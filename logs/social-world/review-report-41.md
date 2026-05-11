# PR #41 代码审查报告

> **PR**: #41 (`design/issue-4`)  
> **标题**: design: add assets for issue #4  
> **审查日期**: 2026-04-18  
> **审查结论**: ❌ 需修改

---

## 1. PR 变更概览

| 文件 | 变更类型 | 说明 |
|------|---------|------|
| `designs/issue-4/design-spec.md` | 修改 | 设计规范更新（+10/-9 行） |
| `designs/issue-4/docs-to-update.txt` | 修改 | 文档更新清单扩展（+8/-5 行） |
| `designs/issue-4/mockup.html` | 修改 | 高保真原型重构（+426/-354 行） |
| `designs/issue-4/preview.dart` | 新增 | Flutter Web 独立预览文件（1351 行） |
| `designs/issue-4/wireframe.svg` | 修改 | 低保真线框图重构（+480/-259 行） |

**变更范围**：仅限 `designs/issue-4/` 目录，未触及 `apps/`、`packages/`、`docs/`（除设计资产目录外）、`.maestro/`。

---

## 2. 历史上下文说明

> ⚠️ 审查前必须明确：**本 PR 并非 issue #4 的首次设计资产提交**。
>
> - PR #33（commit `5e5ead8`）已于 2026-04-16 完成首次设计资产交付，并**同步更新了** `docs/api/matching.md`、`docs/api/socket-events.md`、`docs/modules/matching.md`、`prd/v1-matching.md`、`docs/api/INDEX.md`、`CHANGELOG.md`。
> - PR #41 是在 PR #33 基础上的**设计资产迭代**，核心新增物为 `preview.dart` 及更新的视觉原型（`mockup.html`、`wireframe.svg`）。
>
> 因此，本审查的重点是：**迭代后的设计资产是否补齐了首次交付的缺口，以及是否仍违反项目的设计流程规范。**

---

## 3. 逐条审查问题

### 3.1 ❌ `docs-to-update.txt` 状态与事实不符

**问题描述**：`designs/issue-4/docs-to-update.txt` 使用 "新增"、"更新" 等动词描述以下文档：

```text
docs/api/matching.md|新增 Matching 模块 API 文档...
docs/api/socket-events.md|新增 Socket.io 事件文档...
docs/modules/matching.md|更新 Matching 模块实现文档...
prd/v1-matching.md|顶部添加设计资产链接注释 designs/issue-4/
```

但上述文档**已在 PR #33 中完成新增/更新**（`git show 5e5ead8 --stat` 可证）。`docs-to-update.txt` 作为内部追踪文件，在本 PR 中继续声称需要 "新增" 这些文档，会对后续 Agent 产生误导，使其误以为文档未同步而重复工作。

**修复指令**：
- **文件**: `designs/issue-4/docs-to-update.txt`
- **修改建议**: 将内容更新为反映当前真实状态的清单。例如改为：
  ```text
  # PR #41 已完成的文档同步（由 PR #33 前置完成）
  docs/api/matching.md|已于 PR #33 新增，本 PR 无需修改
  docs/api/socket-events.md|已于 PR #33 新增，本 PR 无需修改
  docs/modules/matching.md|已于 PR #33 新增，本 PR 无需修改
  prd/v1-matching.md|已于 PR #33 添加设计资产链接，本 PR 无需修改
  # 本 PR 待办（如后续 Full-Stack PR 需要）
  docs/design-docs/flutter.md|补充 Matching 组件 Widgetbook 用例执行记录（Full-Stack PR 阶段）
  .maestro/flows/matching/|新增 Matching 用户流程 E2E 测试（Full-Stack PR 阶段）
  ```

---

### 3.2 ❌ 缺失 `docs/design-docs/ui-specs/matching/` 线框图归档文档

**问题描述**：项目现有规范要求每个核心功能在 `docs/design-docs/ui-specs/<feature>/` 下归档 UI 规格文档。现有目录结构：

