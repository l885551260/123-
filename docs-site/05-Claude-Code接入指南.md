# Claude Code 接入指南

[Claude Code](https://github.com/anthropics/claude-code) 是 Anthropic 的终端 AI 编程工具，直接在命令行里读代码、写代码、跑命令。

---

## 第一步：安装

**macOS / Linux**（推荐）：
```bash
curl -fsSL https://claude.ai/install.sh | bash
```

**Windows**（PowerShell）：
```powershell
irm https://claude.ai/install.ps1 | iex
```

**或者用 npm**（需要 Node.js 22+）：
```bash
npm install -g @anthropic-ai/claude-code@latest
```

装完验证：
```bash
claude --version
```

---

## 第二步：配置接入我们平台

找到（或新建）配置文件：
- macOS / Linux：`~/.claude/settings.json`
- Windows：`C:\Users\你的用户名\.claude\settings.json`

写入以下内容：

```json
{
  "env": {
    "ANTHROPIC_BASE_URL": "https://www.aytdai.com/v1",
    "ANTHROPIC_API_KEY": "sk-你的密钥",
    "ANTHROPIC_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_OPUS_MODEL": "deepseek-v4-pro",
    "ANTHROPIC_DEFAULT_HAIKU_MODEL": "deepseek-v4-flash"
  }
}
```

> 把 `sk-你的密钥` 换成你自己的 API Key（[控制台](https://www.aytdai.com) → 创建API密钥）。

然后新建（或编辑）`~/.claude.json`（注意没有文件夹，是用户目录下的文件）：

```json
{
  "hasCompletedOnboarding": true
}
```

这一步是跳过 Anthropic 官方登录引导。

---

## 第三步：清除旧配置（重要！）

如果你之前用过 Anthropic 官方密钥，必须清掉残留的环境变量，否则会覆盖 settings.json 的配置：

**macOS / Linux**：
```bash
unset ANTHROPIC_API_KEY
unset ANTHROPIC_BASE_URL
```

**Windows PowerShell**：
```powershell
$env:ANTHROPIC_API_KEY = ""
$env:ANTHROPIC_BASE_URL = ""
```

如果这些变量写在 `~/.bashrc`、`~/.zshrc` 或系统环境变量里，也要去删掉。

---

## 第四步：启动验证

```bash
claude
```

进去后如果提示"信任此文件夹"，选信任。

**验证是否接入成功**：输入 `/status`，看到 Base URL 指向 `https://www.aytdai.com/v1` 就对了。

也可以用诊断命令检查环境：
```bash
claude doctor
```

---

## 推荐模型

| 填什么 | 适合 |
|--------|------|
| `deepseek-v4-pro` | 写代码首选（推荐） |
| `deepseek-v4-flash` | 简单任务，省钱 |
| `kimi-k2.7-code` | 代码审查、看图 |

切换模型：运行 `/model` 命令。完整模型列表看 [首页](/)。

> 注意：`kimi/kimi-k3` 和 `MiniMax/MiniMax-M3` 不支持 Claude 格式，别填这两个。

---

## 常见问题

**连不上 / 认证失败？**
- 检查密钥是不是 `sk-` 开头，没有多余空格
- 检查 Base URL 是不是 `https://www.aytdai.com/v1`（有 /v1）
- 检查有没有残留的 Anthropic 官方环境变量（`echo $ANTHROPIC_API_KEY` 看看）
- 检查余额够不够

**启动时还是跳到 Anthropic 登录？**
- 确认 `~/.claude.json` 里有 `"hasCompletedOnboarding": true`
- 确认 settings.json 路径正确（是 `~/.claude/settings.json`，不是 `~/.claude.json`）

**回答被截断？**
- DeepSeek 模型会先"思考"再回答，思考也占 Token
- 把 max_tokens 调大（4096+）

**开关思考模式？**
- `/config` 里改 Thinking mode，或快捷键 `Alt+T`
