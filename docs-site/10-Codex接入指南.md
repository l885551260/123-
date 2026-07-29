# Codex 接入指南

[Codex](https://developers.openai.com/codex/) 是 OpenAI 官方的桌面 AI 编程 Agent，能读代码、改文件、跑命令。

---

## 第一步：安装

从 [OpenAI Codex 页面](https://developers.openai.com/codex/) 下载安装包，双击装好。

或者用命令行：
```bash
npm install -g @openai/codex
```

装完验证：
```bash
codex --version
```

---

## 第二步：配置接入我们平台

编辑配置文件（没有就新建）：
- macOS / Linux：`~/.codex/config.toml`
- Windows：`C:\Users\你的用户名\.codex\config.toml`

写入：

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

保存后重启 Codex 即可。

**字段说明**：
| 字段 | 含义 |
|------|------|
| `model` | 默认使用的模型名 |
| `model_provider` | 指向下面定义的 provider 名称 |
| `base_url` | 我们平台的 API 地址 |
| `experimental_bearer_token` | 你的 API Key |
| `wire_api` | 必须填 `"chat"`（走 Chat Completions 协议） |

---

## 第三步：使用

```bash
codex
```

进入后直接输入任务描述就行。Codex 会自己读代码、改文件、跑命令。

---

## 切换模型

改 `config.toml` 里的 `model = "模型名"` 就行：

| 模型 | 适合 |
|------|------|
| `deepseek-v4-pro` | 写代码首选（推荐） |
| `deepseek-v4-flash` | 简单任务，省钱 |
| `kimi-k2.7-code` | 代码专精 |

完整列表看 [首页](/)。

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
- `base_url` 是不是 `https://www.aytdai.com/v1`（有 /v1）
- `wire_api` 是不是 `"chat"`
- 余额够不够

**回答截断？** 调大 `model_context_window` 或在请求中设 max_tokens。

**报错 "model not found"？**
- 模型名拼写要和平台上的完全一致（区分大小写）
- 别加 `openai/` 前缀，直接写 `deepseek-v4-pro`
