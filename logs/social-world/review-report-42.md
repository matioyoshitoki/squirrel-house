## 审查结论

审查结论: **NEEDS_FIX**

本 PR（`feat/issue-4`）核心目标是实现 Matching 模块的 `last_active_at` 活跃度更新机制、空状态降级策略，并同步配套文档与 E2E 测试。整体架构链路完整（DB Migration → API → Mobile），代码质量良好，但存在 **1 个阻塞级问题（Maestro E2E 语法错误）** 和 **3 个严重问题（测试覆盖缺失、状态机缺陷）**，必须在合并前修复。

---

## 阻塞问题（Blocking）

### 1. Maestro E2E 流程 YAML 语法错误导致测试无法解析
- **位置**: `.maestro/flows/matching/swipe_and_match.yaml:29-30, 33-34`
- **问题**: `assertVisible` 和 `tapOn` 使用了「命令缩写 + 同行属性」的混合格式，Maestro 无法正确解析：
  ```yaml
  - assertVisible: "配对成功!"
    optional: true
  - tapOn: "继续滑动"
    optional: true
  ```
  在 YAML 中，这会将列表项解析为包含两个顶层键（`assertVisible` 和 `optional`）的 Mapping，而 Maestro 要求每个列表项必须是单键 Mapping（命令名 → 配置对象）。
- **影响**: `make test-mobile-e2e` 执行到该流程时会直接报错，Matching 模块的 E2E 验证完全失效。
- **修复**: 改为显式对象语法：
  ```yaml
  - assertVisible:
      text: "配对成功!"
      optional: true
  - tapOn:
      text: "继续滑动"
      optional: true
  ```
- **依据**: `.maestro/flows/matching/recommendations.yaml:23-28` 使用了正确的嵌套格式；`swipe_and_match.yaml` 与之不一致。

---

## 严重问题（Major）

### 1. `recordAction` 更新 `last_active_at` 的行为缺少单元测试覆盖
- **位置**: `apps/api/src/modules/matching/tests/matching.service.spec.ts`
- **问题**: PR 在 `matching.service.ts:323-327` 新增了滑动成功后更新 `last_active_at` 的逻辑，但 `matching.service.spec.ts` 中仅对 `getRecommendations` 补充了断言（`"should update last_active_at before querying"`），`recordAction` 的全部 7 个测试用例均未断言 `mockProfileRepo.update` 被调用。
- **影响**: 后续重构可能意外删除或破坏该活跃度更新逻辑，导致推荐排序中的「近期活跃」维度失真。
- **修复**: 在 `describe("recordAction")` 下新增测试：
  ```typescript
  it("should update last_active_at after recording action", async () => {
    mockMatchActionRepo.insert.mockResolvedValue({});
    mockMatchActionRepo.findOne.mockResolvedValue(null);

    await service.recordAction("user-1", "user-2", "like");

    expect(mockProfileRepo.update).toHaveBeenCalledWith(
      { user_id: "user-1" },
      expect.objectContaining({ last_active_at: expect.any(Date) }),
    );
  });
  ```
- **依据**: `docs/design-docs/testing.md` 要求「新增行为必须有对应单元测试覆盖」；`git diff` 显示 `matching.service.ts` 新增了 `profileRepo.update` 调用，但 `matching.service.spec.ts` 未同步增加对应断言。

### 2. 登录流程更新 `last_active_at` 的行为缺少单元测试断言
- **位置**: `apps/api/src/modules/auth/tests/auth.service.spec.ts`
- **问题**: PR 在 `auth.service.ts` 登录成功后新增了 `profileRepo.update({ last_active_at: new Date() })`，但 `auth.service.spec.ts` 的 `"should login existing user and update lastLoginAt"` 测试仅断言了 `userRepo.update` 被调用，未断言 `profileRepo.update`。
- **影响**: 登录活跃度更新逻辑无测试守护，回归风险高。
- **修复**: 在该测试末尾追加：
  ```typescript
  expect(profileRepo.update).toHaveBeenCalledWith(
    { user_id: "user-1" },
    expect.objectContaining({ last_active_at: expect.any(Date) }),
  );
  ```
- **依据**: `git diff` 中 `auth.service.ts` 新增了 `profileRepo.update`，但 `auth.service.spec.ts` 的 diff 仅增加了 `update: jest.fn()` mock，无对应断言。

### 3. `ExternalMatchReceived` 在空状态会丢失当前 mode，导致弹窗关闭后模式回退到 nearby
- **位置**: `apps/mobile/lib/presentation/blocs/matching/matching_bloc.dart:112-132`
- **问题**: 当 `MatchingBloc` 当前状态为 `MatchingLoadEmpty(mode: 'city')` 时收到 socket `match:created` 事件，`_onExternalMatchReceived` 会进入 `else` 分支：
  ```dart
  } else {
    emit(MatchingMatchCreated(event.matchSummary));
  }
  ```
  `MatchingMatchCreated` 的默认值 `previousMode = 'nearby'`。用户关闭弹窗后，`_onMatchModalDismissed` 会发射 `MatchingLoadEmpty(mode: 'nearby')`，导致用户从「同城」被意外切回「附近」。
