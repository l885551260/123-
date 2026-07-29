# 欢迎使用安逸通达

安逸通达 www.aytdai.com 是一个 AI 智能接口平台，把多种主流 AI 模型统一接入到一个接口里。**一个账号、一个 API Key 就能调用所有模型**。

## 快速开始

1. **注册账号** — 访问 [www.aytdai.com](https://www.aytdai.com)，支持用户名/邮箱/手机号注册，新用户赠送 **¥1** 体验额度
2. **获取 API Key** — 登录后进入 **控制台** → **创建API密钥** 
3. **开始使用** — 支持 OpenAI、Claude 两种 API 格式，选你熟悉的即可
4. **充值** — 体验额度用完后，支持支付宝/微信扫码充值，¥10 起充

## 当前支持的模型

| 厂商 | 模型 | 适合场景 |
|------|------|--------|
| DeepSeek | `deepseek-v4-flash` | 日常对话、翻译、轻量任务 |
| DeepSeek | `deepseek-v4-pro` | 复杂推理、代码生成 |
| 月之暗面 | `kimi-k2.7-code` | 编程、代码审查（支持图片输入） |
| 月之暗面 | `kimi/kimi-k3` | 旗舰综合模型 |
| 通义千问 | `qwen3.7-flash` / `qwen3.7-plus` / `qwen3.7-max` | 轻量→通用→复杂推理 |
| 智谱 | `glm-5.2` | 通用大模型 |
| MiniMax | `MiniMax-M2.5` / `MiniMax/MiniMax-M3` | 通用大模型 |

> 模型列表可能随平台接入情况调整，以 `GET /v1/models` 返回为准。详见 [API 使用文档](01-API使用文档.md)。

## 支持的客户端

支持 ChatGPT-Next-Web、Lobe Chat、Cherry Studio、Cursor、Windsurf 等任何兼容 OpenAI API 的客户端。

配置详情见 [API 使用文档 — 第三方客户端配置速查](01-API使用文档.md#第三方客户端配置速查)。

## 文档导航

- **[API 使用文档](01-API使用文档.md)** — OpenAI / Claude 两种格式的接口说明和代码示例
- **[定价与充值](02-定价与充值说明.md)** — 模型价格、充值方式、计费规则
- **[常见问题](03-常见问题FAQ.md)** — 注册、使用、充值相关 FAQ
- **[错误码参考](04-错误码参考.md)** — 错误排查指南

### AI 编程工具接入指南

- **[Claude Code 接入指南](05-Claude-Code接入指南.md)** — Anthropic 官方终端编程 Agent
- **[OpenClaw 接入指南](06-OpenClaw接入指南.md)** — 本地运行的个人 AI 助手
- **[TRAE 接入指南](07-TRAE接入指南.md)** — 字节跳动 AI IDE
- **[Cursor 接入指南](08-Cursor接入指南.md)** — AI 驱动的智能代码编辑器
- **[Hermes Agent 接入指南](09-Hermes-Agent接入指南.md)** — Nous Research 自进化 AI Agent
- **[Codex 接入指南](10-Codex接入指南.md)** — OpenAI 本地终端编程 Agent
- **[其他工具接入指南](11-其他工具接入指南.md)** — Dify、Cherry Studio、Zed、LangChain 等