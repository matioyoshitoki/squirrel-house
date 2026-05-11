## 审查结果：PR #6 (`feat/issue-1`)

| 检查项 | 状态 | 说明 |
|----
----|------|------|
| 全链路完整性 | ❌ | 后端新增了 Auth/Admin/Blocks/Reports/ ContentModeration/Health 模块及完整 DB Migration，但 **Mobile 端零实现**（无 BLo C、无 Page、无 Network/Storage 层）。同时 Profile / Matching / IM 的 API Control ler 完全缺失，仅有 OpenAPI YAML 骨架和 DB 表。 |
| 规范符合性 | ❌ | NestJS 分 层基本正确，但 AuthController 与 OpenAPI 契约严重不符；缺失 `jwt-auth.guard.ts` 和 `profile-completion.guard.ts`；AuthService.login 返回硬编码 stub；BDD Step 全 为空实现。 |
| 契约一致性 | ❌ | `auth.yaml` 约定 `phone+sms_code`，但 Controll er 实现为 `phone+birth_date`；`matching.yaml`/`profile.yaml`/`im.yaml` 仅有 YAML ，无对应后端 DTO 及 `shared-types`/`mobile models` 生成代码。 |
| 测试覆盖 | ❌ | 仅 2 个单元测试文件（`auth.service.spec.ts`、`age-calculator.spec.ts`），BDD S tep Definitions 全部为 `() => {}` 空桩，未验证任何业务逻辑。 |
| 文档同步 | ❌ | `docs/QUALITY_SCORE.md` 未更新（仍显示 C 级/未开始）；`docs/modules/auth.md` 引用了不存在的 `jwt-auth.guard.ts`，会导致 `make check-docs` 死链检查失败；执行 计划状态未刷新。 |
| 可运行性 | ❌ | 无法直接运行工具，但通过代码分析可判定 `ma ke check-docs`（死链）、`make check-contract-sync`（缺失 generated types for mat ching/profile/im）必然失败；`make test` 中 BDD 虽然空跑通过，但不具备验收价值。 |

---

## 问题详情

### 1. Mobile 端完全缺失 — 违反全链路交付原则
**问题 **：`apps/mobile/lib/` 下仅有 `main.dart`（一个空壳 Widget）和 OpenAPI 生成的 `models/generated/auth/`。没有 `presentation/blocs/`、`presentation/pages/`、`core/network/`、`core/storage/` 等任何实现。 **影响**：严重违反 AGENTS.md 黄金原 则第 1 条「一个需求必须同时包含 DB + API + Mobile 的实现」。 **建议**：
- 补
齐 `apps/mobile/lib/presentation/blocs/auth/`（Auth BLoC / Cubit）
- 补齐 `apps
/mobile/lib/presentation/pages/login_page.dart`
- 补齐 `apps/mobile/lib/core/st
orage/secure_storage.dart`（Token 安全存储）
- 补齐 `apps/mobile/lib/core/netwo
rk/auth_interceptor.dart`（Dio 401 自动刷新）

---

### 2. Auth API 契约与实
现严重不符 **问题**：
- `docs/api-contracts/auth.yaml` 约定 `POST /auth/login`
请求体为 `{ phone, sms_code }`；
- 实际 `apps/api/src/modules/auth/auth.control
ler.ts` 中：
  - `POST /auth/register` 接收 `{ phone, birth_date }`（DTO 中没有
`sms_code`）
  - `POST /auth/login` 不接收任何 Body，直接调用 `authService.logi
n()` 返回硬编码 stub token。

**影响**：前后端无法联调，OpenAPI 生成的 ClientSDK 与实际接口不兼容。 **建议**：
1. 修改 `auth.controller.ts` 的 `login` 方
法，使其接收 `LoginRequest`（`phone` + `sms_code`）；
2. 注册/登录合一（按执行
计划决策日志）使用 `POST /auth/login` 统一入口；
3. 在 DTO 层增加 `sms_code` 字
段校验，移除不合理的 `birth_date` 从 Auth DTO（年龄校验应在 Profile 模块）。


---

### 3. 关键 API 模块仅有契约 YAML，无 Controller/Service 实现
**问题**： 以下 YAML 已存在，但对应 NestJS 模块完全缺失：
- `docs/api-contracts/profile.ya
ml` → 无 `apps/api/src/modules/profile/`
- `docs/api-contracts/matching.yaml` →
无 `apps/api/src/modules/matching/`
- `docs/api-contracts/im.yaml` → 无 `apps/a
pi/src/modules/im/`

**影响**：DB Migration 已创建了 `user_profiles`、`match_actions`、`matches`、`conversations`、`messages` 等表，但没有任何业务代码操作这些 表，形成「有库无 API」的断层。 **建议**：
- 按 NestJS 分层规范补齐 Profile /
Matching / IM 的 Controller、Service、DTO、Module；
- 或在本次 PR 中删除未实现
的 YAML，避免 `make check-contract-sync` 失败（见问题 5）。

---

