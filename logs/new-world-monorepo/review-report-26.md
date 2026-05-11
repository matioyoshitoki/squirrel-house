# PR #26 审查报告：feature/issue-18-ui-infrastructure

**审查日期**: 基于仓库 HEAD  
**审查范围**: `designs/issue-18/` 全部新增文件（19 个文件，+1096 行）  
**关联 Issue**: #18  
**关联 PRD**: PRD-004 (`prd/PRD-004-ui-component-library-redesign.md`，标注关联 Issue #17）  

---

## 审查结论

**审查结论: NEEDS_MAJOR_FIX**

本 PR 提交的设计资产与项目实际技术栈存在**根本性 mismatch**。`apps/phaser3` 是基于 Phaser3 游戏引擎的客户端，全部 UI 通过 Canvas/WebGL 渲染；但 PR 交付的是一套 React + Tailwind CSS + DOM 的 UI 基础设施设计。该设计资产无法被项目使用，也无法迁移到 `apps/phaser3/src/ui/`。此外，所有 `.ts` 文件均为 `declare` 存根（stub），不含可执行实现，且未按 `designs/README.md` 规定的目录结构交付。

---

## 阻塞问题（Blocking）

### 1. 技术栈根本性不匹配：React/Tailwind 设计用于 Phaser3 项目
- **位置**: `designs/issue-18/` 全部文件
- **问题**: 项目 `apps/phaser3` 使用 Phaser 3.86 作为唯一 UI 渲染引擎（`package.json` 依赖 `"phaser": "^3.86.0"`，无 React）。现有组件如 `UIButton.ts`、`UIAvatar.ts` 均继承 `Phaser.GameObjects.Container`，使用 `scene.add.text()`、`scene.tweens.add()`、`hasTexture()` 等 Phaser API。但 PR 交付的设计资产是一套 React DOM 组件基础设施：
  - 多处 `import type * as React from 'react'` (`types/component.ts:11`, `types/event.ts:10`, `utils/dom.ts:10`, `setup/config.ts:11` 等)
  - `React.Context`、`React.FC`、`React.useId`、`React.CSSProperties`
  - `clsx` + `tailwind-merge` 类名合并 (`utils/classnames.ts`)
  - CSS 变量主题切换 via `data-theme` + `document.documentElement` (`setup/theme.ts`)
  - DOM 无障碍工具：`ariaProps`、`focusWithoutScrolling`、`ownerDocument` (`utils/accessibility.ts`, `utils/dom.ts`)
  - JSX 使用示例 (`examples/usage-example.tsx`)
- **影响**: 整套设计资产对项目零价值。Phaser3 使用 Canvas/WebGL 渲染，不存在 DOM、CSS 类名、或 React  reconciler。无法迁移到 `apps/phaser3/src/ui/`，也无法作为独立包运行。
- **修复**: 彻底重写为 Phaser3 原生基础设施：
  - 类型应基于 `Phaser.Scene`、`Phaser.GameObjects.Container`、`Phaser.Types.GameObjects.Text.TextStyle`
  - 工具函数应操作 Phaser GameObjects 和纹理（如 `hasTexture`、`createInkBackground`）
  - 主题系统应使用 `UITheme` 常量对象（已在 `apps/phaser3/src/ui/utils.ts` 中定义）
  - 移除所有 React、DOM、Tailwind 相关设计
- **依据**: 
  - `apps/phaser3/package.json` 第 22 行：`"phaser": "^3.86.0"`，无 react 依赖
  - `apps/phaser3/src/ui/UIButton.ts` 第 25 行：`class UIButton extends Phaser.GameObjects.Container`
  - `apps/phaser3/src/ui/utils.ts`：全部工具基于 `Phaser.Scene` 和 GameObjects API
  - PRD-004 Section 5.1 明确指定目录为 `src/ui/` 且组件继承 `Phaser.GameObjects.Container`

### 2. 完全违背 PRD-004 规范
- **位置**: `designs/issue-18/design-spec.md` 及全部实现文件
- **问题**: PRD-004 明确规划了 Phase 1 基础设施内容：
  - 创建 `src/ui/types.ts` 公共类型（`UIPositionConfig`、`UISizeConfig`、`UIQualityLevel`、`UINotificationType` 等）
  - 重写 `src/ui/utils.ts`（统一 `UITheme`、`UITweens`、`hasTexture`、`createInkBackground`）
  - 建立 `src/ui/__tests__/setup.ts` 测试基础设施
  - 重写 `src/ui/index.ts` 统一导出全部 22 个组件
  但 PR 交付的内容与 PRD-004 的接口定义、类型命名、文件位置完全无关。例如：
  - PRD 要求 `UIPositionConfig` / `UISizeConfig` → PR 交付 `BaseComponentProps` / `AsProp` / `PolymorphicProps`
  - PRD 要求 `UIQualityLevel` / `UINotificationType` → PR 交付 `BrandScale` / `NeutralScale` / `SemanticColorToken`
  - PRD 要求动画工具基于 Phaser tween → PR 交付 `cx` / `tailwind-merge`
- **影响**: 设计资产无法作为 PRD-004 的执行参考，下游开发 Agent 无法按此实现。
- **修复**: 严格按 PRD-004 Section 5.1~5.3 重写类型与工具设计：
  - `types.ts` 应包含 `UIPositionConfig`、`UISizeConfig`、`UIEventCallback`、`UIQualityLevel`、`UINotificationType`
  - `utils.ts` 应强化现有 `UITheme`、`UITweens`、`hasTexture`、`createInkBackground`、`getInkTextStyle`
  - 新增 `__tests__/setup.ts` 设计（MockScene、simulateClick、simulateHover）
- **依据**: PRD-004 Section 5.1（目录结构）、5.2（公共类型定义）、Phase 1 迁移路线图

### 3. 所有 `.ts` 文件为不可执行的 `declare` 存根
- **位置**: 
  - `designs/issue-18/utils/classnames.ts:25`：`export declare function cx(...)`
  - `designs/issue-18/utils/dom.ts:13`：`export declare function canUseDOM()`
  - `designs/issue-18/utils/merge.ts:28`：`export declare function mergeProps<T, U>(...)`
  - `designs/issue-18/utils/accessibility.ts:25`：`export declare function useId(...)`
  - `designs/issue-18/setup/config.ts:42`：`export declare const ConfigProvider: React.FC<...>`
  - `designs/issue-18/setup/install.ts:17`：`export declare function registerComponent(...)`
  - `designs/issue-18/setup/theme.ts:15`：`export declare const ThemeProvider: React.FC<...>`
- **问题**: 全部实现文件使用 `declare` 关键字，没有函数体、没有运行时逻辑。文件无法被 TypeScript 编译为可执行代码（仅能作为 `.d.ts` 类型声明使用）。文件中虽有 "实现提示" 注释，但注释不能替代代码。
- **影响**: 即使技术栈匹配，这些文件也无法运行、无法测试、无法被下游 Agent 直接复制使用。
- **修复**: 将所有 `declare function` / `declare const` 替换为实际实现，或明确将这些文件标记为 `.d.ts` 类型声明文件。若定位为设计文档，应改为 Markdown 规范文档而非 `.ts` 文件。
- **依据**: `designs/README.md` 第 3 步说明设计代码"可直接 copy 到 apps/phaser3/src/"；`declare` 存根不可 copy 即用。

---

## 严重问题（Major）

### 4. 未按 `designs/` 目录规范交付
- **位置**: `designs/issue-18/` 目录结构
- **问题**: `designs/README.md` 明确规定了设计交付物的目录结构：
  ```
  designs/issue-xxx/
  ├── design.md           # 设计文档
  ├── src/                # UI 代码（可直接 copy 到 apps/phaser3/src/）
  │   ├── main.ts         # 入口：导出 init(container) 函数
  │   ├── ui/             # UI 组件代码
  │   ├── scenes/         # Scene 代码
  │   └── mock/           # Mock 数据
  ├── placeholders/       # 占位资源代码
  └── screenshots/        # 截图基线
  ```
  但 PR 实际交付的目录结构为：
  ```
  designs/issue-18/
  ├── design-spec.md      # 设计文档（文件名不符）
  ├── index.ts            # 非 main.ts，无 init(container)
  ├── types/              # React 类型层
  ├── utils/              # React 工具层
  ├── setup/              # React 配置层
  └── examples/           # 示例代码
  ```
  缺少 `src/`、`scenes/`、`mock/`、`placeholders/`、`screenshots/` 目录。
- **影响**: Developer Agent 无法按规范路径找到可迁移代码，设计产物无法直接进入开发阶段。
- **修复**: 按 `designs/README.md` 规范重组目录：
  - 将可迁移的 Phaser3 基础设施代码放入 `designs/issue-18/src/ui/`
  - 创建 `designs/issue-18/src/main.ts` 导出 `init(container)` 演示入口
  - 如有占位资源，放入 `designs/issue-18/placeholders/`
  - 设计文档重命名为 `design.md` 或保留 `design-spec.md` 但需在目录中说明
- **依据**: `designs/README.md` 目录结构章节与使用流程

### 5. 未修改任何实际应用源码
- **位置**: `apps/phaser3/src/ui/`
- **问题**: PR 的 19 个变更文件全部位于 `designs/issue-18/`，未修改 `apps/phaser3/src/ui/utils.ts`、`apps/phaser3/src/ui/index.ts` 等任何实际源码文件。PRD-004 Phase 1 明确要求重写这些文件作为基础设施交付物。
- **影响**: 本次 PR 合并后，项目代码没有任何变化，UI 基础设施重建工作零进展。
- **修复**: 按 PRD-004 Phase 1 直接修改实际源码文件：
  - 修改/重构 `apps/phaser3/src/ui/utils.ts`
  - 新建 `apps/phaser3/src/ui/types.ts`
  - 新建/重构 `apps/phaser3/src/ui/index.ts`
  - 新建 `apps/phaser3/src/ui/__tests__/setup.ts`
  设计文档可作为 `designs/issue-18/design.md` 辅助说明，但不应替代源码变更。
- **依据**: PRD-004 Phase 1 迁移路线图；`git diff --stat` 仅显示 `designs/issue-18/` 变更

### 6. 测试基础设施完全缺失
- **位置**: `apps/phaser3/src/ui/__tests__/`
- **问题**: PRD-004 Section 7.3 和 9.2 明确要求：
  - 新增 `src/ui/__tests__/setup.ts` 提供 MockScene、simulateClick、simulateHover
  - 所有 Layer 1 / Layer 2 组件配套独立测试文件
  - 覆盖率 ≥ 80%
  但 PR 未提供任何测试相关文件。
- **影响**: 无法验证设计资产的可测试性，下游开发 Agent 需要从零搭建测试基础设施。
- **修复**: 
  - 设计阶段至少应产出 `src/ui/__tests__/setup.ts` 的接口设计与实现规划
  - 若设计文档中包含测试策略章节，应明确 MockScene 的接口和 helper 函数签名
- **依据**: PRD-004 Section 7.3（测试要求）、9.2（技术验收）

### 7. 项目文档未同步更新
- **位置**: `docs/frontend/ui-components.md`、`docs/modules/client/ui.md`、`docs/design-docs/ui-style-guide/03-component.md`
- **问题**: PRD-004 Section 9.3 明确要求开发完成后更新：
  - `docs/frontend/ui-components.md`
  - `docs/modules/client/ui.md`
  - `docs/modules/client/INDEX.md`
  - `docs/design-docs/ui-style-guide/03-component.md`
  AGENTS.md 维护规则第 6 条也要求："新增/修改 proto、ent、配置、架构分层时，必须同步更新对应文档"。
  `.doc-coverage.txt` 显示本次 PR 未变更任何文档文件。
- **影响**: 文档与代码脱节，后续 Agent 无法获取最新的 UI 基础设施信息。
- **修复**: 若本次 PR 意图为设计阶段（非源码实现），至少应在设计文档中明确标注：
  - 待更新的文档清单
  - 各文档的变更范围摘要
- **依据**: PRD-004 Section 9.3；`AGENTS.md` 维护规则第 6 条；`.doc-coverage.txt`

---

## 轻微问题（Minor）

### 8. 包名与 monorepo 结构不符
- **位置**: `designs/issue-18/examples/package-exports.json:2`
- **问题**: 文件假设包名为 `@new-world/ui`，但项目 `packages/` 目录下仅有 `proto`，不存在 `@new-world/ui` 包。`apps/phaser3` 是 private 应用包，并非可发布的 UI 库。
- **影响**: 误导后续开发者认为应创建独立 UI 包，与现有 monorepo 结构不符。
- **修复**: 如确需独立 UI 包，应在 `packages/` 下新建并配置 `package.json`；否则示例应使用 `apps/phaser3/src/ui/` 的相对路径导入。
- **依据**: `packages/` 目录实际内容仅 `proto/`

### 9. Issue / PRD 编号不一致
- **位置**: PR 上下文
- **问题**: 本 PR 关联 Issue #18，但提供的 PRD-004 明确标注 "关联 Issue: #17"。`prd/` 目录下也找不到 Issue #18 对应的 PRD。
- **影响**: 审查者无法确认 Issue #18 的需求边界，无法判断设计资产是否覆盖了 Issue #18 的全部要求。
- **修复**: 在 PR 描述中明确 Issue #18 与 PRD-004 的关系（是子任务？独立需求？），或补充 Issue #18 的 PRD。
- **依据**: PRD-004 第 9 行；`find prd/ -name "*18*"` 无结果

---

## 验证结果

| 维度 | 状态 | 说明 |
|------|------|------|
| 功能完整性 | ❌ | PR 未实现任何可执行功能。全部文件为 `declare` 存根，且技术栈错误。 |
| 代码规范性 | ❌ | 目录结构违背 `designs/README.md`；类型命名与 PRD-004 不一致；文件使用 `declare` 而非实际实现。 |
| 全链路完整性 | ❌ | 仅新增设计资产文件，未触及 `apps/phaser3/src/ui/` 任何源码。 |
| 规范符合性 | ❌ | 完全违背 `ui-style-guide/` 5 层规范体系（该规范面向 Phaser3，非 React/DOM）。 |
| 测试覆盖 | ❌ | 无任何测试文件或测试基础设施设计。 |
| 文档同步 | ❌ | 未更新 `docs/frontend/ui-components.md`、`docs/modules/client/ui.md` 等 PRD-004 要求的文档。 |

---

## 风险与建议

### 高风险：下游 Agent 误导
如果本 PR 被合并，后续开发 Agent 可能基于 `designs/issue-18/` 中的 React 设计进行实现，导致在 Phaser3 项目中错误地引入 React、DOM API、CSS 变量等完全不兼容的依赖，产生大量无效工作量。

**建议**：
1. **明确 PR 目标**：先确认 Issue #18 是否就是 PRD-004 的执行任务。如果是，应严格按 PRD-004 的 Phaser3 规范重写。
2. **改为源码实现 PR**：Issue #18 的 UI 基础设施工作应直接修改 `apps/phaser3/src/ui/` 下的实际源码，而非仅在 `designs/` 中编写不可执行的设计文档。
3. **保留现有 Phaser3 工具链**：`apps/phaser3/src/ui/utils.ts` 中已存在 `UITheme`、`UITweens`、`hasTexture`、`createInkBackground`、`getInkTextStyle` 等成熟工具。基础设施重建应在此基础上扩展（如提取 `types.ts`、统一 `index.ts` 导出、建立 `__tests__/setup.ts`），而非推翻重来。
4. **设计资产改为 Markdown 规范**：如果设计阶段只想产出接口规范，建议用 Markdown 文档（如 `design-spec.md`）描述接口，而不是用不可执行的 `.ts` 存根文件冒充源代码。

---

*本审查基于以下证据生成：*
- *`apps/phaser3/package.json` 依赖清单*
- *`apps/phaser3/src/ui/utils.ts` 与 `UIButton.ts` 现有实现*
- *`prd/PRD-004-ui-component-library-redesign.md` 完整规范*
- *`designs/README.md` 设计交付物规范*
- *`AGENTS.md` 项目导航与维护规则*
- *`docs/design-docs/ui-style-guide/01-foundation.md` 视觉规范*
- *`git diff HEAD~1 --stat` 变更统计*
