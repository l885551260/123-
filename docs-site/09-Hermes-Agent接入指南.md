# Hermes Agent 接入指南

[Hermes Agent](https://github.com/NousResearch/hermes-agent) 是 Nous Research 做的开源 AI Agent，有持久记忆和自学习能力，支持 CLI、Telegram、Discord 等多平台。

---

## 安装

```bash
curl -fsSL https://raw.githubusercontent.com/NousResearch/hermes-agent/main/scripts/install.sh | bash
```

验证：`hermes doctor`

---

## 配置

```bash
hermes model
```

1. Provider 选 **OpenAI Compatible**
2. 填写：

| 填什么 | 填多少 |
|--------|--------|
| Base URL | `https://www.aytdai.com/v1` |
| API Key | 你的 `sk-xxx` 密钥 |

3. 模型选 `deepseek-v4-pro`

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

---

## 使用

```bash
hermes
```

直接对话就行。Hermes 会记住你的偏好和项目上下文，用得越多越好用。

---

## 推荐模型

写代码用 `deepseek-v4-pro`，省钱用 `deepseek-v4-flash`。完整列表看 [首页](/)。

---

## 常见问题

**连不上？** 检查密钥、Base URL、余额。

**回答截断？** 调大 max_tokens（4096+）。

**开关思考？** `/reasoning none` 关闭，`/reasoning high` 开启。
