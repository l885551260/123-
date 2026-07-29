# Codex 接入指南

> 在 Codex 桌面客户端中接入本平台模型，实现 AI 编程。

[**Codex**](https://developers.openai.com/codex/) 是 OpenAI 官方的桌面端 AI 编程 Agent，可读取代码库、修改文件、执行命令，支持自定义模型 Provider 接入。

---

## 前置条件

| 条件 | 要求 |
|------|------|
| Codex 客户端 | 已安装（从官方页面下载） |
| API Key | 本平台签发的 `sk-` 前缀密钥 |
| 配置文件 | `~/.codex/config.toml`（不存在则新建） |

---

## 安装 Codex

从 [OpenAI Codex 页面](https://developers.openai.com/codex/) 下载并安装 Codex 桌面客户端。

或通过命令行安装：

```bash
npm install -g @openai/codex
```

验证安装：

```bash
codex --version
```

---

## 配置本平台 API

### 执行步骤

**步骤 1**：打开（或新建）配置文件：

- macOS / Linux：`~/.codex/config.toml`
- Windows：`C:\Users\你的用户名\.codex\config.toml`

**步骤 2**：写入以下内容：

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

将 `sk-你的密钥` 替换为本平台 API Key（[控制台](https://www.aytdai.com) → 创建 API 密钥）。

**步骤 3**：保存文件，重启 Codex。

### 字段说明

| 字段 | 值 | 说明 |
|------|-----|------|
| `model` | `deepseek-v4-pro` | 默认使用的模型标识符 |
| `model_provider` | `aytdai` | 指向下方 `[model_providers.aytdai]` 段定义的 Provider |
| `model_context_window` | `131072` | 模型上下文窗口大小（Token 数） |
| `name` | `AytdAI` | Provider 显示名称 |
| `base_url` | `https://www.aytdai.com/v1` | 本平台 OpenAI 兼容端点 |
| `experimental_bearer_token` | `sk-你的密钥` | 认证密钥，以 Bearer Token 方式发送 |
| `wire_api` | `chat` | API 协议类型，本平台使用 Chat Completions 协议 |

【补充说明：`wire_api` 字段决定 Codex 与后端通信使用的 API 协议。值为 `"chat"` 时使用 OpenAI Chat Completions 协议（`/v1/chat/completions`）；值为 `"responses"` 时使用 OpenAI Responses 协议（`/v1/responses`）。本平台支持 Chat Completions 协议，因此必须填 `"chat"`。】

### 预期结果

重启 Codex 后，即可使用本平台模型进行 AI 编程。

---

## 配置模型能力目录（可选）

Codex 可以通过自定义模型目录识别模型的多模态输入、reasoning effort（thinking 开关）、system prompt、工具类型以及其他详细参数。配置完成后，在 Codex 中输入 `/model`，即可在模型列表中看到模型及其可选 reasoning level。

### 执行步骤

**步骤 1**：在 `~/.codex/config.toml` 中增加一行：

```toml
model_catalog_json = "~/.codex/model-catalogs/custom-catalog.json"
```

**步骤 2**：新建文件 `~/.codex/model-catalogs/custom-catalog.json`，写入模型详细配置：

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
      "slug": "kimi-k2.7-code",
      "display_name": "Kimi K2.7 Code",
      "description": "代码专精，支持图片输入",
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
      "experimental_supported_tools": [],
      "input_modalities": ["text", "image"]
    }
  ]
}
```

**步骤 3**：保存文件，重启 Codex 使模型目录生效。

### 字段含义

| 字段 | 说明 |
|------|------|
| `slug` / `display_name` | 模型在 Codex 配置与 `/model` 列表中的标识和展示名称，`slug` 须与 API 中使用的模型名一致 |
| `default_reasoning_level` | 默认 reasoning effort。任意非 `none` 值开启思考；`none` 关闭思考 |
| `supported_reasoning_levels` | 在 `/model` 中可切换的 reasoning 选项 |
| `base_instructions` | Codex 使用该模型时附加的基础 system prompt，可用于声明模型身份和协作方式 |
| `supports_reasoning_summaries` | 开启 Codex 对该模型的 Responses API reasoning 路径。设为 `true` 后 Codex 才会发送 `reasoning.effort`；否则即使配置了 `default_reasoning_level`，Codex 也会省略 `reasoning` 字段 |
| `default_reasoning_summary` | 设为 `none` 表示不额外请求 reasoning summary |
| `shell_type` | 声明模型适配的 shell 工具调用类型，使用 `shell_command` |
| `visibility` / `supported_in_api` / `priority` | 控制模型是否出现在列表中、是否可通过 API 使用，以及排序优先级 |
| `supports_parallel_tool_calls` | 声明模型支持并行工具调用 |
| `experimental_supported_tools` | 预留的实验性工具能力列表；无额外工具时保持空数组 |
| `input_modalities` | 声明模型支持的输入模态。`["text"]` 为纯文本；`["text", "image"]` 支持文本和图片 |
| `truncation_policy` | 控制上下文截断策略，按字节数限制工具或上下文保留内容 |

---

## 模型切换

修改 `~/.codex/config.toml` 中的 `model` 字段：

```toml
model = "deepseek-v4-flash"
```

重启 Codex 生效。

推荐模型：

| 模型 ID | 适用场景 |
|---------|----------|
| `deepseek-v4-pro` | 编程主力模型（推荐） |
| `deepseek-v4-flash` | 轻量任务、降低成本 |
| `kimi-k2.7-code` | 代码专精、多模态（图片） |

完整可用模型列表见 [首页](/)。

---

## 异常排查

### 连接失败

**排查清单**：

1. `experimental_bearer_token` 是否为 `sk-` 开头的有效密钥
2. `base_url` 是否为 `https://www.aytdai.com/v1`（含 `/v1`）
3. `wire_api` 是否为 `"chat"`
4. 账户余额是否充足

### 报错 "model not found"

**原因**：模型标识符拼写错误。

**处理**：
- 模型名须与本平台 API 接受的 model 参数完全一致（区分大小写）
- 不要添加 `openai/` 或其他前缀，直接写 `deepseek-v4-pro`

### 回答截断

**处理**：调大 `model_context_window` 值，或在模型目录中调整 `truncation_policy`。

### /model 列表中看不到模型

**原因**：未配置模型目录，或 `model_catalog_json` 路径错误。

**处理**：
1. 确认 `config.toml` 中 `model_catalog_json` 路径正确
2. 确认 JSON 文件格式合法（无多余逗号）
3. 确认 `slug` 与 `config.toml` 中 `model` 字段值一致
4. 重启 Codex

---

## 配置回滚

若需恢复使用 OpenAI 官方服务：

1. 删除 `config.toml` 中的 `[model_providers.aytdai]` 段
2. 将 `model_provider` 改为 OpenAI 默认值（或删除该字段）
3. 将 `model` 改为 OpenAI 官方模型名
4. 设置 `OPENAI_API_KEY` 环境变量为 OpenAI 官方密钥
5. 重启 Codex

完整可用模型列表见 [首页](/)。
