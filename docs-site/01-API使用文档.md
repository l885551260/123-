# API 使用文档

## 快速开始

### 1. 注册账号

访问 [https://www.aytdai.com](https://www.aytdai.com) 注册账号，支持以下注册方式：

- **用户名 + 密码**
- **邮箱 + 验证码**（支持 Gmail、163、QQ、Outlook 等主流邮箱）
- **手机号 + 短信验证码**

新用户注册即送 **¥1 体验额度**，无需充值即可开始使用。

### 2. 获取 API Key

登录后进入 **控制台** → **创建API密钥**。

你可以：

- 创建多个令牌，分别用于不同项目
- 为令牌设置**模型限制**（只允许调用指定模型）
- 为令牌设置 **IP 白名单**（只允许指定 IP 调用）
- 随时禁用或删除令牌

> **安全提示**：API Key 等同于账户权限，请勿泄露或提交到公开代码仓库。

### 3. 选择 API 格式

本平台同时兼容 **OpenAI** 和 **Claude (Anthropic)** 两种 API 格式，使用同一个平台 API Key，调用同一组模型（个别模型仅支持 OpenAI 格式，见下方“可用模型”节注明）：

| | OpenAI 格式 | Claude 格式 |
|---|---|---|
| **Base URL** | `https://www.aytdai.com/v1` | `https://www.aytdai.com/v1` |
| **聊天端点** | `/v1/chat/completions` | `/v1/messages` |
| **认证头** | `Authorization: Bearer sk-xxx` | `x-api-key: sk-xxx` |
| **流式输出** | ✅ `stream: true` | ✅ `stream: true` |
| **适用场景** | OpenAI SDK、大多数客户端 | Anthropic SDK、Claude Code 等 |

选择你熟悉的格式即可，下文分别给出两种格式的完整示例。

---

## 可用模型

平台当前提供以下模型，**均支持流式输出**，绝大多数模型同时兼容 OpenAI 和 Claude 两种格式：

### DeepSeek

| 模型 ID | 定位说明 |
|---------|----------|
| `deepseek-v4-flash` | 快速响应，性价比高，适合轻量任务、简单问答 |
| `deepseek-v4-pro` | 更强推理能力（推荐），适合复杂编程、代码生成 |

### 月之暗面（Kimi）

| 模型 ID | 定位说明 |
|---------|----------|
| `kimi-k2.7-code` | 代码专精，支持图片输入，适合编程、代码审查 |
| `kimi/kimi-k3` | Kimi 旗舰模型，综合能力强 |

### 阿里通义千问（Qwen）

| 模型 ID | 定位说明 |
|---------|----------|
| `qwen3.7-flash` | 3.7 快速版，极致性价比，适合轻量任务 |
| `qwen3.7-plus` | 3.7 增强版，性能均衡，适合通用任务 |
| `qwen3.7-max` | 3.7 旗舰版，能力最强，适合复杂推理 |

### 智谱（GLM）

| 模型 ID | 定位说明 |
|---------|----------|
| `glm-5.2` | 智谱 GLM 通用大模型 |

### MiniMax

| 模型 ID | 定位说明 |
|---------|----------|
| `MiniMax-M2.5` | MiniMax 通用大模型 |
| `MiniMax/MiniMax-M3` | MiniMax M3 大模型 |

> **注意**：模型名中含 `/` 的模型（`kimi/kimi-k3`、`MiniMax/MiniMax-M3`）目前仅支持 OpenAI 格式，Claude 格式暂不可用。其余模型均同时支持两种格式。
>
> 模型列表可能随平台接入情况调整，以 `GET /v1/models` 返回为准。

也可以通过 API 查询当前可用模型列表：

**端点**: `https://www.aytdai.com/v1/models`

```bash
curl https://www.aytdai.com/v1/models \
  -H "Authorization: Bearer sk-xxxxxxxxxxxxxxxx"
```

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-xxxxxxxxxxxxxxxx",
    base_url="https://www.aytdai.com/v1"
)

models = client.models.list()
for model in models.data:
    print(model.id)
