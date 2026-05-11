## 审查结果

| 检查项 | 状态 | 说明 |
|--------|------|------|
| 全链路完整性 | ✅ | 包含 DB 迁移/实体、API Controller/Service、Mobile 页面/BLoC、Design System 组件、文档更新 |
| 规范符合性 | ⚠️ | 后端代码符合 NestJS 规范；Mobile 页面存在与 wireframe 不一致的交互细节 |
| 契约一致性 | ✅ | `packages/shared-types` 与 Mobile generated models 和 API OpenAPI 定义基本同步 |
| 测试覆盖 | ❌ | API 单元测试 34/34 通过，但 BDD 测试有 1 个失败（profile_complete_minimum_profile） |
| 文档同步 | ✅ | 新增/更新了 `docs/modules/profile.md`、`docs/api-contracts/profile.yaml`、`docs/api-contracts/oss.yaml`、`docs/modules/INDEX.md` |
| UI/设计一致性 | ❌ | 前端/UI 变更缺少 `designs/issue-3/` 设计资产，且实现与 wireframe 存在多处不一致 |

---

## 问题详情

### 1. 设计资产缺失（硬规则违规）

**问题**：PR 包含大量前端/UI 变更文件（`apps/mobile/lib/presentation/pages/profile_setup_page.dart`、`profile_edit_page.dart`，以及 `packages/design-system/lib/src/molecules/` 下的 `sw_bottom_sheet.dart`、`sw_gender_selector.dart`、`sw_photo_grid.dart`、`sw_range_slider.dart`、`sw_tag_chip.dart` 等），但项目根目录下**不存在 `designs/` 目录**，更无 `designs/issue-3/design-spec.md`、`mockup.html` 或 `wireframe.svg`。`.design-assets.txt` 亦不存在。

**影响**：无法验证颜色、字号、间距、圆角、布局结构是否符合设计规范，前端变更审查结论直接判定为 ❌。

**建议**：
1. 在 `designs/issue-3/` 目录下补充设计资产：
   - `design-spec.md`：明确颜色、字号、间距、圆角、组件使用规范
   - `mockup.html` 或 `wireframe.svg`：提供视觉原型供还原度对比
2. 若 Issue #3 的设计资产已存在于其他位置，请迁移至 `designs/issue-3/` 并更新 `.design-assets.txt`（如需要）。

---

### 2. BDD 测试失败：`Missing required field blocks completion`

**问题**：`make test-bdd` 运行时，`tests/bdd/step-definitions/profile_complete_minimum_profile.steps.js` 失败：

```
Expected value to strictly be equal to: 1000
Received: 0
```

场景为：用户未上传头像时保存资料，系统应返回 `code=1000` 拒绝。

**根因**：
- `prd/v1-profile.md` 第 58 行明确规定最小资料的**必填字段**包含 `avatar_url`。
- 但 `apps/api/src/modules/profile/dto/create-profile.dto.ts` 中 `avatar_url` 使用了 `@IsOptional()`，未标记为必填。
- `apps/api/src/modules/profile/profile.service.ts` 的 `createProfile()` 方法在校验必填字段时，仅检查了 `nickname`、`gender`、`birth_date`、`latitude`、`longitude`，**未校验 `avatar_url`**：

```typescript
if (
  !dto.nickname ||
  !dto.gender ||
  !dto.birth_date ||
  !dto.latitude ||
  !dto.longitude
) { ... }
```

**建议**：
1. 修改 `apps/api/src/modules/profile/dto/create-profile.dto.ts`，移除 `avatar_url` 的 `@IsOptional()` 装饰器（或添加 `@IsNotEmpty()`）。
2. 修改 `apps/api/src/modules/profile/profile.service.ts`，在 `createProfile()` 的必填字段校验中加入 `!dto.avatar_url`：
   ```typescript
   if (
     !dto.nickname ||
     !dto.gender ||
     !dto.birth_date ||
     !dto.latitude ||
     !dto.longitude ||
     !dto.avatar_url
   ) { ... }
   ```
