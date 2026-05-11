## 审查结论

审查结论: **NEEDS_MAJOR_FIX**

本 PR 为 `design/issue-43`（Matching 模块 UI/UX 设计），交付了线框图、高保真原型、设计规范文档，并同步更新了 API 文档、模块文档和 PRD 引用。设计资产本身质量较高，覆盖了 PRD 定义的全部页面和状态。但存在 **4 项 Major 问题**：其中最关键的是 PR 将回滚 `main` 分支上已修复的设计模板，且未按当前模板要求交付 Flutter 纯 UI 代码。

---

## 阻塞问题（Blocking）

无。

---

## 严重问题（Major）

### 1. TEMPLATE-design.md 将回滚 main 分支上的模板修复
- **位置**: `docs/exec-plans/templates/TEMPLATE-design.md`（PR diff 中的整体变更）
- **问题**: `main` 分支在 `ad62d97`（`fix(design-template): 设计任务必须产出 Flutter Widget + Widgetbook + Golden File`）中已明确修正设计模板，要求 Design PR 必须包含 **Flutter Widget 纯 UI 实现 + Widgetbook 用例 + Golden File**。本 PR 基于 `7062b85` 之前的旧状态，其 `TEMPLATE-design.md` 将设计任务降级为仅产出"原型图、设计稿、视觉资产"。若合并，将直接覆盖 `main` 上的修正，导致后续 Agent 误将 Design PR 当作纯美术任务、跳过 UI 代码实现。
- **影响**: 全项目后续 Design PR 的流程标准被错误降级，破坏 `flutter.md` 与 `TEMPLATE-design.md` 之间已对齐的 Design-Only PR 定义。
- **修复**:
  ```bash
  # 从 PR 中移除 TEMPLATE-design.md 的变更，保留 main 上的 ad62d97 修正
  git checkout main -- docs/exec-plans/templates/TEMPLATE-design.md
  ```
- **依据**: `main` 上 `ad62d97` 的提交日志明确说明："设计任务 ≠ 美术任务，设计任务产出的是'可运行的 UI 代码'"；`git diff main..design/issue-43 -- docs/exec-plans/templates/TEMPLATE-design.md` 显示 PR 将回滚该修复。

### 2. 未按当前模板要求交付 Flutter 纯 UI 代码
- **位置**: PR 整体范围
- **问题**: 根据 `main` 上当前生效的 `TEMPLATE-design.md`，Design PR 的范围边界明确包含：
  - `packages/design-system/` 组件
  - `apps/mobile/lib/presentation/` 纯 UI 骨架（无 BLoC/网络）
  - `packages/design-system/widgetbook/main.dart` 用例注册
  - `test/goldens/` 截图基线
  本 PR 的变更文件列表中完全没有 `packages/design-system/` 或 `apps/mobile/lib/presentation/` 的代码文件，唯一的 Mobile 相关变更是空的 `apps/mobile/assets/images/.gitkeep`。
- **影响**: 本 PR 作为 "Design-Only PR" 不完整；后续 feat PR 将缺乏可直接 review 设计方向的 runnable UI 代码，增加返工风险。
- **修复**: 补充以下任一方案：
  - **方案 A（推荐）**: 在本 PR 中补充 Matching 相关 design-system 组件（`SwActionButton`、`SwipeCard`、`SwipeCardStack`、`MatchSuccessDialog`、`MatchListItem`）的纯 UI 实现、Widgetbook 用例注册及 Golden File 截图。
  - **方案 B**: 若本 PR 仅作为"美术预设计"，则应在 PR 标题/描述中明确说明这不是 Design-Only PR，并创建后续独立的 `design/issue-43-flutter` PR 补充代码。
- **依据**: `main:docs/exec-plans/templates/TEMPLATE-design.md` 第 1 行定义范围边界为 "UI/UX 设计 + Flutter Widget 纯 UI 实现"；`main:docs/design-docs/flutter.md` 第 101-109 行同样要求 Design-Only PR 必须包含 UI 代码和 Widgetbook。

