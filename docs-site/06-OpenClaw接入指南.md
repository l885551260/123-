# OpenClaw 接入指南

[OpenClaw](https://openclaw.ai) 是一个本地运行的开源 AI 助手，可以接入微信、Telegram、Discord 等聊天平台远程操控。

---

## 第一步：安装

**前提**：需要 Node.js 22+（`node --version` 检查）。

**macOS / Linux**：
```bash
curl -fsSL https://openclaw.ai/install.sh | bash
```

**Windows**（PowerShell）：
```powershell
iwr -useb https://openclaw.ai/install.ps1 | iex
```

装完运行初始化引导：
```bash
openclaw onboard --install-daemon
```

引导会让你选模型提供商和填密钥。**可以先随便选一个跳过**，后面我们手动改成自己的。

---

## 第二步：配置接入我们平台

编辑配置文件 `~/.openclaw/openclaw.json`（Windows 为 `C:\Users\你的用户名\.openclaw\openclaw.json`）。

在文件里加入（或合并到已有的 JSON 中）：

```json
{
  "models": {
    "mode": "merge",
    "providers": {
      "aytdai": {
        "baseUrl": "https://www.aytdai.com/v1",
        "apiKey": "sk-你的密钥",
        "api": "openai-completions",
        "models": [
          {
            "id": "deepseek-v4-pro",
            "name": "DeepSeek V4 Pro",
            "reasoning": true,
            "input": ["text"],
            "contextWindow": 131072,
            "maxTokens": 16384
          },
          {
            "id": "deepseek-v4-flash",
            "name": "DeepSeek V4 Flash",
            "reasoning": true,
            "input": ["text"],
            "contextWindow": 131072,
            "maxTokens": 16384
          }
        ]
      }
    }
  },
  "agents": {
    "defaults": {
      "model": {
        "primary": "aytdai/deepseek-v4-pro"
      },
      "models": {
        "aytdai/deepseek-v4-pro": {
          "alias": "ds-pro"
        },
        "aytdai/deepseek-v4-flash": {
          "alias": "ds-flash"
        }
      }
    }
  }
}
```

> 把 `sk-你的密钥` 换成你的 API Key（[控制台](https://www.aytdai.com) → 创建API密钥）。

**两个关键点**：
1. `models.providers` 里定义提供商和模型列表
2. `agents.defaults.models` 里把模型加入允许列表（**漏了这步会报 "model not allowed"**）

保存后应用配置：
```bash
openclaw gateway config.apply --file ~/.openclaw/openclaw.json
```

Gateway 会自动重启。

---

## 第三步：验证

```bash
openclaw dashboard
```

浏览器会打开控制面板。在聊天框发条消息，能回复就成功了。

或者用命令行检查：
```bash
openclaw gateway status
```

看到 Gateway 在 18789 端口监听就正常。

---

## 切换模型

在聊天中输入：
```
/model ds-pro
```
或
```
/model ds-flash
```

想加更多模型？在 `models.providers.aytdai.models` 数组里追加，同时在 `agents.defaults.models` 里加对应的允许条目。

完整模型列表看 [首页](/)。

---

## 常见问题

**报 "model not allowed"？**
- 你只做了第一步（定义 provider），没做第二步（加允许列表）
- 确认 `agents.defaults.models` 里有 `"aytdai/模型名"` 这个键

**连不上 / 没回复？**
- 检查密钥是不是 `sk-` 开头
- 检查 `baseUrl` 是不是 `https://www.aytdai.com/v1`
- 检查余额够不够
- 用 curl 测试 API 是否通：
```bash
curl https://www.aytdai.com/v1/chat/completions \
  -H "Authorization: Bearer sk-你的密钥" \
  -H "Content-Type: application/json" \
  -d '{"model":"deepseek-v4-pro","messages":[{"role":"user","content":"hi"}]}'
```

**回答截断？** DeepSeek 有思考过程占 Token，在模型定义里把 `maxTokens` 调大（比如 32000）。

**开关思考？** 对话中输入 `/think off` 关闭，`/think adaptive` 恢复。
