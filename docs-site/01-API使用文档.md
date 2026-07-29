# API 使用文档

> 这个页面是给**开发者**看的。如果你只是想用聊天客户端或编程工具，直接看首页的推荐就行。

---

## 你需要准备什么

1. 一个平台账号（[注册](https://www.aytdai.com)，送 ¥1 体验金）
2. 一个 API Key（登录后 → 控制台 → 创建API密钥，得到 `sk-xxx` 格式的字符串）

> API Key 相当于你的"通行证"，别泄露给别人。

---

## 核心概念（30 秒看完）

- **Base URL**：`https://www.aytdai.com/v1`（所有请求都发到这个地址）
- **认证方式**：请求头带上 `Authorization: Bearer sk-你的密钥`
- **计费**：按 Token 收费（Token 可以理解为"字数"，1个中文字 ≈ 2个 Token）
- **格式**：兼容 OpenAI 和 Claude 两种格式，用哪个都行

---

## 可用模型

| 模型 ID | 说明 |
|---------|------|
| `deepseek-v4-flash` | 快速便宜，日常任务 |
| `deepseek-v4-pro` | 强推理，写代码（推荐） |
| `kimi-k2.7-code` | 代码专精，支持图片 |
| `kimi/kimi-k3` | Kimi 旗舰（仅 OpenAI 格式） |
| `qwen3.7-flash` | 最便宜 |
| `qwen3.7-plus` | 均衡 |
| `qwen3.7-max` | 通义最强 |
| `glm-5.2` | 智谱通用 |
| `MiniMax-M2.5` | MiniMax 通用 |
| `MiniMax/MiniMax-M3` | MiniMax 旗舰（仅 OpenAI 格式） |

> 带"仅 OpenAI 格式"的模型不能用 Claude 格式调用，其他模型两种格式都可以。

查询最新模型列表：

```bash
curl https://www.aytdai.com/v1/models -H "Authorization: Bearer sk-你的密钥"
```

---

## 用 OpenAI 格式调用（最常用）

绝大多数客户端和 SDK 都用这个格式。

### curl 示例

```bash
curl https://www.aytdai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer sk-你的密钥" \
  -d '{
    "model": "deepseek-v4-pro",
    "messages": [{"role": "user", "content": "你好"}]
  }'
```

### Python

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-你的密钥",
    base_url="https://www.aytdai.com/v1"
)

response = client.chat.completions.create(
    model="deepseek-v4-pro",
    messages=[{"role": "user", "content": "你好"}]
)
print(response.choices[0].message.content)
```

### Node.js

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-你的密钥',
  baseURL: 'https://www.aytdai.com/v1'
});

const response = await client.chat.completions.create({
  model: 'deepseek-v4-pro',
  messages: [{ role: 'user', content: '你好' }]
});
console.log(response.choices[0].message.content);
```

### Go

```go
config := openai.DefaultConfig("sk-你的密钥")
config.BaseURL = "https://www.aytdai.com/v1"
client := openai.NewClientWithConfig(config)

resp, _ := client.CreateChatCompletion(context.Background(),
    openai.ChatCompletionRequest{
        Model:    "deepseek-v4-pro",
        Messages: []openai.ChatCompletionMessage{{Role: "user", Content: "你好"}},
    })
fmt.Println(resp.Choices[0].Message.Content)
```

### 流式输出

加一个参数就行：`"stream": true`

---

## 用 Claude 格式调用

适合用 Anthropic SDK 或 Claude Code 的场景。

**和 OpenAI 格式的区别**：
- 认证头不同：用 `x-api-key: sk-你的密钥`（不是 Authorization）
- 端点不同：`/v1/messages`（不是 /v1/chat/completions）
- 必须加一个头：`anthropic-version: 2023-06-01`

### curl 示例

```bash
curl https://www.aytdai.com/v1/messages \
  -H "Content-Type: application/json" \
  -H "x-api-key: sk-你的密钥" \
  -H "anthropic-version: 2023-06-01" \
  -d '{
    "model": "deepseek-v4-pro",
    "max_tokens": 4096,
    "messages": [{"role": "user", "content": "你好"}]
  }'
```

### Python

```python
import anthropic

client = anthropic.Anthropic(
    api_key="sk-你的密钥",
    base_url="https://www.aytdai.com/v1"
)

message = client.messages.create(
    model="deepseek-v4-pro",
    max_tokens=4096,
    messages=[{"role": "user", "content": "你好"}]
)

# 注意：DeepSeek 模型会先返回"思考过程"，再返回正式回答
# 要找 type == "text" 的块才是最终答案
for block in message.content:
    if block.type == "text":
        print(block.text)
```

### TypeScript

```typescript
import Anthropic from '@anthropic-ai/sdk';

const client = new Anthropic({
  apiKey: 'sk-你的密钥',
  baseURL: 'https://www.aytdai.com/v1',
});

const message = await client.messages.create({
  model: 'deepseek-v4-pro',
  max_tokens: 4096,
  messages: [{ role: 'user', content: '你好' }],
});

const textBlock = message.content.find(b => b.type === 'text');
console.log(textBlock?.text);
```

> **坑点提醒**：DeepSeek 模型用 Claude 格式时，会先输出一段"思考过程"（thinking），这段也计入 Token 费用。`max_tokens` 建议设 4096 以上，否则可能只拿到思考过程就被截断了。

---

## 客户端怎么填地址？

| 客户端 | 地址填 | 说明 |
|--------|--------|------|
| ChatGPT-Next-Web | `https://www.aytdai.com` | 它自己会加 /v1 |
| Lobe Chat | `https://www.aytdai.com/v1` | 要手动加 /v1 |
| Cherry Studio | `https://www.aytdai.com/v1` | 选"自定义提供商" |
| OpenAI SDK | `https://www.aytdai.com/v1` | 填 base_url |
| Anthropic SDK | `https://www.aytdai.com/v1` | 填 base_url |

---

## 其他端点

| 端点 | 干嘛的 | 能用吗 |
|------|--------|--------|
| `/v1/chat/completions` | 聊天（OpenAI 格式） | ✅ |
| `/v1/messages` | 聊天（Claude 格式） | ✅ |
| `/v1/completions` | 文本补全 | ✅ |
| `/v1/embeddings` | 文本向量化 | ⏳ 还没接 |
| `/v1/images/generations` | 画图 | ⏳ 还没接 |

---

## 计费规则

- 按 Token 数扣费，输入（你发的）和输出（AI 回的）分开算
- 不管用 OpenAI 格式还是 Claude 格式，价格一样
- 详细价格看 → [定价与充值说明](02-定价与充值说明.md)