### 4. 缺
失核心 Guard **问题**：
- `docs/modules/auth.md` 引用了 `apps/api/src/common/g
uards/jwt-auth.guard.ts`，但文件不存在；
- `docs/exec-plans/TASK-profile.md` 要
求的 `ProfileCompletionGuard` 亦不存在。

**影响**：所有新增模块的接口均处于无 认证状态，且无法拦截未完善资料的用户进入匹配流程。 **建议**：
1. 创建 `jwt-a
uth.guard.ts`，从 `Authorization: Bearer <token>` 提取并校验 JWT；
2. 创建 `pro
file-completion.guard.ts`，检查 `user_profiles.is_complete`；
3. 将 `JwtAuthGua
rd` 应用至需要保护的 Controller（Blocks、Reports、Admin 等）；
4. 将 `ProfileCo
mpletionGuard` 应用至 Matching / IM 相关路由。

---

### 5. `shared-types` 
与 `mobile models` 生成不完整 **问题**：
- `packages/shared-types/src/generate
d/` 下仅有 `auth`，无 `profile`/`matching`/`im`；
- `apps/mobile/lib/models/ge

erated/` 下同样仅有 `auth`；
- `packages/shared-types/src/index.ts` 仅 `export 
* from "./generated/auth/src"`。

**影响**：`make check-contract-sync` 必然会 报差异（YAML 存在但生成代码缺失）。 **建议**：
- 运行 `make generate-types` 
重新生成所有契约对应的 TypeScript 和 Dart 代码；
- 确保 `index.ts` 同步导出所有
生成模块。

---

### 6. BDD 验收测试全为空桩
**问题**：`tests/bdd/step-definitions/common.js` 中所有 `given/when/then` 均为 `() => {}` 空函数；各 `.steps.js` 文件仅做 `loadFeature + autoBindSteps(commonSteps)`，无额外断言。 **影响** ：`make test-bdd` 虽然可能 exit 0，但不验证任何业务规则，失去验收测试意义。 * *建议**：
- 至少为本次已实现的核心流程（手机号注册/登录、年龄校验）编写真实 Ste
p Definition；
- 使用 `supertest` 调用实际 API 端口，结合测试数据库进行断言。


---

### 7. 文档不同步与死链
**问题**：
- `docs/QUALITY_SCORE.md` 仍显示 A
uth / Profile / Matching / IM 为 C 级「未开始」；
- `docs/modules/auth.md` 包含
死链 `apps/api/src/common/guards/jwt-auth.guard.ts`；
- `docs/exec-plans/INDEX.
md` 中各任务状态仍为 🟡 active，未根据实际完成度更新。

**影响**：`make check-docs` 死链检查阶段会失败。 **建议**：
1. 更新 `docs/QUALITY_SCORE.md`，将已
完成的领域（Auth 骨架、DB Migration、Design System）评分提升；
2. 修复或移除 `d
ocs/modules/auth.md` 中的死链；
3. 更新执行计划状态，已完成项标记为 🟢 complete
d，未开始项保持清晰边界。

---

### 8. AuthService 存在硬编码 Stub
**问题** ：`apps/api/src/modules/auth/auth.service.ts` 中：
- `register()` 返回 `{ user_
id: "stub-user-id" }` 而不写数据库；
- `login()` 返回硬编码 `access_token: "stu
b-access-token"`。

**影响**：接口无法真正创建用户或签发有效 JWT，端到端不可运 行。 **建议**：
- `register()` 使用 `Repository<UserEntity>` 创建真实用户记
录（UUID v7）；
- `login()` 实现 JWT 签发（Access 15min + Refresh 7天），并按决
策日志实现 Token Rotation 与 Redis 黑名单。

---

### 9. 部分模块路由缺少版
本前缀一致性校验 **问题**：`HealthController` 使用 `@Controller("health")`，在 `main.ts` 中全局前缀为 `api/v1`，最终路径为 `/api/v1/health`，这没问题。但 Blocks/Reports/Admin 等模块的路径设计在 YAML 中未体现，无法确认是否与契约一致。 ** 建议**：为所有已实现模块补充对应 OpenAPI YAML 契约，并保持路径一致。

---

#
## 10. PR 范围过大，违反迭代边界
**问题**：`feat/issue-1` 一次性新增了 6 个后端 模块 + 全套 DB Migration + Design System + BDD 桩，但 Mobile 缺失、核心 Guard 缺 失、契约与实现不符。 **建议**：后续 PR 应严格按 Issue 拆分（如 `feat/auth-login`、`feat/profile-crud`），每个 PR 必须保证 DB + API + Mobile 全链路可运行后再 提交。

---

## 总结

**本 PR 不能合并（Not LGTM）**。 虽然代码量在基础 设施层面（DB Migration、通用过滤器/拦截器、部分 Admin/Blocks/Reports 模块）有可 观进展，但存在以下 blocker：
1. **Mobile 端零实现**，违反全链路原则；
2. **Aut
h 契约与实现不符**，无法联调；
3. **Profile / Matching / IM 有 DB 无 API**；
4. **BDD 测试为空桩**，无验收价值；
5. **文档死链 + 契约同步失败**，机械检查无法
通过。

请按上述问题清单逐条修复后重新提交。