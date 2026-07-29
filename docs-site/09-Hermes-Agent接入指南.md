# Hermes Agent 接入指南

[Hermes Agent](https://github.com/NousResearch/hermes-agent) 是 Nous Research 的开源 AI Agent，有持久记忆和自学习能力，支持终端、Telegram、Discord 等多平台。

---

## 第一步：安装

```bash
curl -fsSL https://raw.githubusercontent.com/NousResearch/hermes-agent/main/scripts/install.sh | bash
```

装完验证：
```bash
hermes doctor
```

看到环境检查通过就行。

---

## 第二步：配置接入我们平台

**方法一：交互式配置（推荐）**

```bash
hermes model
```

1. 在 Provider 列表里选 **Custom endpoint (self-hosted / VLLM / etc.)**
2. 按提示填写：

| 提示 | 填什么 |
|------|--------|
| API base URL | `https://www.aytdai.com/v1` |
| API key | 你的 `sk-xxx` 密钥 |
| Model name | `deepseek-v4-pro` |

3. 确认保存

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

**方法二：手动编辑配置文件**

编辑 `~/.hermes/config.yaml`（Windows 为 `C:\Users\你的用户名\.hermes\config.yaml`）：

```yaml
model:
  default: deepseek-v4-pro
  provider: custom
  base_url: https://www.aytdai.com/v1
  api_key: sk-你的密钥
```

**方法三：环境变量**

编辑 `~/.hermes/.env` 文件：

```bash
OPENAI_API_KEY=sk-你的密钥
OPENAI_BASE_URL=https://www.aytdai.com/v1
```

这种方式走 openai-api provider，模型用 `hermes model` 切换。

---

## 第三步：使用

```bash
hermes
```

直接对话就行。Hermes 会记住你的偏好和项目上下文，用得越多越好用。

**切换模型**（在对话中）：
```
/model
```

**切换模型**（在终端中，添加新 provider）：
```bash
hermes model
```

> 注意：`/model` 只能在已配好的 provider 之间切换。要添加新 provider 得退出对话，跑 `hermes model`。

---

## 推荐模型

| 模型 | 适合 |
|------|------|
| `deepseek-v4-pro` | 写代码首选（推荐） |
| `deepseek-v4-flash` | 简单任务，省钱 |
| `kimi-k2.7-code` | 代码审查 |
| `qwen3.7-plus` | 日常对话 |

完整列表看 [首页](/)。

---

## 常见问题

**连不上？**
- 检查密钥是不是 `sk-` 开头
- 检查 base_url 是不是 `https://www.aytdai.com/v1`
- 检查余额够不够
- 跑 `hermes doctor` 看环境有没有问题

**回答截断？** 调大 max_tokens（4096+）。

**开关思考？** `/reasoning none` 关闭，`/reasoning high` 开启。

**hermes model 里没有 Custom endpoint 选项？**
- 更新到最新版：`hermes update`
- 或者直接用方法二/方法三手动配置