### 3. 缺少 `docs/design-docs/ui-specs/INDEX.md`
- **位置**: `docs/design-docs/ui-specs/` 目录
- **问题**: `TEMPLATE-design.md` 的交付前检查清单明确要求："**索引已更新** — `docs/design-docs/ui-specs/INDEX.md`"。该文件在仓库中**不存在**（`main` 和本分支均不存在），且本 PR 新增 `docs/design-docs/ui-specs/matching/` 子目录后仍未创建该索引。目前 `ui-specs/` 下已有 `auth/`、`dev-tools/`、`matching/`、`profile/` 四个子目录，缺乏统一的导航索引。
- **影响**: 后续 Agent 难以快速发现已存在的设计规范，违背"仓库即唯一真实来源"原则。
- **修复**: 在 PR 中新增 `docs/design-docs/ui-specs/INDEX.md`，列出所有功能模块的设计规范入口：
  ```markdown
  # UI 设计规范索引

  | 模块 | 规范文档 | 线框图 |
  |------|----------|--------|
  | auth | [spec.md](auth/login-page-wireframe.md) | — |
  | dev-tools | [spec.md](dev-tools/dev-tools-drawer-wireframe.md) | — |
  | matching | [spec.md](matching/spec.md) | [wireframes/](matching/wireframes/) |
  | profile | — | — |
  ```
- **依据**: `docs/exec-plans/templates/TEMPLATE-design.md` 交付前检查清单最后一项。

### 4. 设计规范文档存在完全重复的两份副本
- **位置**: `designs/issue-43/design-spec.md` 与 `docs/design-docs/ui-specs/matching/spec.md`
- **问题**: 两份文件内容几乎完全一致（仅 Markdown 内部相对链接的深度不同：前者使用 `../../prd/...`，后者使用 `../../../../prd/...`）。这造成维护风险：未来任何设计规范更新（如调整 Token、修改组件路径）必须同时修改两份文件，否则会出现"唯一真实来源"冲突。
- **影响**: 长期维护成本上升，极易出现文档不一致。
- **修复**: 选择以下任一方案：
  - **方案 A（推荐）**: 以 `docs/design-docs/ui-specs/matching/spec.md` 作为唯一规范正本；`designs/issue-43/design-spec.md` 简化为仅包含设计决策摘要和指向正本的链接。
  - **方案 B**: 删除 `docs/design-docs/ui-specs/matching/spec.md`，在 `docs/design-docs/ui-specs/INDEX.md` 中直接链接到 `designs/issue-43/design-spec.md`。
- **依据**: `AGENTS.md` 黄金原则第 3 条："`packages/shared-types/` 是前后端的唯一真实来源"（同理，设计规范也应只有一个真实来源）；`diff` 命令显示两份文件仅 3 行相对路径不同。

---

## 轻微问题（Minor）

### 1. 空文件 `apps/mobile/assets/images/.gitkeep` 无意义
- **位置**: `apps/mobile/assets/images/.gitkeep`
- **问题**: 该文件大小为 0 字节，在 PR 中不产生任何实际价值，且 `assets/images/` 目录目前无其他文件。若后续有真实图片资产，再添加 `.gitkeep` 不迟。
- **修复**: 从 PR 中移除该文件。

### 2. `docs/api/socket-events.md` 被标注为"新增"但早已存在
- **位置**: `designs/issue-43/docs-to-update.txt` 第 8-10 行；`designs/issue-43/design-spec.md` 第 297 行
- **问题**: 文档清单和 design-spec 的"需更新文档"章节均将 `docs/api/socket-events.md` 标注为"新增 Socket.io 事件文档"。但该文件已在 `5e5ead8 design: assets for issue #4 (#33)` 中创建，且当前 `main` 分支上已包含完整的 `match:created` Payload 定义。
- **修复**: 将 docs-to-update.txt 和 design-spec.md 中的措辞从"新增"改为"确认/核对"，或直接移除该项。
- **依据**: `git log --oneline --all -- docs/api/socket-events.md` 显示文件已存在于 `5e5ead8`。

### 3. `prd/v1-matching.md` 与 `docs/api/matching.md` 存在纯格式噪音
- **位置**: `prd/v1-matching.md` 全文；`docs/api/matching.md` 全文
- **问题**: PR 对这两份文档的修改几乎全部为：
  - 在章节标题前增加空行（如 `### 背景` → 前面加空行）
  - 表格列宽对齐（如 `| 参数 | 类型 |` → `| 参数    | 类型      |`）
  这些格式变更不产生语义价值，却增大了 diff 噪音和潜在合并冲突面。
