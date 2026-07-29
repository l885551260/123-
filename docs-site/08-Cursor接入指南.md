# Cursor 接入指南

> 在 Cursor 中接入本平台模型，实现 AI 编程辅助。

[**Cursor**](https://cursor.com) 是 Anysphere 出品的 AI-first IDE，基于 VS Code 二次开发，内置 Agent 与代码库理解能力。通过配置自定义 OpenAI 兼容端点，可将 Chat / Composer / Edit 模式的底层模型替换为本平台提供的模型。

---

## 前置条件

| 条件 | 要求 |
|------|------|
| Cursor 订阅 | **高级会员（Pro）及以上**。免费用户无法配置自定义模型 |
| API Key | 本平台签发的 `sk-` 前缀密钥 |
| 环境变量 | 需清除 `OPENAI_API_KEY` 和 `OPENAI_BASE_URL`（若存在） |

> **⚠️ 重要提示：订阅要求**
>
> 若非 Cursor 高级会员，配置自定义模型时将出现以下错误：
>
> `The model deepseek-v4-pro does not work with your current plan or api key`
>
> 此时需升级 Cursor 订阅至 Pro 或更高级别。

---

## 安装 Cursor

1. 通过 [Cursor 官网](https://cursor.com/) 下载并安装 Cursor
2. 打开 Cursor，点击右上角"设置"按钮，进入设置界面
3. 点击"Sign in"按钮，登录 Cursor 账户

---

## 配置本平台 API

### 前置操作：清除 OpenAI 环境变量

> **⚠️ 重要提示**
>
> 配置前，必须确保清除以下 OpenAI 相关环境变量，以免干扰自定义端点的正常使用：
>
> - `OPENAI_API_KEY`
> - `OPENAI_BASE_URL`
>
> 若上述变量在系统环境变量或 shell 配置文件（`~/.bashrc`、`~/.zshrc`）中被永久设置，需同步删除。

### 已知限制："Override OpenAI Base URL" 为全局设置

> **⚠️ 重要提示**
>
> 开启 "Override OpenAI Base URL" 后，该设置将作用于 Cursor 中**所有**已配置的 API Key——包括 Cursor 自带模型使用的 Anthropic / GPT Key。
>
> 若开启 Override 后 Cursor 自带的 Claude / GPT 模型停止工作，需在不使用本平台模型时关闭 Override。Cursor 目前不支持按模型设置不同 Base URL。
>
> 【补充说明：此为 Cursor 产品层面的架构限制，非本平台或任何第三方 API 的问题。Cursor 社区已有相关讨论及 Feature Request，但截至目前官方尚未提供按模型独立配置 Base URL 的能力。】

### 执行步骤

**步骤 1**：点击左侧栏的 **"Models"**，进入模型配置页面。

**步骤 2**：展开 **"API Keys"** 部分，配置 API 信息：

- 勾选 **"Override OpenAI Base URL"**
- 在下方输入框填入本平台的调用地址：`https://www.aytdai.com/v1`

**步骤 3**：在 **OpenAI API Key** 输入框，填入本平台的 API Key（`sk-` 前缀）。

> API Key 获取方式：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

**步骤 4**：点击 **"OpenAI API Key"** 栏右侧的验证按钮。

**步骤 5**：在弹出的窗口中点击 **"Enable OpenAI API Key"** 按钮，完成设置验证。

**步骤 6**：在 **Models** 板块中，点击 **"View All Models"** 按钮，再点击 **"Add Custom Model"** 按钮。

**步骤 7**：输入模型名称，点击 **"Add"** 按钮。

推荐添加以下模型（每个模型需单独添加一次）：

| 模型名称（严格区分大小写） | 适用场景 |
|---------------------------|----------|
| `deepseek-v4-pro` | 编程主力模型（推荐） |
| `deepseek-v4-flash` | 轻量任务、降低成本 |
| `kimi-k2.7-code` | 代码专精、多模态 |

> **⚠️ 注意**：模型名称须与本平台 API 实际接受的 model 参数完全一致（区分大小写），否则请求将返回模型不存在错误。

**步骤 8**：启用刚添加的模型（确保模型列表中对应开关为开启状态）。

**步骤 9**：在聊天面板顶部的模型下拉菜单中，选择已添加的模型（如 `deepseek-v4-pro`），开始使用。

---

## 预期结果

配置完成后：

- 在 **Chat**（`Ctrl+L`）中可选择自定义模型进行对话
- 在 **Composer** 中可使用自定义模型进行多文件编辑
- 在 **Edit**（`Ctrl+K`）中可使用自定义模型进行代码修改

---

## 功能限制

### Tab 自动补全不受自定义模型影响

Cursor Tab 自动补全是 Cursor Pro 自带能力，始终使用 Cursor 自家模型，不受自定义 API Key 影响。Custom Model 仅在 Chat / Composer / Edit 等模式下生效。

【补充说明：此为 Cursor 产品设计决定。Tab 补全的推理路径与 Chat/Composer 完全独立，不经过用户配置的 OpenAI API Key 通道。】

### Override 的全局影响

如前文所述，开启 Override 后所有 API Key 均受影响。若需同时使用 Cursor 自带模型和本平台模型，需手动切换 Override 开关。

---

## 异常排查

### 模型没有返回任何内容

**处理**：将 Cursor 中的 **"Network"** 设置更改为 **HTTP/1.0**。

路径：Settings → Network → 改为 HTTP/1.0。

【补充说明：此问题通常与 HTTP/2 连接复用或流式传输兼容性有关。切换为 HTTP/1.0 可规避部分网络环境下的连接异常。】

### 报错 "does not work with your current plan or api key"

**原因**：Cursor 订阅级别不足。

**处理**：升级 Cursor 订阅至 Pro 或更高级别。

### 验证失败 / 红色报错

**排查清单**：

1. API Key 是否以 `sk-` 开头，无多余空格或换行
2. Override OpenAI Base URL 是否为 `https://www.aytdai.com/v1`（含 `/v1`，末尾无斜杠）
3. 系统环境中是否存在残留的 `OPENAI_API_KEY` 或 `OPENAI_BASE_URL` 覆盖了 Cursor 设置
4. 账户余额是否充足
5. 网络是否可正常访问 `https://www.aytdai.com`

### Cursor 自带模型停止工作

**原因**：Override OpenAI Base URL 为全局设置，开启后 Cursor 自带的 Anthropic / GPT 模型也会尝试通过自定义 Base URL 发送请求。

**处理**：在不使用本平台模型时，取消勾选 "Override OpenAI Base URL"。

### 回答截断

**原因**：DeepSeek 系列模型的思考过程（thinking tokens）计入输出 Token，导致实际回答内容被截断。

**处理**：在 Settings → Models 中调大 Max Tokens（建议 8192+）。

---

## 配置回滚

若需恢复 Cursor 默认行为（使用 Cursor 自带模型）：

1. 取消勾选 **"Override OpenAI Base URL"**
2. 清空 **OpenAI API Key** 输入框
3. 在 Models 列表中禁用或删除已添加的自定义模型
4. 重启 Cursor

完整可用模型列表见 [首页](/)。
