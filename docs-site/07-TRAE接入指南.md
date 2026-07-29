# TRAE 接入指南

> 在 TRAE 中接入本平台模型，实现 AI 编程辅助。

[**TRAE**](https://www.trae.ai) 是字节跳动出品的 AI-native IDE，主打智能协作与 Agent 自动化。支持通过 API Key 接入自定义 OpenAI 兼容模型服务。

---

## 安装 TRAE

### 步骤 1：下载并安装

- 国际版：访问 [trae.ai](https://www.trae.ai) 下载（macOS / Windows / Linux）
- 中国版：访问 [trae.cn](https://www.trae.cn) 下载（macOS / Windows）

【补充说明：两个版本均支持自定义模型配置，操作流程一致。中国版内置部分国内模型预设（如 MiniMax-CN），国际版内置 GPT / Claude / Gemini 等预设。本平台无预设服务商，需通过 OpenAI 兼容方式手动添加。】

### 步骤 2：完成初始设置

TRAE 首次启动时将进入初始设置页面，根据指引完成基础配置。

### 步骤 3：登录 TRAE

登录 TRAE 账户。

---

## 配置本平台自定义模型

### 执行步骤

**步骤 1**：在 AI 对话框右上角，点击 **设置** 图标。

**步骤 2**：选择 **模型** 页签。

**步骤 3**：点击 **+ 添加模型** 按钮。

**步骤 4**：服务商选择 **OpenAI**，模型选择 **Custom Model**（自定义模型，位于下拉列表末尾）。

【补充说明：MiniMax 等部分服务商在 TRAE 中拥有预设服务商入口（如 MiniMax-CN），可直接选用。本平台暂无 TRAE 预设服务商，因此需通过 OpenAI 兼容协议 + Custom Model 方式接入。】

**步骤 5**：填写配置信息：

| 字段 | 填写内容 | 说明 |
|------|----------|------|
| Model ID | `deepseek-v4-pro` | 模型标识符，须与本平台 API 接受的 model 参数完全一致 |
| API Key | `sk-你的密钥` | 本平台 API Key（[控制台](https://www.aytdai.com) → 创建 API 密钥） |
| Custom Request URL | `https://www.aytdai.com/v1/chat/completions` | 完整的 Chat Completions 端点路径 |

> **⚠️ 重要提示：Custom Request URL 必须包含完整路径**
>
> 自 TRAE v3.3.51 起，Custom Request URL 字段为**原样使用**，TRAE 不再自动追加 `/chat/completions` 路径。
>
> - ✅ 正确：`https://www.aytdai.com/v1/chat/completions`
> - ❌ 错误：`https://www.aytdai.com/v1`（将导致 404）
> - ❌ 错误：`https://www.aytdai.com`（将导致 404）

**步骤 6**：点击 **添加模型** 按钮。

### 预期结果

TRAE 将调用服务商的接口来检测 API Key 是否有效。可能的结果如下：

- **连接成功**：该自定义模型被添加至模型列表，可在对话中选用
- **连接失败**：添加模型窗口中展示错误信息和服务商返回的错误日志，可参考这些信息排查问题

---

## 使用

添加成功后，回到编辑器，在 AI 对话框顶部的**模型下拉菜单**中选择已添加的模型（如 `deepseek-v4-pro`）。

TRAE 提供三种 AI 交互模式：

| 模式 | 说明 | 触发方式 |
|------|------|----------|
| Builder | Agent 模式，AI 直接读写文件、执行命令 | 对话框切换 |
| Chat | 侧边栏问答 | 对话框切换 |
| Inline Chat | 编辑器内联对话 | `Ctrl+I` |

---

## 添加更多模型

每个模型需单独添加一条配置（重复上述步骤 3-6，仅更改 Model ID）：

| Model ID | 适用场景 |
|----------|----------|
| `deepseek-v4-pro` | 编程主力模型（推荐） |
| `deepseek-v4-flash` | 轻量任务、降低成本 |
| `kimi-k2.7-code` | 代码专精、多模态 |
| `qwen3.7-plus` | 日常对话 |

完整可用模型列表见 [首页](/)。

---

## 异常排查

### 添加模型时连接失败 / 404

**首要排查项**：Custom Request URL 是否包含完整路径。

- 必须为 `https://www.aytdai.com/v1/chat/completions`
- 不能只写到 `/v1` 或裸域名

**其他排查项**：
1. API Key 是否以 `sk-` 开头，无多余空格
2. 账户余额是否充足
3. 网络是否可正常访问 `https://www.aytdai.com`

### Builder 模式卡住 / 工具调用失败

**可能原因**：部分模型在 Agent 模式下对 tool calling 的支持不够稳定。

**处理**：
1. 切换至 `deepseek-v4-pro`（Agent 模式下最稳定）
2. 避免使用过小的模型（如 flash 系列）执行复杂多文件任务
3. 检查上下文长度是否超限

### 回答截断

**原因**：DeepSeek 系列模型的思考过程（thinking tokens）计入输出 Token。

**处理**：在设置中调大 max_tokens（建议 4096+）。

### 找不到 Custom Model 选项

**处理**：
1. 确认 TRAE 为最新版本（Help → Check for Updates）
2. 确认服务商选择的是 **OpenAI**（非其他预设服务商）
3. 在 Model 下拉列表中滚动至最底部，Custom Model 为最后一项

---

## 配置回滚

若需移除本平台模型：

1. 打开设置 → 模型页签
2. 找到已添加的自定义模型
3. 点击对应模型的删除/禁用按钮

完整可用模型列表见 [首页](/)。
