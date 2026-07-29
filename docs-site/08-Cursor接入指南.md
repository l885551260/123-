# Cursor 接入指南

> 在 Cursor 中接入本平台模型，进行 AI 编程。

[**Cursor**](https://cursor.com) 是 Anysphere 出品的 AI-first IDE，基于 VS Code 二次开发，内置 Agent 与代码库理解能力。本平台提供 OpenAI 兼容 API，可通过自定义模型方式接入。

---

## 安装 Cursor

1. 通过 [Cursor 官网](https://cursor.com/) 下载并安装 Cursor
2. 打开 Cursor，点击右上角"设置"按钮，进入设置界面，点击"Sign in"登录你的 Cursor 账户

---

## 配置本平台 API

> **重要提示：使用前请先清除 OpenAI 环境变量**
>
> 请确保清除以下环境变量，以免影响本平台 API 的正常使用：
>
> - `OPENAI_API_KEY`
> - `OPENAI_BASE_URL`

> **注意：Cursor 仅支持订阅高级会员（Pro）及以上的用户配置自定义模型**
>
> 若非 Cursor 高级会员，配置时将出现以下错误：
>
> `The model deepseek-v4-pro does not work with your current plan or api key`

> **已知问题："Override OpenAI Base URL" 是全局设置**
>
> 开启后会作用到 Cursor 中**所有**已配置的 API Key——包括 Cursor 自带模型用的 Key。如果开启 Override 后 Cursor 自带的 Claude / GPT 模型停止工作，请在不使用本平台模型时关闭 Override。

### 配置步骤

1. 点击左侧栏的 **"Models"**，进入模型配置页面

2. 展开 **"API Keys"** 部分，配置 API 信息：
   - 勾选 **"Override OpenAI Base URL"**
   - 在下方输入本平台的调用地址：`https://www.aytdai.com/v1`

3. 在 **OpenAI API Key** 输入框，输入你的平台 API Key（`sk-` 开头）

   > **获取 API Key**：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

4. 点击 **"OpenAI API Key"** 栏右侧的验证按钮

5. 在弹出的窗口中点击 **"Enable OpenAI API Key"**，完成设置验证

6. 在 **Models** 板块中，点击 **"View All Models"** → **"Add Custom Model"**

7. 输入模型名称，点击"Add"按钮：

   | 模型 ID | 说明 |
   |---------|------|
   | `deepseek-v4-pro` | 更强推理能力（推荐） |
   | `deepseek-v4-flash` | 快速响应，性价比高 |
   | `kimi-k2.7-code` | 代码专精，支持图片输入 |
   | `kimi/kimi-k3` | Kimi 旗舰 |
   | `qwen3.7-max` | 通义千问旗舰 |
   | `qwen3.7-plus` | 通义千问增强 |
   | `qwen3.7-flash` | 通义千问快速 |
   | `glm-5.2` | 智谱 GLM |
   | `MiniMax-M2.5` | MiniMax |
   | `MiniMax/MiniMax-M3` | MiniMax M3 |

   > 模型名称需严格输入完整 ID（注意大小写）。

8. 启用刚添加的模型

9. 在聊天面板中选择对应模型，开始使用

---

## 注意事项

- **Cursor Tab 自动补全无法由自定义模型驱动。** Tab 是 Cursor Pro 自带能力，始终使用 Cursor 自家模型。自定义模型仅在 Chat / Composer / Edit 等模式下生效。
- 如果模型没有返回任何内容，可尝试将 Cursor 中的 **"Network"** 设置更改为 **HTTP/1.0** 解决。

---

## 常见问题

### 验证 API Key 失败

- 确认 API Key 填写正确（`sk-` 开头）
- 确认 Base URL 为 `https://www.aytdai.com/v1`（注意末尾 `/v1`）
- 确认账户有可用额度（登录控制台查看）
- 确认已清除系统环境变量中的 OpenAI 配置

### 响应中断 / 不完整

- DeepSeek 模型带有思考过程（thinking），会消耗输出 token
- 如遇截断，可尝试调大 max_tokens（建议 4096+）

---

## 费用说明

- 按 **Token** 用量计费，输入和输出分别计价
- 新用户注册即送 **¥1 体验额度**
- 详见 [定价与充值说明](https://www.aytdai.com/docs/#/02-定价与充值说明)