```

---

## OpenAI 格式（最常用）

适用于 OpenAI SDK、ChatGPT-Next-Web、Lobe Chat、Cursor 等大多数客户端。

### 地址与认证

```
Base URL: https://www.aytdai.com/v1
认证头:   Authorization: Bearer sk-xxxxxxxxxxxxxxxx
```

> ChatGPT-Next-Web 会自动拼接 `/v1`，只需填 `https://www.aytdai.com`；Lobe Chat 需要手动加 `/v1`。

### 聊天补全

**端点**: `https://www.aytdai.com/v1/chat/completions`

```bash
curl https://www.aytdai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer sk-xxxxxxxxxxxxxxxx" \
  -d '{
    "model": "deepseek-v4-flash",
    "messages": [
      {"role": "system", "content": "你是一个有帮助的助手。"},
      {"role": "user", "content": "你好"}
    ]
  }'
```

**流式输出**：添加 `"stream": true` 参数即可。

#### Python

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-xxxxxxxxxxxxxxxx",
    base_url="https://www.aytdai.com/v1"
)

response = client.chat.completions.create(
    model="deepseek-v4-flash",
    messages=[
        {"role": "system", "content": "你是一个有帮助的助手。"},
        {"role": "user", "content": "你好"}
    ]
)
print(response.choices[0].message.content)
```

#### Node.js

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-xxxxxxxxxxxxxxxx',
  baseURL: 'https://www.aytdai.com/v1'
});

const response = await client.chat.completions.create({
  model: 'deepseek-v4-flash',
  messages: [
    { role: 'system', content: '你是一个有帮助的助手。' },
    { role: 'user', content: '你好' }
  ]
});
console.log(response.choices[0].message.content);
```

#### Go

```go
package main

import (
    "context"
    "fmt"
    openai "github.com/sashabaranov/go-openai"
)

func main() {
    config := openai.DefaultConfig("sk-xxxxxxxxxxxxxxxx")
    config.BaseURL = "https://www.aytdai.com/v1"
    client := openai.NewClientWithConfig(config)

    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: "deepseek-v4-flash",
            Messages: []openai.ChatCompletionMessage{
                {Role: "user", Content: "你好"},
            },
        },
    )
    if err != nil { panic(err) }
    fmt.Println(resp.Choices[0].Message.Content)
}
```

---

## Claude (Anthropic) 格式

适用于 Anthropic Python/TypeScript SDK、Claude Code，以及使用 Claude Messages API 格式的客户端。

### 地址与认证

```
Base URL: https://www.aytdai.com/v1
认证头:   x-api-key: sk-xxxxxxxxxxxxxxxx
          anthropic-version: 2023-06-01
```

> 使用平台 API Key 作为 `x-api-key` 的值，无需 Claude 官方密钥。

### 消息创建

**端点**: `https://www.aytdai.com/v1/messages`

```bash
curl https://www.aytdai.com/v1/messages \
  -H "Content-Type: application/json" \
  -H "x-api-key: sk-xxxxxxxxxxxxxxxx" \
  -H "anthropic-version: 2023-06-01" \
  -d '{
    "model": "deepseek-v4-flash",
    "max_tokens": 1024,
    "messages": [
      {"role": "user", "content": "你好"}
    ]
  }'
```

#### Python

```python
import anthropic

client = anthropic.Anthropic(
    api_key="sk-xxxxxxxxxxxxxxxx",
    base_url="https://www.aytdai.com/v1"
)

message = client.messages.create(
    model="deepseek-v4-flash",
    max_tokens=1024,
    messages=[{"role": "user", "content": "你好"}]
)

# 响应可能包含 thinking（思考过程）和 text（正式回复）
# 需要找到 type 为 text 的内容块
for block in message.content:
    if block.type == "text":
        print(block.text)
        break
```

#### TypeScript

```typescript
import Anthropic from '@anthropic-ai/sdk';

const client = new Anthropic({
  apiKey: 'sk-xxxxxxxxxxxxxxxx',
  baseURL: 'https://www.aytdai.com/v1',
});

const message = await client.messages.create({
  model: 'deepseek-v4-flash',
  max_tokens: 1024,
  messages: [{ role: 'user', content: '你好' }],
});

// 响应可能包含 thinking（思考过程）和 text（正式回复）
const textBlock = message.content.find(b => b.type === 'text');
console.log(textBlock?.text);
```

