# PR #55 审查报告

> 审查分支：`feat/issue-43`  
> 审查范围：Matching V1（推荐卡片、滑动匹配、双向配对、Socket 实时通知、匹配列表）  
> 审查日期：2026-04-22  
> 依据文档：`AGENTS.md`、`docs/design-docs/flutter.md`、`docs/design-docs/nestjs.md`、`docs/design-docs/api.md`、`docs/design-docs/testing.md`、`prd/v1-matching.md`

---

## 审查结论

**审查结论: NEEDS_MAJOR_FIX**

PR #55 整体实现了 Matching V1 的核心功能链路（DB → API → Mobile），代码结构清晰，测试覆盖较完整（单元测试、BDD、Maestro E2E、Golden Test 均已提供）。但存在 **2 个 Blocking 问题** 必须在合并前修复：重复 like 请求会导致 `match:created` 事件重复发射，以及缺少强制要求的 Socket 事件 YAML 契约。此外存在若干 Major 规范违反（NestJS DTO、Flutter DI、Socket 重连、异常信息掩盖等），建议在合并前或紧随其后的修复 PR 中处理。

---

## 阻塞问题（Blocking）

### 1. 重复滑动请求会导致 `match:created` 事件重复发射
- **位置**: `apps/api/src/modules/matching/matching.service.ts:143-179`（`recordSwipeAction` 方法）
- **问题**: 当 `matchActionRepo.save` 捕获到重复键异常（幂等）后，代码未提前返回，而是继续执行 `checkAndCreateMutualMatch` 和 `emitMatchCreated`。具体触发场景：
  1. 用户因网络抖动重复发送 like 请求；
  2. 两个用户并发互相 like，后提交的事务因唯一索引冲突回滚，但回滚后查询到已有 match 并返回 `isMatch: true`。
  这两种情况都会导致同一个 match 的 `match:created` 事件被发射两次，前端会弹出两次配对成功弹窗。
- **影响**: 用户体验严重受损（弹窗重复闪烁），且可能引发前端状态异常。
- **修复**: 在 `recordSwipeAction` 的 `catch (err)` 分支中，若确认是重复键异常，直接 `return { isMatch: false }`，不再进入 `checkAndCreateMutualMatch`。
  ```typescript
  // apps/api/src/modules/matching/matching.service.ts
  } catch (err: unknown) {
    const isDuplicate = ...;
    if (!isDuplicate) throw err;
    // 新增：重复记录说明此前已处理过，直接返回，不再检查 mutual like
    return { isMatch: false };
  }
  ```
- **依据**: `prd/v1-matching.md` US-MATCH-2 技术要求："重复插入被数据库拦截，后端捕获异常返回成功"；`api.md` 要求事件发射必须与事务状态一致，避免重复通知。

### 2. 缺少强制要求的 Socket 事件 YAML 契约
- **位置**: `docs/api-contracts/` 目录下无 `matching-events.yaml`
- **问题**: `api.md` 的「实时通信契约」章节强制要求：
  > 契约文档：`docs/api-contracts/<module>-events.yaml`（事件名、Payload Schema、触发时机）
  
  且「不变量 A」要求后端每新增一个 `server.emit()` 事件，必须同时满足 `docs/api-contracts/<module>-events.yaml` 中定义了该事件的 Schema。当前仅有 `docs/api/socket-events.md`（Markdown 描述），缺少 YAML 格式契约。这会导致自动化契约比对工具（如 `precommit-checks.sh`）无法扫描到事件定义。
- **影响**: 前后端事件契约不同步的风险无法被 CI 检测；新 Agent 无法通过 OpenAPI 生成事件类型。
- **修复**: 新建 `docs/api-contracts/matching-events.yaml`，内容至少包含：
  ```yaml
  events:
    match:created:
      description: 双向匹配成功时推送
      trigger: 第二个用户 Like 后，事务提交成功
      recipients: [user_a_id, user_b_id]
      payload:
        $ref: './matching.yaml#/components/schemas/MatchCreatedEvent'
  ```