```
docs/design-docs/ui-specs/
├── auth/login-page-wireframe.md
├── dev-tools/dev-tools-drawer-wireframe.md
└── profile/profile-edit-wireframe.md
    profile-setup-wireframe.md
```

Matching 作为核心用户流程（首页推荐、滑动、配对成功弹窗、匹配列表），**缺失对应的 `ui-specs/matching/` 目录及线框图归档文档**。这导致后续开发 Agent 无法像 auth/dev-tools/profile 一样在统一位置找到 Matching 的 UI 规格。

**修复指令**：
- **文件**: 新建 `docs/design-docs/ui-specs/matching/matching-wireframe.md`
- **修改建议**: 参照 `auth/login-page-wireframe.md` 的格式，产出 Matching 核心页面的低保真线框图及设计系统组件映射，至少覆盖：
  - 首页推荐卡片（HomePage）
  - 空状态（EmptyRecommendationView）
  - 配对成功弹窗（MatchSuccessDialog）
  - 匹配列表页（MatchListPage）

---

### 3.3 ❌ Design-Only PR 未集成实际 design-system 与 Widgetbook

**问题描述**：根据 `docs/design-docs/flutter.md` 的 **Design-Only PR 强制规范**：

> "只修改 `packages/design-system/` 和 `apps/mobile/lib/presentation/` 中的纯 UI 骨架"
> "所有新组件必须在 `widgetbook/main.dart` 中注册"

本 PR 虽然产出了 `preview.dart`（1351 行的独立 Flutter 预览），但：
- **未创建** `packages/design-system/lib/src/atoms/sw_action_button.dart`
- **未创建** `packages/design-system/lib/src/molecules/swipe_card.dart`
- **未创建** `packages/design-system/lib/src/organisms/swipe_card_stack.dart`
- **未创建** `packages/design-system/lib/src/organisms/match_success_dialog.dart`
- **未创建** `packages/design-system/lib/src/molecules/match_list_item.dart`
- **未在** `packages/design-system/widgetbook/main.dart` 中注册任何 Matching 相关用例

`preview.dart` 使用的是本地 `_Tokens` 和私有 Widget（`_ActionButton`、`_SwipeableCard`、`_CardStack` 等），**无法被实际 Mobile App 或 Widgetbook 复用**。这导致本次 Design-Only PR 的产出无法作为后续 Full-Stack PR 的直接依赖。

**修复指令**：
- **文件 1**: `packages/design-system/lib/src/atoms/sw_action_button.dart`
- **文件 2**: `packages/design-system/lib/src/molecules/swipe_card.dart`
- **文件 3**: `packages/design-system/lib/src/organisms/swipe_card_stack.dart`
- **文件 4**: `packages/design-system/lib/src/organisms/match_success_dialog.dart`
- **文件 5**: `packages/design-system/lib/src/molecules/match_list_item.dart`
- **文件 6**: `packages/design-system/widgetbook/main.dart`
- **修改建议**: 
  1. 将 `preview.dart` 中的通用组件提取为 design-system 组件，使用 `DesignTokens`（而非本地 `_Tokens`）
  2. 在 `widgetbook/main.dart` 中为上述 5 个组件注册 use-case，覆盖设计规范中要求的状态（如 `SwActionButton` 的 Pass/Like 变体、`SwipeCard` 的默认态/拖动中态等）
  3. `preview.dart` 可以保留在 `designs/issue-4/` 作为独立预览，但不应替代 Widgetbook 集成

---

### 3.4 ❌ 缺少 Golden File 截图或 Widgetbook 运行截图

**问题描述**：`docs/design-docs/flutter.md` 的 Design-Only PR 验收 Checklist 明确要求：

> - [ ] PR 中附带 Golden File 截图或 Widgetbook 运行截图，供人类直接查看视觉产物

