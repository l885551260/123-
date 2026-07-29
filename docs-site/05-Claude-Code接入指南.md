# Claude Code 接入指南

[Claude Code](https://github.com/anthropics/claude-code) 是 Anthropic 做的终端 AI 编程工具。下面是接入我们平台的步骤。

---

## 第一步：安装 Claude Code

看 [官方文档](https://code.claude.com/docs/en/setup)，装好就行。

---

## 第二步：改配置文件

找到（或新建）这个文件：
- macOS / Linux：`~/.claude/settings.json`
- Windows：`用户目录/.claude/settings.json`

写入：

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

> 把 `sk-你的密钥` 换成你自己的 API Key（控制台 → 创建API密钥）。

然后再新建（或编辑）`~/.claude.json`：

```json
{
  "hasCompletedOnboarding": true
}
```

---

## 第三步：清除旧配置（重要！）

如果你之前配过 Anthropic 官方密钥，必须清掉：

```bash
unset ANTHROPIC_AUTH_TOKEN
unset ANTHROPIC_BASE_URL
```

如果这两行写在 `~/.bashrc` 或 `~/.zshrc` 里，也要删掉。

---

## 第四步：启动

```bash
claude
```

进去后选"信任此文件夹"，就能用了。

**验证配置**：输入 `/status`，看到 Base URL 指向 `https://www.aytdai.com/v1` 就对了。

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
- 检查密钥是不是 `sk-` 开头
- 检查 Base URL 是不是 `https://www.aytdai.com/v1`
- 检查有没有残留的 Anthropic 官方环境变量
- 检查余额够不够

**回答被截断？**
- DeepSeek 模型会先"思考"再回答，思考也占 Token
- 把 max_tokens 调大（4096+）

**开关思考模式？**
- `/config` 里改 Thinking mode，或快捷键 `Alt+T`