- **依据**: `docs/design-docs/api.md` 第 89-113 行（实时通信契约章节）及「不变量 A」清单。

---

## 严重问题（Major）

### 3. `MatchingBloc` 异常处理掩盖具体后端错误信息
- **位置**: `apps/mobile/lib/presentation/blocs/matching/matching_bloc.dart:55-57` 及 `105-107`
- **问题**: `catch (Exception _)` 将所有非 `ProfileIncompleteException` 的异常统一显示为固定的 `networkErrorMessage`（"网络异常，请检查连接后重试"）。当后端返回业务错误（如 "Cannot swipe yourself" / code 1000）时，用户看到的是网络错误提示，无法获知真实原因。
- **影响**: 用户体验受损，问题排查困难。
- **修复**: 区分 `DioException` 的业务错误响应与其他异常。若 `DioException` 携带了后端的 `message`，优先展示后端消息；仅在网络层真正失败时使用兜底文案。
  ```dart
  } on DioException catch (e) {
    final msg = e.response?.data?['message'] as String?;
    emit(MatchingFailure(msg ?? networkErrorMessage));
  } on Exception catch (_) {
    emit(const MatchingFailure(networkErrorMessage));
  }
  ```
- **依据**: `docs/design-docs/api.md` 统一响应体章节（`code/message/data` 结构）；`flutter.md` 要求错误信息必须包含有意义的描述。

### 4. `MatchingSocketClient` 缺少重连机制与错误监听
- **位置**: `apps/mobile/lib/core/network/matching_socket_client.dart`
- **问题**: 仅实现了 `connect()` 和 `disconnect()`，未监听 `onDisconnect`、`onConnectError`、`onError` 事件，也未实现自动重连。移动网络（4G/5G/Wi-Fi 切换、电梯、地铁）极易导致 Socket 断开，断开后用户将永久丢失实时配对通知。
- **影响**: 核心实时功能（配对弹窗）在弱网环境下不可用。
- **修复**: 
  1. 添加 `onDisconnect` 监听器，触发指数退避重连（或使用 `socket_io_client` 的 `reconnect` 选项）；
  2. 添加 `onConnectError` 日志记录；
  3. 将连接状态暴露给 `MatchingBloc`，在断线时提示用户或降级为轮询。
- **依据**: `docs/design-docs/api.md` 第 155-165 行（与 REST API 的降级策略）；`prd/v1-matching.md` 要求 "实时推送通知"。

### 5. `MatchingService` 多处返回内联类型，违反 Service 层 DTO 规范
- **位置**: 
  - `apps/api/src/modules/matching/matching.service.ts:95` — `recordSwipeAction` 返回内联类型 `{ isMatch: boolean; matchId?: string; conversationId?: string }`
  - `apps/api/src/modules/matching/matching.service.ts:423-430` — `getMatchList` 返回内联类型 `{ items: MatchListResponseDto[]; next_cursor: string | null; has_more: boolean }`
  - `apps/api/src/modules/matching/matching.service.ts:454-465` — `getMatchSummary` 返回内联类型
- **问题**: `nestjs.md` 强制要求 "所有出参使用显式 DTO，禁止直接返回 TypeORM Entity"。内联类型虽然结构化，但不是显式 DTO 类/接口，无法被其他模块或契约生成器引用，也不利于维护。
- **影响**: 代码可维护性下降；Controller 层无法明确获知返回结构；不符合模块文档约定。
- **修复**: 
  1. 新建 `SwipeActionResponseDto`（含 `isMatch`, `matchId`, `conversationId`）；
  2. 新建 `MatchListDataDto`（含 `items`, `next_cursor`, `has_more`）；
  3. 新建 `MatchSummaryDto`（与 `MatchCreatedEvent` 结构对齐，或复用 `shared-types` 中的类型）。
- **依据**: `docs/design-docs/nestjs.md` 第 82-86 行（DTO 与校验章节）。

