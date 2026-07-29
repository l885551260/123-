# Claude Code 接入指南

> 在 Claude Code 中接入本平台模型，实现终端 AI 编程。

[**Claude Code**](https://github.com/anthropics/claude-code) 是 Anthropic 出品的终端原生编程 Agent，可读取代码库、执行命令、进行多文件编辑。通过配置自定义端点，可将其底层模型替换为本平台提供的模型。

---

## 安装 Claude Code

参考 [Claude Code 官方文档](https://code.claude.com/docs/en/setup) 完成安装。

**快速安装命令**：

macOS / Linux：
```bash
curl -fsSL https://claude.ai/install.sh | bash
```

Windows（PowerShell）：
```powershell
irm https://claude.ai/install.ps1 | iex
```

或通过 npm（需 Node.js 22+）：
```bash
npm install -g @anthropic-ai/claude-code@latest
```

安装完成后执行 `claude --version` 确认安装成功。

---

## 配置本平台 API

> **⚠️ 重要提示**
>
> 配置前，必须确保清除以下 Anthropic 相关环境变量，否则其优先级高于配置文件，将导致请求被路由至 Anthropic 官方端点：
>
> - `ANTHROPIC_AUTH_TOKEN`
> - `ANTHROPIC_BASE_URL`
>
> ```bash
> unset ANTHROPIC_AUTH_TOKEN
> unset ANTHROPIC_BASE_URL
> ```
>
> 若以上变量在 `~/.bashrc` / `~/.zshrc` 中被永久导出，需同步删除对应行，否则新开 shell 会再次注入。
>
> Windows PowerShell 用户：
> ```powershell
> $env:ANTHROPIC_AUTH_TOKEN = ""
> $env:ANTHROPIC_BASE_URL = ""
> ```
> 若在系统环境变量中永久设置，需前往「系统属性 → 环境变量」中手动删除。

---

### 步骤一：编辑配置文件

编辑或创建 Claude Code 的配置文件：

- macOS / Linux：`~/.claude/settings.json`
- Windows：`C:\Users\你的用户名\.claude\settings.json`

写入以下内容（将 `sk-你的密钥` 替换为你在本平台创建的 API Key）：

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "https://www.aytdai.com/v1",
    "ANTHROPIC_AUTH_TOKEN": "sk-你的密钥",
    "ANTHROPIC_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_OPUS_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_HAIKU_MODEL": "deepseek-v4-flash"
  }
}
```

**字段说明**：

| 环境变量 | 作用 |
|----------|------|
| `ANTHROPIC_BASE_URL` | API 端点地址，指向本平台 |
| `ANTHROPIC_AUTH_TOKEN` | 认证令牌，即你的 `sk-` 开头 API Key |
| `ANTHROPIC_MODEL` | 默认使用的主模型 |
| `ANTHROPIC_DEFAULT_SONNET_MODEL` | Sonnet 别名对应的实际模型 |
| `ANTHROPIC_DEFAULT_OPUS_MODEL` | Opus 别名对应的实际模型 |
| `ANTHROPIC_DEFAULT_HAIKU_MODEL` | Haiku 别名对应的实际模型 |

【补充说明：Claude Code 内部以 Sonnet / Opus / Haiku 三个别名调度不同场景的模型。通过上述配置，可将所有别名统一映射到本平台模型，确保无论 Claude Code 内部调用哪个别名，实际请求均发往本平台。】

> API Key 获取方式：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

---

### 步骤二：新增 onboarding 标记

编辑或新增 `~/.claude.json` 文件（注意：此文件位于用户主目录下，非 `.claude/` 文件夹内）：

- macOS / Linux：`~/.claude.json`
- Windows：`C:\Users\你的用户名\.claude.json`

写入：

```json
{
  "hasCompletedOnboarding": true
}
```

【补充说明：此标记用于跳过 Claude Code 首次启动时的 Anthropic 官方账户登录引导。若不设置，Claude Code 会强制要求登录 Anthropic 账户或设置官方 API Key，无法直接使用第三方端点。】

---

### 步骤三：启动 Claude Code

配置完成后，进入工作目录，在终端中运行：

```bash
claude
```

启动后，选择 **信任此文件夹 (Trust This Folder)**，以允许 Claude Code 访问当前目录中的文件。

---

## 验证配置生效

启动 `claude` 后，在 TUI 中依次输入以下命令确认配置正确：

```text
/status
/model
```

**预期结果**：

- `/status` 应显示 `ANTHROPIC_BASE_URL` 指向 `https://www.aytdai.com/v1`
- `/model` 应显示当前模型为 `deepseek-v4-pro`