本 PR 的变更均为代码产物（`preview.dart`、`mockup.html`、`wireframe.svg`），**未提交任何可直接在 GitHub 上查看的 PNG/Golden 截图**。虽然 `mockup.html` 可在浏览器打开查看，但 Flutter 组件的真实渲染效果（尤其是动画、手势反馈）无法通过 HTML/SVG 验证。

**修复指令**：
- **方式 A（推荐）**: 运行 `flutter test --update-goldens` 生成 Golden File（`.png`），将生成的图片提交到 `packages/design-system/test/goldens/matching/` 或 `designs/issue-4/screenshots/`，并在 PR 描述中引用。
- **方式 B**: 运行 `flutter run -d chrome -t packages/design-system/widgetbook/main.dart`，对 Matching 相关组件截图，将截图粘贴到 PR 描述中。

---

### 3.5 ❌ 设计规范未覆盖 Accessibility / Maestro E2E 语义要求

**问题描述**：`docs/design-docs/flutter.md` 的 Accessibility 编码规范是**强制性**的，要求：

- 所有需 Maestro 定位的 Widget 必须声明 `Semantics(identifier: 'xxx', label: 'xxx', button: true, container: true)`
- `GestureDetector` 必须使用 `onTapDown` 而非 `onTap`
- `Stack` 中悬浮 Widget 必须使用 `Positioned` + 明确尺寸，**严禁 `Align`**

但 `designs/issue-4/design-spec.md` 在以下方面完全未提及 Accessibility：
- `SwActionButton` 的 `Semantics` 属性（如 `identifier: 'pass_button'` / `'like_button'`）
- `SwipeCard` 拖动时的语义反馈
- `MatchSuccessDialog` 中 CTA 按钮的语义标识
- `MatchListItem` 中聊天入口按钮的语义标识
- 首页模式切换 Tab 的语义标识

同时，`preview.dart` 中全部使用 `GestureDetector(onTap: ...)`，**违反 Maestro 兼容性规范**（`onTap` 无法被 Maestro `tapOn` 触发）。

**修复指令**：
- **文件**: `designs/issue-4/design-spec.md`
- **修改建议**: 在 "交互说明" 或新增 "Accessibility 规范" 章节中，为每个可交互元素指定 `Semantics` 属性。例如：
  ```markdown
  ### Accessibility 标识

  | 元素 | identifier | label | 说明 |
  |------|-----------|-------|------|
  | Pass 按钮 | `pass_button` | `不喜欢` | `button: true, container: true` |
  | Like 按钮 | `like_button` | `喜欢` | `button: true, container: true` |
  | 模式 Tab-附近 | `mode_tab_nearby` | `附近` | `button: true` |
  | 匹配列表项 | `match_list_item_{user_id}` | `{nickname}` | `container: true` |
  | 立即聊天按钮 | `match_chat_button` | `立即聊天` | `button: true` |
  ```
- **文件**: `designs/issue-4/preview.dart`
- **修改建议**: 将 `GestureDetector(onTap: ...)` 统一改为 `GestureDetector(onTapDown: (_) => onTap())`，并为关键交互元素包裹 `Semantics`。

---

### 3.6 ⚠️ 缺少 Matching 用户流程的 Maestro E2E 测试

**问题描述**：`.maestro/flows/` 目前仅包含：

```
.maestro/flows/
├── auth/login.yaml
├── devtools/open_drawer.yaml
├── devtools/edit_server_url.yaml
└── profile/setup.yaml
```

Matching 是核心用户流程（首页 → 滑动卡片 → 配对成功 → 匹配列表），但 **`.maestro/flows/` 下无任何 Matching 相关 E2E 流**。根据 `docs/design-docs/issue-workflow.md`：

> "涉及新的用户流程（登录、注册、资料设置、**新页面交互**）→ **必须**新建 `.maestro/flows/<flow>.yaml`"

