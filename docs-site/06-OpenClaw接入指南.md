# OpenClaw 接入指南

[OpenClaw](https://github.com/openclaw/openclaw) 是一个本地运行的 AI 助手，可以接入微信、Telegram 等平台远程操控。

---

## 安装

**macOS / Linux**：
```bash
curl -fsSL https://openclaw.bot/install.sh | bash
```

**Windows**：
```powershell
iwr -useb https://openclaw.ai/install.ps1 | iex
```

---

## 配置

装完后会自动进入配置引导。如果没有，手动运行 `openclaw configure`。

按提示选：
1. **Where will the Gateway run?** → Local (this machine)
2. **Select sections to configure** → Model
3. **Model/auth provider** → OpenAI Compatible

然后填：

| 填什么 | 填多少 |
|--------|--------|
| Base URL | `https://www.aytdai.com/v1` |
| API Key | 你的 `sk-xxx` 密钥 |
| Model | `deepseek-v4-pro` |

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

---

## 验证

```bash
openclaw tui
```

能正常对话就成功了。

---

## 推荐模型

写代码用 `deepseek-v4-pro`，省钱用 `deepseek-v4-flash`，日常用 `qwen3.7-plus`。

完整列表看 [首页](/)。

---

## 常见问题

**连不上？** 检查密钥、Base URL、余额。

**回答截断？** DeepSeek 有思考过程占 Token，调大 max_tokens（4096+）。

**开关思考？** 对话中输入 `/think off` 关闭，`/think adaptive` 恢复。