若显示异常，参照下方「异常排查」章节。

---

## 推荐模型配置

| 模型 ID | 适用场景 | 配置位置 |
|---------|----------|----------|
| `deepseek-v4-pro` | 编程主力模型（推荐） | `ANTHROPIC_MODEL` / `SONNET` / `OPUS` |
| `deepseek-v4-flash` | 轻量任务、降低成本 | `HAIKU` |
| `kimi-k2.7-code` | 代码审查、多模态（图片） | 按需替换 |

切换模型：在 Claude Code 中运行 `/model` 命令。完整模型列表见 [首页](/)。

> **⚠️ 兼容性限制**：`kimi/kimi-k3` 和 `MiniMax/MiniMax-M3` 不支持 Anthropic Messages 协议格式，不可配置为 Claude Code 的底层模型。

---

## 扩展思考（Extended Thinking）

本平台 DeepSeek 系列模型支持 Claude Code 的扩展思考功能，默认开启。

- 开关方式：运行 `/config`，将 **Thinking mode** 设为 `true` 或 `false`
- 快捷键：`Option+T`（macOS）或 `Alt+T`（Windows/Linux）

【补充说明：DeepSeek 模型的思考过程（thinking tokens）会计入输出 Token 消耗。若回答被截断，可适当调大 max_tokens 参数（建议 4096 以上），或在 `/config` 中关闭思考模式以节省 Token。】

---

## 异常排查

### 启动后仍跳转 Anthropic 登录页

**原因**：`~/.claude.json` 中缺少 onboarding 标记，或文件路径不正确。

**处理**：
1. 确认文件路径为 `~/.claude.json`（用户主目录下），而非 `~/.claude/settings.json`
2. 确认文件内容为 `{"hasCompletedOnboarding": true}`
3. 确认文件为合法 JSON 格式（无多余逗号、编码为 UTF-8）

### 连接失败 / 认证错误（401）

**排查清单**：
1. API Key 是否以 `sk-` 开头，无多余空格或换行
2. `ANTHROPIC_BASE_URL` 是否为 `https://www.aytdai.com/v1`（含 `/v1` 路径）
3. 是否存在残留的 Anthropic 官方环境变量覆盖了配置文件（执行 `echo $ANTHROPIC_AUTH_TOKEN` 和 `echo $ANTHROPIC_BASE_URL` 检查）
4. 账户余额是否充足

### 模型响应为空 / 回答截断

**原因**：DeepSeek 模型的思考过程消耗大量输出 Token，导致实际回答内容被截断。

**处理**：
- 在 `/config` 中调大 Max Output Tokens（建议 8192+）
- 或关闭 Thinking mode 以减少思考 Token 消耗

### 环境变量优先级问题

【补充说明：Claude Code 的配置优先级为：系统环境变量 > `settings.json` 中的 `env` 字段。若系统环境中已导出 `ANTHROPIC_AUTH_TOKEN` 或 `ANTHROPIC_BASE_URL`（例如在 `~/.bashrc`、`~/.zshrc`、Windows 系统变量中），则 `settings.json` 中的同名配置将被覆盖，请求会发往错误地址。务必确保系统环境中无残留。】
