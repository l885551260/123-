# Hermes Agent 接入指南

> 在 Hermes Agent 中接入本平台模型，进行自主 AI 编程。

[**Hermes Agent**](https://github.com/NousResearch/hermes-agent) 是 Nous Research 出品的开源自我进化 AI Agent 框架。本平台提供 OpenAI 兼容 API，可直接作为 Hermes Agent 的模型后端。

---

## 前提条件

- 拥有本平台 API Key（登录 [控制台](https://www.aytdai.com) → 创建 API 密钥）
- 一台可访问终端的电脑（macOS、Linux 或 Windows WSL2）

---

## 安装 Hermes Agent

Hermes Agent 具备跨会话持久记忆、内置自改进学习能力、40+ 集成工具以及多平台支持（CLI、Telegram、Discord、Slack、WhatsApp）。

在终端中运行一键安装命令：

```bash
curl -fsSL https://raw.githubusercontent.com/NousResearch/hermes-agent/main/scripts/install.sh | bash
```

验证安装：

```bash
hermes doctor
```

更多信息请参考 [Hermes Agent 文档](https://hermes-agent.nousresearch.com/docs/)。

---

## 配置本平台模型

运行模型选择器：

```bash
hermes model
```

1. 从 provider 列表中选择 **OpenAI Compatible**（或 **Custom**）

2. 填写配置：

   | 配置项 | 值 |
   |--------|-----|
   | Base URL | `https://www.aytdai.com/v1` |
   | API Key | 你的平台 API Key（`sk-` 开头） |

3. 选择模型：`deepseek-v4-pro`（推荐）

---

## 可用模型

| 模型 ID | 说明 | 适用场景 |
|---------|------|----------|
| `deepseek-v4-pro` | 更强推理能力（推荐） | 复杂编程、代码生成 |
| `deepseek-v4-flash` | 快速响应，性价比高 | 轻量任务、简单问答 |
| `kimi-k2.7-code` | 代码专精，支持图片输入 | 编程、代码审查 |
| `kimi/kimi-k3` | Kimi 旗舰 | 综合任务 |
| `qwen3.7-max` | 通义千问旗舰 | 复杂推理 |
| `qwen3.7-plus` | 通义千问增强 | 通用任务 |
| `qwen3.7-flash` | 通义千问快速 | 轻量任务 |
| `glm-5.2` | 智谱 GLM | 通用任务 |
| `MiniMax-M2.5` | MiniMax | 通用任务 |
| `MiniMax/MiniMax-M3` | MiniMax M3 | 通用任务 |

> 模型列表以 `GET /v1/models` 返回为准。

---

## 开始使用

运行 `hermes`，开始与由本平台模型驱动的 Hermes Agent 对话。

> **提示**：Hermes Agent 的持久记忆和自改进技能系统意味着使用越多效果越好——你的编程模式、项目上下文和偏好会自动跨会话记忆。

---

## 开关思考

运行时用 `/reasoning` 开关思考：`none` 关闭，其他档位（minimal/low/medium/high/xhigh）开启。

> **提示**：DeepSeek 模型的思考过程（thinking）也计入输出 token，如遇响应截断，建议调大 max_tokens。

---

## 常见问题

### 连接失败 / 认证错误

- 确认 API Key 填写正确（`sk-` 开头）
- 确认 Base URL 为 `https://www.aytdai.com/v1`
- 确认账户有可用额度（登录控制台查看）

### 响应中断 / 不完整

- DeepSeek 模型带有思考过程，会消耗输出 token
- 如遇截断，调大 max_tokens（建议 4096+）

---

## 费用说明

- 按 **Token** 用量计费，输入和输出分别计价
- 新用户注册即送 **¥1 体验额度**
- 详见 [定价与充值说明](https://www.aytdai.com/docs/#/02-定价与充值说明)