### 6. `MatchingGateway` 的 `emitMatchCreated` 使用内联 Payload 类型
- **位置**: `apps/api/src/modules/matching/matching.gateway.ts:62-92`
- **问题**: `emitMatchCreated` 的参数类型是手写内联对象类型，没有使用 `packages/shared-types` 中已生成的 `MatchCreatedEvent` 类型。`nestjs.md` 要求：
  > Payload 类型：必须使用 `shared-types` 中定义的 DTO/Interface，禁止内联 `any`
- **影响**: 类型不同步风险（如字段名 `avatarUrl` vs `avatar_url`）；契约变更时 Gateway 代码不会被类型检查捕获。
- **修复**: 导入并使用 `@social-world/shared-types` 中的 `MatchCreatedEvent`：
  ```typescript
  import { MatchCreatedEvent } from '@social-world/shared-types';
  // ...
  emitMatchCreated(userAId: string, userBId: string, payloadForA: MatchCreatedEvent, payloadForB: MatchCreatedEvent) { ... }
  ```
- **依据**: `docs/design-docs/nestjs.md` 第 143-151 行（Gateway 规范）。

### 7. `HomePage` 滑动按钮缺少防抖，快速双击导致重复 API 请求
- **位置**: `apps/mobile/lib/presentation/pages/home_page.dart:258-271`（`_submitSwipe`）
- **问题**: `_submitSwipe` 内先执行 `setState` 动画，再通过 `Future.delayed` 发送 BLoC 事件。整个 200ms 动画期间没有任何标志位阻止再次点击。用户快速双击 Like/Pass 按钮会触发两次 `_submitSwipe`，发送两次 `SwipeActionSubmitted`。
- **影响**: 重复 API 请求；可能触发后端幂等逻辑；UI 动画异常。
- **修复**: 添加 `_isAnimating` 标志，在动画期间禁用按钮：
  ```dart
  bool _isAnimating = false;

  void _submitSwipe(String targetId, SwipeAction action) {
    if (_isAnimating) return;
    _isAnimating = true;
    // ... setState animation ...
    Future.delayed(const Duration(milliseconds: 200), () {
      if (!mounted) return;
      context.read<MatchingBloc>().add(...);
      setState(() {
        _isAnimating = false;
        // ... reset animation ...
      });
    });
  }
  ```
- **依据**: `prd/v1-matching.md` US-MATCH-2 要求 "操作幂等"，但前端也应避免无意义重复请求；`flutter.md` 要求状态管理避免竞态。

### 8. `HomePage` `initState` 中不安全使用 `context.read` 与直接实例化 `SecureStorage`
- **位置**: `apps/mobile/lib/presentation/pages/home_page.dart:62-76`（`_initSocket`）
- **问题**: 
  1. `initState` 中调用 `context.read<AuthBloc>().state` 依赖 BuildContext 获取其他 BLoC，在复杂路由场景下可能不安全；
  2. `SecureStorage()` 直接实例化，而 `flutter.md` / `di.dart` 中已将其注册为 `getIt` 单例，应通过依赖注入获取。
- **影响**: 不符合架构规范；可能导致测试时无法 mock `SecureStorage`。
- **修复**: 
  1. 将 `_initSocket` 移至 `didChangeDependencies`，或完全依赖 `BlocListener<AuthBloc, AuthState>`（已在 `build` 中存在）初始化 Socket；
  2. 改为 `getIt<SecureStorage>().readAccessToken()`。
- **依据**: `docs/design-docs/flutter.md` 第 70-81 行（DioClient 单例模式与 DI 规范）。

