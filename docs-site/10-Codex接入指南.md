# Codex 接入指南

[Codex](https://developers.openai.com/codex/) 是 OpenAI 官方的桌面 AI 编程 Agent。

---

## 安装

从 [OpenAI Codex 页面](https://developers.openai.com/codex/) 下载安装。

---

## 配置

编辑 `~/.codex/config.toml`（Windows 为 `用户目录/.codex/config.toml`）：

```toml
model = "deepseek-v4-pro"
model_provider = "aytdai"
model_context_window = 131072

[model_providers.aytdai]
name = "AytdAI"
base_url = "https://www.aytdai.com/v1"
experimental_bearer_token = "sk-你的密钥"
wire_api = "chat"
```

> 把 `sk-你的密钥` 换成你的 API Key（[控制台](https://www.aytdai.com) → 创建API密钥）。

重启 Codex 即可。

---

## 切换模型

改 `config.toml` 里的 `model = "模型名"` 就行。

推荐：`deepseek-v4-pro`（写代码）、`deepseek-v4-flash`（省钱）。完整列表看 [首页](/)。

---

## 高级：自定义模型目录（可选）

如果你想让 Codex 的 `/model` 列表显示模型的详细信息（思考开关、多模态支持等），可以配置模型目录。

在 `config.toml` 加一行：

```toml
model_catalog_json = "~/.codex/model-catalogs/custom-catalog.json"
```

然后创建 `~/.codex/model-catalogs/custom-catalog.json`：

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
      "supports_parallel_tool_calls": true,
      "input_modalities": ["text"]
    },
    {
      "slug": "kimi-k2.7-code",
      "display_name": "Kimi K2.7 Code",
      "description": "代码专精，支持图片",
      "default_reasoning_level": "high",
      "supported_reasoning_levels": [
        { "effort": "none", "description": "Think-Off" },
        { "effort": "high", "description": "Deep" }
      ],
      "shell_type": "shell_command",
      "visibility": "list",
      "supported_in_api": true,
      "priority": 1,
      "supports_parallel_tool_calls": true,
      "input_modalities": ["text", "image"]
    }
  ]
}
```

重启 Codex 生效。

---

## 常见问题

**连不上？**
- `experimental_bearer_token` 是不是你的 `sk-xxx` 密钥
- `base_url` 是不是 `https://www.aytdai.com/v1`
- 余额够不够

**回答截断？** 调大 `model_context_window` 或 max_tokens。