**备注**：由于本 PR 是 Design-Only PR（未产出实际可运行的 App 页面），E2E 的缺失在本次审查中记为 **警告（⚠️）** 而非阻断项。但必须在 **后续 Full-Stack PR 中强制补齐**，建议在本 PR 的 `docs-to-update.txt` 中明确标注为待办。

**修复指令（建议在本 PR 或后续立即执行）**：
- **文件 1**: 新建 `.maestro/flows/matching/swipe_cards.yaml`
- **文件 2**: 新建 `.maestro/flows/matching/match_list.yaml`
- **修改建议**: 
  - `swipe_cards.yaml` 覆盖：进入首页 → 等待加载 → 断言卡片可见 → 点击 Like 按钮 → 断言卡片切换
  - `match_list.yaml` 覆盖：点击底部导航「消息」→ 等待加载 → 断言匹配列表可见 → 点击首项 → 断言进入聊天页

---

## 4. 验收 Checklist 逐项核查

根据 `designs/issue-4/design-spec.md` 第 9 节 "验收 Checklist（供 Design-Only PR 使用）"：

| Checklist 项 | 状态 | 说明 |
|-------------|------|------|
| 低保真/高保真原型图已产出 | ✅ | `wireframe.svg` + `mockup.html` + `preview.dart` |
| 所有新增组件已在 `widgetbook/main.dart` 中注册 use-case | ❌ | 未注册任何 Matching 组件 |
| PR 中附带 Golden File 截图或 Widgetbook 运行截图 | ❌ | 仅有代码产物，无截图 |
| `pages/` 中未使用硬编码颜色/原生 Material 组件 | ✅ | 本 PR 未修改 `pages/` |
| 页面中的状态在原型图中有明确定义 | ✅ | Loading、空态、错误、配对成功弹窗均有定义 |

---

## 5. 机械检查

| 检查项 | 命令 | 结果 |
|-------|------|------|
| 文档健康检查 | `make check-docs` | ✅ 通过 |
| Agent 配置检查 | `make agent-check` | ✅ 通过 |
| BDD 测试存在性 | `ls tests/bdd/matching_*.feature` | ✅ 4 个 feature 文件已存在（前置 PR 完成） |

---

## 6. 总结与结论

### 通过项
- 设计资产完整性：`wireframe.svg`、`mockup.html`、`preview.dart` 覆盖了 Matching 核心场景（首页卡片、空状态、配对成功、匹配列表、加载中）。
- 设计规范覆盖：状态、动画、BLoC 事件/状态设计、组件清单均有定义。
- 文档机械检查通过。
- 前置 PR #33 已完成 API 文档、模块文档、PRD 链接同步。

### 阻塞项（必须修复）
1. `docs-to-update.txt` 内容 stale，需更新为反映真实文档状态的追踪清单。
2. 缺失 `docs/design-docs/ui-specs/matching/matching-wireframe.md`，未遵循项目 UI 规格归档规范。
3. Design-Only PR 未产出可复用的 `packages/design-system/` 组件，也未在 `widgetbook/main.dart` 注册。
4. 未提交 Golden File 或 Widgetbook 截图，人类无法在 GitHub 上直接视觉验收。
5. 设计规范完全缺失 Accessibility / Maestro 语义标识要求。

### 建议项（建议本次或下个 PR 补齐）
- `.maestro/flows/matching/` E2E 测试流需在 Full-Stack PR 阶段补齐。
- `preview.dart` 中的 `GestureDetector(onTap)` 应改为 `onTapDown` 以符合 Maestro 兼容性规范。

### 最终结论

**❌ 需修改**

本 PR 作为 Design-Only PR 的迭代，在视觉原型层面有增量价值（尤其是 `preview.dart`），但在项目设计流程合规性上存在明显缺口：未集成 Widgetbook、未产出 design-system 组件、未覆盖 Accessibility 规范、缺少 UI 规格归档文档。建议按上述 5 条修复指令修改后重新提交审查。
