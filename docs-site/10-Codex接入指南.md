# Codex 接入指南

> 在 Codex 桌面客户端中接入本平台模型，进行 AI 编程。

[**Codex**](https://developers.openai.com/codex/) 是 OpenAI 官方的桌面端 AI 编程 Agent。本平台提供 OpenAI 兼容 API，可通过自定义 provider 方式接入。

---

## 安装 Codex

从 [OpenAI Codex 页面](https://developers.openai.com/codex/) 下载并安装 Codex 桌面客户端。

---

## 配置本平台 API

### 1. 编辑配置文件

打开 `~/.codex/config.toml`（Windows 为 `用户目录/.codex/config.toml`），加入以下内容，并将 `sk-xxxxxxxxxxxxxxxx` 替换为你的平台 API Key：

```toml
model = "deepseek-v4-pro"
model_provider = "aytdai"
model_context_window = 131072

[model_providers.aytdai]
name = "AytdAI"
base_url = "https://www.aytdai.com/v1"
experimental_bearer_token = "sk-xxxxxxxxxxxxxxxx"
wire_api = "chat"
```

> **获取 API Key**：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

### 2. 重启 Codex

重启 Codex，即可开始使用本平台模型。

---

## 可用模型

| 模型 ID | 说明 | 适用场景 |
|---------|------|----------|
| `deepseek-v4-pro` | 更强推理能力（推荐） | 复杂编程、代码生成 |
| `deepseek-v4-flash` | 快速响应，性价比高 | 轻量任务、简单问答 |
| `kimi-k2.7-code` | 代码专精，支持图片输入 | 编程、代码审查 |
| `kimi/kimi-k3` | Kimi 旗舰 | 综合任务 |
| `qwen3.7-max` | 通义千问旗舰 | 复杂推理 |
| `qwen3.7-plus` | 通义千问增强 | 通用任务 |
| `qwen3.7-flash` | 通义千问快速 | 轻量任务 |
| `glm-5.2` | 智谱 GLM | 通用任务 |
| `MiniMax-M2.5` | MiniMax | 通用任务 |
| `MiniMax/MiniMax-M3` | MiniMax M3 | 通用任务 |

切换模型：修改 `config.toml` 中的 `model = "模型ID"` 即可。

---

## 配置模型能力目录（可选）

Codex 可以通过自定义模型目录识别模型的多模态输入、reasoning effort（thinking 开关）等详细参数。配置完成后，在 Codex CLI 中输入 `/model`，即可在模型列表中看到对应模型。

在 `~/.codex/config.toml` 中增加一行：

```toml
model_catalog_json = "~/.codex/model-catalogs/custom-catalog.json"
```

然后新建 `~/.codex/model-catalogs/custom-catalog.json`，写入模型详细配置：

```json
{
  "models": [
    {
      "slug": "deepseek-v4-pro",
      "display_name": "DeepSeek V4 Pro",
      "description": "强推理编程模型",
      "default_reasoning_level": "high",
      "supported_reasoning_levels": [
        { "effort": "none", "description": "Think-Off" },
        { "effort": "high", "description": "Deep" }
      ],
      "shell_type": "shell_command",
      "visibility": "list",
      "supported_in_api": true,
      "priority": 0,
      "base_instructions": "You are Codex, a coding agent powered by DeepSeek V4 Pro. You and the user share the same workspace and collaborate to achieve the user's goals.",
      "supports_reasoning_summaries": true,
      "default_reasoning_summary": "none",
      "support_verbosity": false,
      "truncation_policy": { "mode": "bytes", "limit": 10000 },
      "supports_parallel_tool_calls": true,
      "experimental_supported_tools": [],
      "input_modalities": ["text"]
    },
    {
      "slug": "deepseek-v4-flash",
      "display_name": "DeepSeek V4 Flash",
      "description": "快速轻量模型",
      "default_reasoning_level": "high",
      "supported_reasoning_levels": [
        { "effort": "none", "description": "Think-Off" },
        { "effort": "high", "description": "Deep" }
      ],
      "shell_type": "shell_command",
      "visibility": "list",
      "supported_in_api": true,
      "priority": 1,
      "base_instructions": "You are Codex, a coding agent powered by DeepSeek V4 Flash. You and the user share the same workspace and collaborate to achieve the user's goals.",
      "supports_reasoning_summaries": true,
      "default_reasoning_summary": "none",
      "support_verbosity": false,
      "truncation_policy": { "mode": "bytes", "limit": 10000 },
      "supports_parallel_tool_calls": true,
      "experimental_supported_tools": [],
      "input_modalities": ["text"]
    },
    {
      "slug": "kimi-k2.7-code",
      "display_name": "Kimi K2.7 Code",
      "description": "代码专精模型",
      "default_reasoning_level": "high",
      "supported_reasoning_levels": [
        { "effort": "none", "description": "Think-Off" },
        { "effort": "high", "description": "Deep" }
      ],
      "shell_type": "shell_command",
      "visibility": "list",
      "supported_in_api": true,
      "priority": 2,
      "base_instructions": "You are Codex, a coding agent powered by Kimi K2.7 Code. You and the user share the same workspace and collaborate to achieve the user's goals.",
      "supports_reasoning_summaries": true,
      "default_reasoning_summary": "none",
      "support_verbosity": false,
      "truncation_policy": { "mode": "bytes", "limit": 10000 },
      "supports_parallel_tool_calls": true,
      "experimental_supported_tools": [],
      "input_modalities": ["text", "image"]
    }
  ]
}
```

修改后重启 Codex，使新的模型目录生效。

**常用字段说明**：

- `slug` / `display_name`：模型在 `/model` 列表中的标识和展示名称，需与 API 中的模型名一致
- `default_reasoning_level`：默认 reasoning effort，`none` 关闭思考，其他值开启
- `supported_reasoning_levels`：在 `/model` 中可切换的 reasoning 选项
- `base_instructions`：使用该模型时附加的基础 system prompt
- `input_modalities`：模型支持的输入模态（text / image）

---

## 常见问题

### 连接失败 / 认证错误

- 确认 `experimental_bearer_token` 填写的是平台 API Key（`sk-` 开头）
- 确认 `base_url` 为 `https://www.aytdai.com/v1`
- 确认账户有可用额度（登录控制台查看）

### 响应中断 / 不完整

- DeepSeek 模型带有思考过程（thinking），会消耗输出 token
- 如遇截断，可尝试调大 `model_context_window` 或在请求中设置更大的 max_tokens

---

## 费用说明

- 按 **Token** 用量计费，输入和输出分别计价
- 新用户注册即送 **¥1 体验额度**
- 详见 [定价与充值说明](https://www.aytdai.com/docs/#/02-定价与充值说明)