### 9. `MatchingBloc` `_onExternalMatchReceived` 在非法状态下可能空指针异常
- **位置**: `apps/mobile/lib/presentation/blocs/matching/matching_bloc.dart:121-128`
- **问题**: 仅在 `state is MatchingLoadSuccess || state is MatchingLoadEmpty` 时保存 `_lastListState`。如果 Socket 事件在 `MatchingLoadInProgress`、`MatchingFailure` 或 `MatchingProfileIncomplete` 状态下到达，`_lastListState` 仍为 `null`。用户关闭弹窗后调用 `_onMatchModalDismissed`，执行 `emit(_lastListState!)` 会抛出空指针异常。
- **影响**: App 崩溃。
- **修复**: 为 `_lastListState` 提供 fallback：
  ```dart
  void _onExternalMatchReceived(...) {
    if (state is MatchingLoadSuccess || state is MatchingLoadEmpty) {
      _lastListState = state;
    } else if (_lastListState == null) {
      _lastListState = const MatchingLoadEmpty(mode: 'nearby');
    }
    emit(MatchingMatchCreated(event.matchSummary));
  }
  ```
- **依据**: `flutter.md` 要求 BLoC 状态流转必须覆盖所有边界情况；代码可读性规范要求避免 `!` 非空断言在不安全场景下使用。

---

## 轻微问题（Minor）

### 10. `MatchListItem` 的 `Semantics identifier` 不唯一
- **位置**: `packages/design-system/lib/src/molecules/match_list_item.dart:20`
- **问题**: 所有列表项的 `identifier` 固定为 `'match_list_item'`，Maestro 无法区分和定位特定列表项。
- **修复**: 改为 `'match_list_item_${matchId}'` 或 `'match_list_item_${userId}'`，由调用方传入唯一 key。
- **依据**: `docs/design-docs/flutter.md` 第 256-266 行（Semantics 三要素）。

### 11. `MatchingRepositoryImpl` 中 `limit` 硬编码为 20
- **位置**: `apps/mobile/lib/data/repositories/matching_repository_impl.dart:23` 及 `83`
- **问题**: `getRecommendations` 和 `getMatches` 中 `limit` 硬编码，忽略了接口参数和调用方传入的值。
- **修复**: 将 `limit` 作为可选参数暴露给 Repository，默认 20。
- **依据**: `prd/v1-matching.md` 技术要求支持 `limit` 参数。

### 12. `MatchingService` `buildRecommendations` 重复查询当前用户 profile
- **位置**: `apps/api/src/modules/matching/matching.service.ts:231-237`
- **问题**: 缓存命中路径中，`buildRecommendations` 再次查询 `currentProfile`，而 `getRecommendations` 在缓存未命中路径中已经查询过。可优化为传参避免重复查询。
- **修复**: 修改 `buildRecommendations` 签名接收 `currentProfile?: UserProfileEntity`。
- **依据**: `docs/design-docs/nestjs.md` 性能与可读性规范。

### 13. SQL 中 Haversine 公式重复计算 3 次
- **位置**: `apps/api/src/modules/matching/matching.service.ts:166-222`
- **问题**: 同一条 SQL 中 Haversine 距离被计算了 3 次（SELECT 列、score 公式、WHERE 过滤），可通过子查询或变量优化。
- **修复**: 使用 `HAVING distance <= ?` 替代 WHERE 中的第三次计算，或将距离计算提取为派生表。
- **依据**: `prd/v1-matching.md` 性能要求 "首页推荐卡片加载时间 < 800ms（P95）"。

### 14. `MatchingGateway` CORS 配置方式非标准
- **位置**: `apps/api/src/modules/matching/matching.gateway.ts:28-35`
- **问题**: `afterInit` 中直接修改 `server.engine.opts.cors.origin` 属于内部属性访问，非标准 Socket.io API，可能在版本升级时失效。
- **修复**: 使用 `@WebSocketGateway({ namespace: 'matching', cors: { origin: '...' } })` 装饰器参数。
- **依据**: `@nestjs/websockets` 官方文档。

### 15. `MatchingSocketClient` URL 协议未从 HTTP 转换为 WebSocket
- **位置**: `apps/mobile/lib/core/network/matching_socket_client.dart:24`
- **问题**: `baseUrl` 若为 `http://`/`https://`，代码未显式转换为 `ws://`/`wss://`。虽然 `socket_io_client` 在 `transports: ['websocket']` 下可能自动处理，但存在环境兼容性风险。
- **修复**: 
  ```dart
  final serverUrl = baseUrl
    .replaceFirst(RegExp(r'/api/v1/?$'), '')
    .replaceFirst('http://', 'ws://')
    .replaceFirst('https://', 'wss://');
  ```