3. 同步更新 OpenAPI 定义 `docs/api-contracts/profile.yaml` 中 `UpdateProfileRequest` 的 `avatar_url` 为 `required`。
4. 重新运行 `make generate-types` 和 `make test-bdd` 确保契约与测试同步。

---

### 3. Mobile 页面实现与 UI Wireframe 不一致

基于现有 `docs/design-docs/ui-specs/profile/` 中的 wireframe 文档，发现以下不一致：

#### 3.1 `ProfileSetupPage` — "下一步" 按钮状态控制不符

**问题**：`docs/design-docs/ui-specs/profile/profile-setup-wireframe.md` Step 2 的交互状态规定：
> 「下一步」按钮：`isDisabled = true`（昵称、性别、出生日期为必填）

但 `apps/mobile/lib/presentation/pages/profile_setup_page.dart` 中，按钮始终可点击（仅通过 `_onSave` 最后统一校验），没有根据当前步骤的填写情况动态禁用：

```dart
SwButton(
  text: _currentStep < 2 ? '下一步' : '完成',
  variant: SwButtonVariant.filled,
  isLoading: isLoading,
  onPressed: isLoading ? null : _onNext,
)
```

**建议**：在 `_buildStep()` 或 `_onNext` 中增加步骤级字段校验，当当前步骤必填字段未完整填写时，禁用按钮或给出更明确的引导。

#### 3.2 `ProfileEditPage` — 标签数量限制与 wireframe 矛盾

**问题**：`docs/design-docs/ui-specs/profile/profile-edit-wireframe.md` 中规定：
> 标签数量限制：最多 10 个

但 `prd/v1-profile.md` 和 `apps/api/src/modules/profile/profile.service.ts` 均限制为 **8 个**。 wireframe 与 PRD 存在冲突，应以 PRD 为准，但 wireframe 文档需修正以避免后续歧义。

**建议**：
- 确认标签上限为 8（与后端/PRD 一致）。
- 修改 `docs/design-docs/ui-specs/profile/profile-edit-wireframe.md`，将 "最多 10 个" 更正为 "最多 8 个"。
- 在 `ProfileEditPage` 中增加前端标签数量上限校验（当前仅依赖后端报错）。

---

### 4. 内部文档矛盾：头像是否为必填项

**问题**：`prd/v1-profile.md` 规定 `avatar_url` 为必填字段，但 `docs/design-docs/ui-specs/profile/profile-setup-wireframe.md` Step 1 中写明：
> 头像为选填项，系统提供默认头像

**建议**：产品侧需统一口径。若按 PRD 执行（头像必填），请同步修正 wireframe；若按 wireframe 执行（头像选填），请同步修正 PRD、DTO、Service 校验和 BDD 测试。

---

### 5. 机械检查环境说明

| 检查项 | 结果 | 说明 |
|--------|------|------|
| `make check-docs` | ✅ 通过 | 仅存在 glossary 术语未使用警告，与本次 PR 无关 |
| `make lint`（API） | ✅ 通过 | `eslint` 无报错 |
| `make test`（API 单元测试） | ✅ 通过 | 34/34 通过 |
| `make test-bdd` | ❌ 失败 | 1 个 profile BDD case 失败（详见问题 2） |
| `make check-contract-sync` | ⚠️ 无法执行 | 环境缺少 openapi-generator-cli jar 文件，非代码问题 |
| Flutter analyze / test | ⚠️ 无法执行 | 审查环境未安装 Flutter SDK，非代码问题 |

---

## 总结

本次 PR #16（`feat/issue-3`）在**全链路完整性**和**文档同步**方面做得较好，涵盖了后端 API、数据库、Mobile 页面、BDD 测试和模块文档。但存在以下**阻塞性问题**，必须在合并前修复：

1. **补充 `designs/issue-3/` 设计资产**（`design-spec.md` 及原型文件），否则无法验证前端实现是否符合设计规范。
2. **修复 `avatar_url` 必填校验缺失**导致的 BDD 测试失败，同步更新 DTO、Service、OpenAPI 契约及生成类型。
3. **统一头像必填性的产品定义**，并同步修正 PRD 或 wireframe 中的矛盾描述。

**在以上阻塞性问题解决前，本 PR 不予 LGTM。**