- **影响**: 用户在空数据态收到配对通知后，关闭弹窗发现模式被切换，体验受损。
- **修复**: 为 `MatchingLoadEmpty` 和 `MatchingFailure` 也保留正确的 `previousMode`：
  ```dart
  } else if (currentState is MatchingLoadEmpty) {
    emit(MatchingMatchCreated(
      event.matchSummary,
      previousCards: const [],
      previousMode: currentState.mode,
    ));
  } else {
    emit(MatchingMatchCreated(event.matchSummary));
  }
  ```
- **依据**: `docs/modules/matching.md` 中 Socket.io 事件监听设计明确说明状态应流转至 `MatchingMatchCreated` 并保留上下文；`designs/issue-4/design-spec.md` 5.3 节要求模式切换需用户明确操作，不应由弹窗关闭隐式触发。

---

## 轻微问题（Minor）

### 1. `docs/design-docs/flutter.md` 未按设计规范同步更新
- **位置**: `docs/design-docs/flutter.md`
- **问题**: `designs/issue-4/design-spec.md` 第 7.2 节明确要求更新 `docs/design-docs/flutter.md`（「在『Design-Only PR 验收 Checklist』中补充 Matching 相关组件的 Widgetbook 用例要求」），但 PR 中该文件无变更。
- **修复**: 在 `flutter.md` 的 Widgetbook 注册规范或 Design-Only PR Checklist 中补充 `SwipeCard`、`SwipeCardStack`、`MatchSuccessDialog`、`SwActionButton`、`MatchListItem`、`SwModeSelector` 的用例要求。
- **依据**: `designs/issue-4/design-spec.md` 第 7.2 节文档同步清单。

### 2. `MatchSuccessDialog` 使用硬编码颜色而非 Design Token
- **位置**: `packages/design-system/lib/src/organisms/match_success_dialog.dart:25`
- **问题**: `Material(color: Colors.black.withOpacity(0.6), ...)` 使用了硬编码颜色。`designs/issue-4/design-spec.md` 2.1 节规定「视觉系统严格复用项目已有 Design Token，禁止新增硬编码色值」。
- **修复**: 在 `tokens.dart` 中新增 `scrim` 或 `overlay` Token（如 `Color(0x99000000)`），并在组件中引用。
- **依据**: `designs/issue-4/design-spec.md` 第 2.1 节色板规范。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 全链路完整性 | ✅ | DB Migration (`AddLastActiveAtToUserProfiles`) → Entity 同步 → API (`matching.service.ts`) → Mobile (`MatchingBloc` / `HomePage`) → Design System (`SwModeSelector` / `EmptyRecommendationView`) 链路完整。 |
| 规范符合性 | ⚠️ | NestJS 分层、Flutter BLoC 状态管理、Design Token 使用基本符合规范；但 `MatchSuccessDialog` 存在硬编码颜色，`flutter.md` 未同步更新。 |
| 测试覆盖 | ❌ | `matching.service.spec.ts` 缺少 `recordAction` 的 `last_active_at` 断言；`auth.service.spec.ts` 缺少登录时 `last_active_at` 断言。Maestro E2E 流程存在语法错误。 |
| 文档同步 | ⚠️ | `docs/api/matching.md`、`docs/exec-plans/tech-debt-tracker.md`、`prd/v1-matching.md`、`docs/modules/matching.md` 均已同步；但 `docs/design-docs/flutter.md` 未按设计规范更新。 |
| OpenAPI 契约 | ✅ | `docs/api-contracts/matching.yaml` 已补充 `mode`、`limit` 参数及 `/matches` 接口定义；`packages/shared-types` 与 `apps/mobile/lib/models/generated` 的生成代码一致。 |
| Accessibility (Maestro) | ⚠️ | 新增组件（`SwActionButton`、`MatchListItem`、`SwModeSelector`、`EmptyRecommendationView`）均正确使用 `Semantics(identifier/label/button)` + `onTapDown`；但 `swipe_and_match.yaml` 语法错误导致 E2E 不可用。 |

---

## 风险与建议

1. **性能风险**：`getRecommendations` 在缓存命中前仍先执行 `UPDATE user_profiles SET last_active_at = NOW()`，在高并发下可能给 DB 带来写压力。建议在后续迭代中评估是否可将活跃度更新异步化（如通过消息队列或定时任务合并写入），或降低更新频率（如 5 分钟内只更新一次）。

2. **状态机健壮性**：`MatchingBloc` 的 `_onEmptyStateFallbackRequested` 中，`MatchingMatchCreated` 状态未被处理，会默认回退到 `'nearby'`。虽然当前 UI 不会从该状态触发降级，但建议显式覆盖所有状态分支，避免未来 UI 变更引入 Bug。

3. **缓存一致性**：`recordAction` 在 MySQL `ER_DUP_ENTRY`（重复滑动）时直接返回，**未更新** `last_active_at`。从业务语义看，用户进行了滑动操作（即使重复），活跃度也应被刷新。建议在 `ER_DUP_ENTRY` 处理分支中也追加 `last_active_at` 更新，或至少在 PRD 中明确该边界行为。

4. **Widgetbook 产物缺失**：PR 中未包含 Golden File 截图或 Widgetbook 运行截图。虽然本 PR 不是纯 Design-Only PR，但涉及大量 design-system 新增组件，建议在 PR 描述中补充至少一张 Widgetbook 运行截图，便于视觉走查。