- **依据**: `socket_io_client` 官方文档要求 WebSocket 传输使用 `ws://`/`wss://` 协议。

### 16. BDD 测试 "Cache hit returns quickly" 时间断言过于严格
- **位置**: `tests/bdd/step-definitions/matching.steps.js:112-113`
- **问题**: 断言 `duration < 100ms`，在 CI 慢节点或高负载时容易不稳定失败。
- **修复**: 放宽为 `< 200ms` 或仅断言使用了缓存（mockRedis.get 被调用）而不断言绝对时间。
- **依据**: `docs/design-docs/testing.md` 第 304 行（测试失败时错误信息必须包含修复指令），过于严格的断言会导致无意义的 flaky 失败。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| **功能完整性** | ⚠️ 部分通过 | 核心功能（推荐、滑动、双向匹配、Socket 通知、匹配列表）均已实现。但重复滑动会导致事件重复发射（Blocking #1），异常处理掩盖业务错误（Major #3），Socket 缺少重连（Major #4）。 |
| **代码规范性** | ❌ 未通过 | 违反 `nestjs.md` Service DTO 规范（Major #5）、Gateway Payload 类型规范（Major #6）；违反 `flutter.md` DI 规范（Major #8）；`matching-events.yaml` 缺失（Blocking #2）。 |
| **全链路完整性** | ✅ 通过 | DB Migration → Entity → Service → Controller → OpenAPI 生成 → Mobile Repository → BLoC → UI 均已连通。 |
| **规范符合性** | ❌ 未通过 | Socket 事件契约缺少 YAML 文件（Blocking #2）；`MatchingGateway` 未使用 `shared-types`（Major #6）；`MatchingService` 多处内联返回类型（Major #5）。 |
| **测试覆盖** | ✅ 通过 | 后端单元测试（`matching.service.spec.ts`）、Flutter BLoC 测试（`matching_bloc_test.dart`、`match_list_cubit_test.dart`）、BDD（`matching.feature` + steps）、Maestro E2E（`swipe_recommendations.yaml`、`view_match_list.yaml`）、Golden Test（`matching_components_golden_test.dart`）均已覆盖核心路径。 |
| **文档同步** | ⚠️ 部分通过 | `docs/modules/matching.md`、`docs/api/matching.md`、`docs/api/socket-events.md` 已更新；技术债务已录入 `tech-debt-tracker.md`。但缺少 `docs/api-contracts/matching-events.yaml`（Blocking #2）。 |

---

## 风险与建议

### 高风险
1. **重复 `match:created` 事件**：这是最直接的用户体验破坏点。建议在合并前必须修复（Blocking #1）。
2. **Socket 断线无恢复**：Matching 的核心卖点是 "实时配对弹窗"，如果 Socket 在弱网下永久断开，产品体验将严重降级（Major #4）。建议在下个迭代中优先补齐重连机制。

### 中风险
3. **异常信息掩盖**：用户看到 "网络异常" 时，实际上可能是 "不能滑动自己" 或 "资料不完整"。这会增加客服成本（Major #3）。
4. **DTO 规范债务**：Service 层内联类型虽然当前功能正确，但随着模块扩展，会成为类型同步的隐患（Major #5、#6）。

### 建议
- **合并前必须修复**：Blocking #1（重复事件）、Blocking #2（事件契约 YAML）。
- **建议同批次修复**：Major #3（异常消息）、Major #5（DTO）、Major #6（Gateway 类型）、Major #7（按钮防抖）、Major #9（空指针保护）。
- **建议下个迭代处理**：Major #4（Socket 重连）、Major #8（DI 规范）、Minor #13（SQL 性能优化）。
- **已认可的技术债务**：`city` 模式简化、`last_login_at` 命名偏差、复合索引缺失均已妥善记录在 `docs/exec-plans/tech-debt-tracker.md`，无需在本次 PR 中解决。
