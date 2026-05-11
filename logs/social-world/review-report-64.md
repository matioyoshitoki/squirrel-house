## 审查结论

审查结论: PASS

## 阻塞问题（Blocking）

无。

## 严重问题（Major）

无。

## 轻微问题（Minor）

### 1. PR Body 自检清单未同步更新
- **位置**: PR #64 Body「自检清单」
- **问题**: 开发者已在最新提交（`420482c`）中修复 `ts-jest` 依赖版本并验证命令通过，但 PR Body 中的自检清单仍显示 `[ ] make test` 与 `[ ] make lint` 未勾选。
- **影响**: 审查者无法第一时间从 PR Body 确认检查项实际执行状态，增加沟通成本。
- **修复**: 将 PR Body 中以下两项更新为 `[x]`：
  - `[x] make test`
  - `[x] make lint`
  - `[x] make check-contract-sync`
- **依据**: `docs/exec-plans/templates/TEMPLATE-fullstack-feature.md`「交付前检查清单」要求 developer 在 PR 中明确标记执行结果。

## 上一轮问题修复状态

| # | 问题标题 | 上一轮级别 | 修复状态 | 说明 |
|---|---------|-----------|---------|------|
| 1 | BDD 离线推送测试断言了未与实际代码集成的 Mock | Major | ✅ 已修复 | 已在上一轮修复并验证 |
| 2 | 模板交付前检查清单未完整执行：`make test` 与 `make lint` 未运行 | Major | ✅ 已修复 | `ts-jest` 已从 `^29.2.0` 升级至 `^29.4.9`（`apps/api/package.json`、`apps/api/package-lock.json`、`pnpm-lock.yaml`）；开发者提交 `420482c` 明确说明「安装依赖并验证 make test/make lint/make check-contract-sync 全部通过」。PR Body 清单未更新为 `[x]`，归入本轮 Minor |
| 3 | 系统消息创建后未通过 Socket.io 实时通知 | Major | ✅ 已修复 | 已在上一轮修复并验证 |
| 4 | Widgetbook 未注册新 IM 组件用例 | Major | ✅ 已修复 | 已在上一轮修复并验证 |
| 5 | 文档与代码实现不一致 | Major | ✅ 已修复 | 已在上一轮修复并验证 |

### 上一轮 Minor 问题修复状态

| # | 问题标题 | 修复状态 | 说明 |
|---|---------|---------|------|
| 1 | `ImBloc` 断连/离线时丢失分页状态 | ✅ 已修复 | `MessageSending` 状态新增 `hasMore`、`nextCursor`、`lastReadMessageId`、`lastReadSequenceId` 字段；`_onImDisconnectRequested`、`_onSocketDisconnected`、`_onReceiveMessage`、`_onMessageAckReceived`、`_onMessageDelivered` 均优先从 `MessageSending` 状态恢复分页字段（`41f7f1c`） |
| 2 | `im.service.ts:52` 类型绕过 `as unknown as string` | ✅ 已修复 | 已移除 `as unknown as string`，改为 `where: { user_id: Not(userId) }` |
| 3 | `ImBloc` 注册为 `LazySingleton` 而非 Factory | ✅ 已修复 | `apps/mobile/lib/core/di/di.dart:54` 已改为 `registerFactory<ImBloc>` |
| 4 | `ImSocketClient.connect()` 未释放旧 Socket | ✅ 已修复 | `im_socket_client.dart:37` 开头已增加 `disconnect()` 调用，防止旧连接泄漏 |

## 修复引入的新问题

无。

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ✅ | 核心 IM（发送/接收/已读/会话列表）已实现；离线推送骨架已完成；系统消息实时广播链路已贯通 |
| 代码规范性 | ✅ | 命名符合项目规范；无硬编码魔法值；`as unknown as string` 已移除；Entity 已在 `ImModule` 注册；`data-source.ts` 通过 glob 覆盖新 Entity |
| 全链路完整性 | ✅ | DB（Migration + Entity）→ API（Controller/Gateway/Service/PushModule）→ Mobile（Page/BLoC/Repository/Socket）链路已贯通；`shared-types` 与 `mobile` 生成模型已同步 |
| 规范符合性 | ✅ | Widgetbook 已按执行计划注册新组件用例；`docs/modules/im.md` 与代码实现已一致；`make arch-check` 已通过 |
| 模板检查清单执行 | ✅ | 开发者已通过提交 `420482c` 修复 `ts-jest` 依赖并验证 `make test`/`make lint`/`make check-contract-sync` 通过；PR Body 未更新 `[x]` 标记（Minor） |
| 测试覆盖 | ✅ | 后端单元测试（`im.service.spec.ts`）、Flutter BLoC 测试（`im_bloc_test.dart`）、Golden 测试、BDD 步骤均非空实现且 assert 了具体状态转换值 |
| 文档同步 | ✅ | `docs/api/im.md`、`docs/api/socket-events.md`、`docs/api/INDEX.md`、`docs/api-contracts/im.yaml`、`docs/modules/im.md` 均已更新 |

## 自检结果

- **State 不可变性**: ✅ 通过（`Message.copyWith` 正确使用 `?? this.field` 回退；`MessagesLoaded` / `MessageSending` / `MessageSent` 状态迁移均显式传递字段）
- **生命周期**: ✅ 通过（`ImBloc.close()` 正确调用 `_socketClient?.disconnect()`；`ImSocketClient.disconnect()` 置空 `_socket`）
- **边界路径**: ✅ 通过（`MessageSending` 断连/离线时分页字段 `hasMore`/`nextCursor` 已正确保留；`ImSocketClient.connect()` 先 `disconnect()` 再创建新实例；`sendMessage` 拒绝 `type === "system"` 的用户请求）
- **回归**: ✅ 通过（`ImBloc` 改为 `Factory` 后，`app.dart` 中 `BlocProvider(create: (_) => getIt<ImBloc>())` 仍保证单实例生命周期；未发现影响未修改文件的预期行为）
- **测试真实性**: ✅ 通过（BLoC 测试 assert 了具体状态转换值如 `messages.length`、`hasMore`、`isLoadingMore`；后端单元测试 assert 了返回值与 Repository 调用参数）

## 风险与建议

1. **本 PR 涉及状态管理与 Socket 生命周期变更，建议合并前人工确认边界场景**：特别是 `ImBloc` 在页面重建时作为 `Factory` 由 `BlocProvider` 托管，需确认 `app.dart` 中 `BlocProvider` 的重建频率不会导致 Socket 反复重连。

2. **PR Body 清单 hygiene**：建议开发者在合并前将 PR Body 中 `make test`、`make lint`、`make check-contract-sync` 标记为 `[x]`，以保持交付可追溯性。