> **注意**：DeepSeek 模型的 Claude 格式响应会先返回一个 `thinking` 思考块，再返回 `text` 正式回复。直接取 `content[0]` 可能拿到的是思考过程而非最终回答，建议遍历 `content` 找到 `type: "text"` 的块。思考过程也计入输出 token，建议 `max_tokens` 设置为 1024 以上。

---

## 其他可用端点

| 端点 | 说明 | 状态 |
|------|------|------|
| `/v1/chat/completions` | 聊天补全 | ✅ 已上线 |
| `/v1/completions` | 文本补全 | ✅ 已上线 |
| `/v1/messages` | Claude 格式消息 | ✅ 已上线 |
| `/v1/embeddings` | 文本嵌入 | ⏳ 待接入模型 |
| `/v1/images/generations` | 图像生成 | ⏳ 待接入模型 |
| `/v1/audio/speech` | 语音合成 (TTS) | ⏳ 待接入模型 |
| `/v1/audio/transcriptions` | 语音识别 (STT) | ⏳ 待接入模型 |

---

## 第三方客户端配置速查

### 聊天客户端

| 客户端 | 地址填写 | 备注 |
|--------|---------|------|
| ChatGPT-Next-Web | `https://www.aytdai.com` | 自动拼接 `/v1` |
| Lobe Chat | `https://www.aytdai.com/v1` | 手动加 `/v1` |
| Cherry Studio | `https://www.aytdai.com/v1` | 自定义提供商 |
| DeepChat | 一键导入 | 控制台提供导入链接 |
| AionUI | 一键导入 | 控制台提供导入链接 |
| OpenCat | `https://www.aytdai.com` | 团队设置 |

### AI 编程 / Agent 工具

以下工具均支持 OpenAI 兼容 API，配置 Base URL 为 `https://www.aytdai.com/v1`，填入平台 API Key 即可使用。

| 工具 | 简介 | GitHub | 详细指南 |
|------|------|--------|---------|
| Claude Code | Anthropic 官方终端编程 Agent | [anthropics/claude-code](https://github.com/anthropics/claude-code) | [接入指南](05-Claude-Code接入指南.md) |
| OpenClaw | 本地运行的个人 AI 助手，支持多平台集成 | [openclaw/openclaw](https://github.com/openclaw/openclaw) | [接入指南](06-OpenClaw接入指南.md) |
| TRAE | 字节跳动 AI IDE，内置智能编程助手 | [bytedance/trae-agent](https://github.com/bytedance/trae-agent) | [接入指南](07-TRAE接入指南.md) |
| Cursor | AI 驱动的智能代码编辑器 | [getcursor/cursor](https://github.com/getcursor/cursor) | [接入指南](08-Cursor接入指南.md) |
| Hermes Agent | Nous Research 自进化 AI Agent | [nousresearch/hermes-agent](https://github.com/nousresearch/hermes-agent) | [接入指南](09-Hermes-Agent接入指南.md) |
| Codex | OpenAI 本地终端编程 Agent | [openai/codex](https://github.com/openai/codex) | [接入指南](10-Codex接入指南.md) |

> **通用配置方法**：在工具的模型设置中选择 **OpenAI Compatible**，Base URL 填 `https://www.aytdai.com/v1`，API Key 填平台密钥，模型填 `deepseek-v4-pro` 等可用模型即可。

其他工具（Dify、Cherry Studio、Zed、LangChain 等）请参考 [其他工具接入指南](11-其他工具接入指南.md)。

---

## 速率限制

平台对请求频率有合理限制，正常使用不会触发。如遇限流请降低请求频率，或添加指数退避重试逻辑。

---

## 计费

- 按 **Token** 计费，输入和输出分别计价
- 1 Token ≈ 0.75 个英文单词 / 0.5 个中文字符
- 无论使用哪种 API 格式，计费规则相同
- 详见 [定价页面](02-定价与充值说明.md)
