# Claude Code 接入指南

> 在 Claude Code 中接入本平台模型，进行 AI 编程。

[**Claude Code**](https://github.com/anthropics/claude-code) 是 Anthropic 出品的官方终端原生编程 Agent。本平台提供 Anthropic 兼容 API，可直接作为 Claude Code 的模型后端。

---

## 安装 Claude Code

参考 [Claude Code 官方文档](https://code.claude.com/docs/en/setup) 进行安装。

---

## 配置本平台 API

> **重要提示**：配置前请确保清除以下 Anthropic 官方环境变量，以免影响本平台 API 的正常使用：
>
> ```bash
> unset ANTHROPIC_AUTH_TOKEN
> unset ANTHROPIC_BASE_URL
> ```
>
> 若以上变量在 `~/.bashrc` / `~/.zshrc` 中被永久导出，请同步删除对应行，否则新开 shell 会再次注入。

### 1. 编辑配置文件

编辑或创建 Claude Code 的配置文件：

- **macOS / Linux**：`~/.claude/settings.json`
- **Windows**：`用户目录/.claude/settings.json`

写入以下内容（将 `sk-xxxxxxxxxxxxxxxx` 替换为你的平台 API Key）：

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "https://www.aytdai.com/v1",
    "ANTHROPIC_AUTH_TOKEN": "sk-xxxxxxxxxxxxxxxx",
    "ANTHROPIC_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_OPUS_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_HAIKU_MODEL": "deepseek-v4-flash"
  }
}
```

> **获取 API Key**：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

**模型说明**：

| 环境变量 | 建议值 | 说明 |
|---------|--------|------|
| `ANTHROPIC_MODEL` | `deepseek-v4-pro` | 默认模型，推理能力强 |
| `ANTHROPIC_DEFAULT_SONNET_MODEL` | `deepseek-v4-pro` | Sonnet 级别任务 |
| `ANTHROPIC_DEFAULT_OPUS_MODEL` | `deepseek-v4-pro` | Opus 级别任务 |
| `ANTHROPIC_DEFAULT_HAIKU_MODEL` | `deepseek-v4-flash` | 轻量任务，速度快 |

### 2. 新增 onboarding 标记

编辑或新增 `.claude.json` 文件：

- **macOS / Linux**：`~/.claude.json`
- **Windows**：`用户目录/.claude.json`

写入：

```json
{
  "hasCompletedOnboarding": true
}
```

---

## 启动 Claude Code

配置完成后，进入工作目录，在终端中运行：

```bash
claude
```

启动后选择 **信任此文件夹 (Trust This Folder)**，即可开始使用。

---

## 验证配置生效

启动 `claude` 后，在 TUI 中依次输入以下命令确认配置：

```
/status
/model
```

- `/status` 应显示 `ANTHROPIC_BASE_URL` 指向 `https://www.aytdai.com/v1`
- `/model` 应显示当前模型为 `deepseek-v4-pro`

---

## 可用模型

| 模型 ID | 说明 | 适用场景 |
|---------|------|----------|
| `deepseek-v4-pro` | 更强推理能力（推荐） | 复杂编程、代码生成 |
| `deepseek-v4-flash` | 快速响应，性价比高 | 轻量任务、简单问答 |
| `kimi-k2.7-code` | 代码专精，支持图片输入 | 编程、代码审查 |
| `qwen3.7-max` | 通义千问旗舰 | 复杂推理 |
| `qwen3.7-plus` | 通义千问增强 | 通用任务 |
| `qwen3.7-flash` | 通义千问快速 | 轻量任务 |
| `glm-5.2` | 智谱 GLM | 通用任务 |
| `MiniMax-M2.5` | MiniMax | 通用任务 |

> 注：`kimi/kimi-k3` 和 `MiniMax/MiniMax-M3` 暂不支持 Claude 格式，请使用 OpenAI 格式调用。模型列表以 `GET /v1/models` 为准。

切换模型：在 Claude Code 中运行 `/model` 命令，或直接修改 `settings.json` 中的模型名称。

---

## 常见问题

### 连接失败 / 认证错误

- 确认 `ANTHROPIC_AUTH_TOKEN` 填写的是平台 API Key（`sk-` 开头）
- 确认 `ANTHROPIC_BASE_URL` 为 `https://www.aytdai.com/v1`
- 确认已清除系统环境变量中的 Anthropic 官方配置
- 确认账户有可用额度（登录控制台查看）

### 响应中断 / 不完整

- DeepSeek 模型带有思考过程（thinking），会消耗输出 token
- 如遇截断，可在 `/config` 中调大 max_tokens

### 思考模式开关

运行 `/config`，将 **Thinking mode** 设为 `true` 或 `false`。也可随时通过 `Option+T`（macOS）或 `Alt+T`（Windows/Linux）快捷切换。

---

## 费用说明

- 按 **Token** 用量计费，输入和输出分别计价
- 新用户注册即送 **¥1 体验额度**
- 详见 [定价与充值说明](https://www.aytdai.com/docs/#/02-定价与充值说明)
