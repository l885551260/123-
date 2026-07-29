# 安逸通达 AI 平台

**一句话说明白**：我们是一个 AI 模型的"中转站"。你只需要一个账号、一个密钥，就能调用 DeepSeek、Kimi、通义千问、智谱、MiniMax 等多家 AI 模型，不用分别去各家注册。

---

## 三步开始用

| 步骤 | 做什么 | 怎么做 |
|------|--------|--------|
| ① 注册 | 创建账号 | 打开 [www.aytdai.com](https://www.aytdai.com)，手机号/邮箱/用户名都能注册。新用户送 **¥1** 体验金 |
| ② 拿密钥 | 获取 API Key | 登录后点 **控制台 → 创建API密钥**，会得到一个 `sk-` 开头的字符串 |
| ③ 用起来 | 接入工具或写代码 | 把密钥填到你想用的工具里（Cursor、Claude Code 等），或者用代码直接调用 |

> 体验金用完了？控制台 → 钱包 → 充值，支付宝/微信扫码，¥10 起充，即时到账。

---

## 有哪些模型可以用？

| 模型 | 谁家出的 | 干什么用 |
|------|---------|---------|
| `deepseek-v4-flash` | DeepSeek | 便宜快速，日常聊天、翻译 |
| `deepseek-v4-pro` | DeepSeek | 能力强，写代码、复杂推理（**推荐**） |
| `kimi-k2.7-code` | 月之暗面 | 专门写代码，还能看图 |
| `kimi/kimi-k3` | 月之暗面 | Kimi 家最强的通用模型 |
| `qwen3.7-flash` | 阿里通义千问 | 最便宜，简单任务 |
| `qwen3.7-plus` | 阿里通义千问 | 均衡，啥都能干 |
| `qwen3.7-max` | 阿里通义千问 | 通义家最强，复杂推理 |
| `glm-5.2` | 智谱 | 通用大模型 |
| `MiniMax-M2.5` | MiniMax | 通用大模型 |
| `MiniMax/MiniMax-M3` | MiniMax | MiniMax 家更强版本 |

> 模型可能随时增减，以实际查询为准。

---

## 我想用聊天客户端（不会写代码）

推荐这几款，下载后填上我们的地址和密钥就能用：

| 客户端 | 类型 | 地址怎么填 |
|--------|------|-----------|
| [Cherry Studio](https://github.com/CherryHQ/cherry-studio) | 桌面软件 | `https://www.aytdai.com/v1` |
| [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) | 网页 | `https://www.aytdai.com`（它会自动加 /v1） |
| [Lobe Chat](https://github.com/lobehub/lobe-chat) | 网页 | `https://www.aytdai.com/v1` |

详细配置方法看 → [API 使用文档](01-API使用文档.md)

---

## 我想用 AI 编程工具

| 工具 | 一句话介绍 | 接入教程 |
|------|-----------|---------|
| Claude Code | 终端里的 AI 程序员 | [教程](05-Claude-Code接入指南.md) |
| Cursor | 最火的 AI 代码编辑器 | [教程](08-Cursor接入指南.md) |
| TRAE | 字节跳动做的 AI IDE | [教程](07-TRAE接入指南.md) |
| Codex | OpenAI 的编程 Agent | [教程](10-Codex接入指南.md) |
| OpenClaw | 本地 AI 助手 | [教程](06-OpenClaw接入指南.md) |
| Hermes Agent | 会自我进化的 AI Agent | [教程](09-Hermes-Agent接入指南.md) |
| 其他（Dify、Zed、LangChain…） | 看这里 | [教程](11-其他工具接入指南.md) |

> **通用配置就三样东西**：选 OpenAI 兼容 → 地址填 `https://www.aytdai.com/v1` → 密钥填你的 `sk-xxx`。

---

## 多少钱？

按用量付费，用多少扣多少，不用不花钱。举几个例子：

- 问一个简单问题：约 ¥0.0004（不到一厘钱）
- 翻译一篇文章：约 ¥0.006
- **充 ¥10 大概能用 1000~5000 次对话**

完整价格表 → [定价与充值说明](02-定价与充值说明.md)

---

## 遇到问题了？

- 先看看 → [常见问题 FAQ](03-常见问题FAQ.md)
- 报错了 → [错误码参考](04-错误码参考.md)
- 还是搞不定 → 发邮件：aytdai@163.com