- **修复**: 回滚这两份文件中的纯格式变更，保留真正有意义的修改（如 PRD 顶部的 `designs/issue-43/` 链接更新、API 文档顶部的 `design-spec.md` 链接更新）。
- **依据**: `git diff main..design/issue-43 -- prd/v1-matching.md` 与 `docs/api/matching.md` 显示无内容变更，仅空白与表格对齐。

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 全链路完整性 | ⚠️ N/A | 本 PR 为设计文档 PR，未涉及 DB/API/Mobile 代码实现，不适用全链路代码审查。但 design-spec 中定义的 BLoC/Event/State 与 `docs/modules/matching.md` 保持一致。 |
| 规范符合性 | ❌ 未通过 | TEMPLATE-design.md 将与 main 分支冲突；`ui-specs/INDEX.md` 缺失；设计规范文档重复。 |
| 测试覆盖 | ⚠️ N/A | 设计 PR 无代码，但 design-spec 中完整定义了 Maestro E2E Flow 建议（`.maestro/flows/matching_swipe.yaml` 等），为后续 feat PR 提供了良好的 E2E 覆盖蓝图。 |
| 文档同步 | ⚠️ 部分通过 | PRD、API 文档、模块文档均已正确引用 `designs/issue-43/`。但 `docs/api/socket-events.md` 的"新增"描述不准确，`flutter.md` 未按 design-spec 第 9.2 节要求补充 Matching 组件的 Widgetbook 用例要求（虽然 spec 备注"已存在通用规范"，但 checklist 中仍列出该项）。 |
| PRD 覆盖 | ✅ 通过 | design-spec.md 的页面/场景清单覆盖了 `prd/v1-matching.md` 中定义的全部 Matching 功能：推荐卡片（US-MATCH-1）、滑动操作（US-MATCH-2）、双向匹配弹窗（US-MATCH-3）、匹配列表（US-MATCH-4）。空态、加载态、错误态均已在原型图中定义。 |
| Accessibility | ✅ 通过 | design-spec 第 8 节定义了完整的 Semantics identifier/label，与 `docs/design-docs/flutter.md` 的 Maestro 兼容性要求一致（如 `GestureDetector(onTapDown)`、Stack 中避免 Align 等）。触摸目标尺寸（64×64pt、48×48pt）符合 Material 规范。 |
| 视觉资产 | ✅ 通过 | 交付了 `wireframe.svg`（低保真线框图，4 个页面状态）和 `mockup.html`（高保真可交互原型，含首页/空态/配对弹窗/匹配列表）。 |

---

## 风险与建议

1. **模板回滚风险（高）**：若本 PR 在未移除 `TEMPLATE-design.md` 变更的情况下合并，`main` 分支上已明确的 "Design PR 必须包含 Flutter 代码" 标准将被静默撤销。后续 Agent 可能仅产出美术文档而跳过 UI 实现，重蹈 `flutter.md` 中提到的 Issue #2（LoginPage 未使用 design-system）覆辙。

2. **文档双源风险（中）**：`designs/issue-43/design-spec.md` 与 `docs/design-docs/ui-specs/matching/spec.md` 长期并存会导致规范碎片化。建议在合并前明确单点真实来源。

3. **Design-Only PR 代码缺口（中）**：本 PR 的设计规范非常详尽（BLoC 状态机、组件清单、Animation 曲线、Accessibility 标签均已具备），具备直接转化为 Flutter 纯 UI 代码的条件。建议利用这些已有设计资产，在本 PR 或紧接的下一个 PR 中完成 design-system 组件和纯 UI 骨架，避免设计上下文在后续开发中丢失。

4. **E2E 语义标识符一致性（低）**：design-spec 第 8.2 节和第 10 节已精确定义了 Maestro 所需的 `Semantics identifier`（如 `recommendation_card_top`、`swipe_like_button`、`match_continue_button`）。后续 feat PR 的 Flutter 实现 Agent 必须严格遵循这些 identifier，否则已规划的 `.maestro/flows/*.yaml` 将无法运行。
