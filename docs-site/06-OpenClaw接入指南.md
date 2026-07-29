# OpenClaw 接入指南

> 在 OpenClaw 中接入本平台模型，实现本地 AI 助手及多平台远程操控。

[**OpenClaw**](https://openclaw.ai) 是一个本地运行的开源 AI 助手框架，支持接入 Discord、Telegram、WhatsApp、Slack、微信等聊天平台进行远程操控。其模型层通过 Provider 插件机制对接 LLM 后端，支持任何 OpenAI 兼容 API 作为自定义 Provider 接入。

---

## 前置条件

| 条件 | 要求 |
|------|------|
| Node.js | 22.22.3+、24.15+ 或 25.9+（推荐 Node 26） |
| API Key | 本平台签发的 `sk-` 前缀密钥 |
| 操作系统 | macOS / Linux / Windows（原生或 WSL2） |

【补充说明：检查 Node.js 版本使用 `node --version`。若未安装或版本不符，参考 OpenClaw 官方 Node 安装指引。Windows 用户可选择原生 Windows Hub 应用、PowerShell 安装器或 WSL2 Gateway 路径。】

---

## 安装 OpenClaw

### macOS / Linux

```bash
curl -fsSL https://openclaw.ai/install.sh | bash
```

### Windows（PowerShell）

```powershell
iwr -useb https://openclaw.ai/install.ps1 | iex
```

【补充说明：其他安装方式（Docker、Nix、npm）参见 OpenClaw 官方 Install 文档。npm 方式命令为 `npm install -g openclaw`。】

---

## 运行初始化引导

```bash
openclaw onboard --install-daemon
```

该向导将依次完成：选择模型提供商、设置 API Key、配置 Gateway、安装守护进程。

【补充说明：引导过程中若提示选择 Provider 并填写 API Key，可先选择任意内置 Provider 跳过（后续通过手动配置替换为本平台）。完整引导参考见 OpenClaw 官方 Onboarding (CLI) 文档。跳过可选步骤后可随时通过 `openclaw configure` 返回补充。】

---

## 配置本平台为自定义 Provider

### 配置文件位置

OpenClaw 配置存储于：

- macOS / Linux：`~/.openclaw/openclaw.json`
- Windows：`C:\Users\你的用户名\.openclaw\openclaw.json`

【补充说明：旧版路径 `~/.clawdbot/clawdbot.json` 已通过符号链接自动兼容至新路径，无需手动迁移。】

### 配置结构说明

OpenClaw 的自定义 Provider 配置分为**两个必要部分**：

1. **Provider 定义**（`models.providers`）：声明 Provider 名称、API 端点、认证密钥、可用模型列表
2. **模型允许列表**（`agents.defaults.models`）：将模型的全限定标识符加入允许列表，并设置主模型

> **⚠️ 重要**：仅完成 Provider 定义而未将模型加入允许列表，将触发 `model not allowed` 错误。两部分缺一不可。

### 执行步骤

编辑 `~/.openclaw/openclaw.json`，写入以下内容（若文件已有内容，将对应字段合并进去）：

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

将 `sk-你的密钥` 替换为本平台 API Key（[控制台](https://www.aytdai.com) → 创建 API 密钥）。

### 字段逐项说明

**Provider 定义部分（`models.providers.aytdai`）**：

| 字段 | 值 | 说明 |
|------|-----|------|
| `baseUrl` | `https://www.aytdai.com/v1` | 本平台 OpenAI 兼容端点 |
| `apiKey` | `sk-你的密钥` | 认证密钥 |
| `api` | `openai-completions` | API 协议类型，OpenAI 兼容端点必须填此值 |
| `models[].id` | 模型标识符 | 必须与本平台 API 实际接受的 model 参数完全一致 |
| `models[].name` | 显示名称 | 用于 UI 展示，可自定义 |
| `models[].reasoning` | `true` / `false` | 是否支持推理模式 |
| `models[].input` | `["text"]` | 支持的输入模态 |
| `models[].contextWindow` | Token 数 | 模型上下文窗口大小 |
| `models[].maxTokens` | Token 数 | 单次最大输出 Token 数 |

【补充说明：`models.providers.*.contextWindow` / `maxTokens` 为 Provider 级默认值；`models.providers.*.models[].contextWindow` / `maxTokens` 为逐模型覆盖值。若 Provider 下多个模型共享相同上下文窗口，可在 Provider 级设置一次，无需逐模型重复。】

**模型允许列表部分（`agents.defaults`）**：

| 字段 | 值 | 说明 |
|------|-----|------|
| `model.primary` | `aytdai/deepseek-v4-pro` | 主模型，全限定格式为 `Provider名/模型id` |
| `models` 对象键 | `aytdai/deepseek-v4-pro` | 允许使用的模型全限定标识符 |
| `models` 对象值.alias | `ds-pro` | 快捷别名，用于 `/model` 命令快速切换 |

【补充说明：模型全限定标识符（Model ref）格式为 `provider/model`，例如 `aytdai/deepseek-v4-pro`。此格式是 OpenClaw 内部路由使用的唯一标识。`alias` 为用户友好的短名称，在聊天中通过 `/model ds-pro` 即可切换。】

---

## 应用配置

保存文件后，执行：

```bash
openclaw gateway config.apply --file ~/.openclaw/openclaw.json
```

Gateway 将自动重启以加载新配置。

---

## 验证配置生效

### 方法一：Control UI

```bash
openclaw dashboard
```

浏览器将打开控制面板（Control UI）。在聊天框发送任意消息，收到 AI 回复即为成功。

### 方法二：Gateway 状态检查

```bash
openclaw gateway status
```

**预期结果**：输出显示 Gateway 正在监听端口 18789。

### 方法三：Control UI 连接测试

在 Control UI 中进入 **Settings → Model Providers**，找到 `aytdai` Provider 卡片，点击 **Test connection**。

**预期结果**：显示连接成功及延迟数据。

【补充说明：Test connection 会发起一次真实 Provider 请求，可能消耗少量 Token。若返回错误，将显示分类信息（authentication / rate-limit / billing / timeout / response error），可据此定位问题。】

### 方法四：CLI 模型列表

```bash
openclaw models list
```

**预期结果**：输出中包含 `aytdai/deepseek-v4-pro` 和 `aytdai/deepseek-v4-flash`。

---

## 模型切换

### 聊天内切换

在对话中输入：

```
/model ds-pro
```

或

```
/model ds-flash
```

### CLI 切换默认模型

```bash
openclaw models set aytdai/deepseek-v4-pro
```

【补充说明：`openclaw configure` 或 `openclaw models auth login` 添加/重新认证 Provider 时，不会覆盖已有的 `agents.defaults.model.primary` 设置。仅在传入 `--set-default` 参数时才会变更主模型。因此日常配置更新不会意外切换已设定的主模型。】

---

## 添加更多模型

若需接入本平台其他模型（如 `kimi-k2.7-code`、`qwen3.7-plus`），执行以下步骤：

1. 在 `models.providers.aytdai.models` 数组中追加新模型对象：

```json
{
  "id": "kimi-k2.7-code",
  "name": "Kimi K2.7 Code",
  "reasoning": true,
  "input": ["text", "image"],
  "contextWindow": 131072,
  "maxTokens": 16384
}
```

2. 在 `agents.defaults.models` 中追加允许条目：

```json
"aytdai/kimi-k2.7-code": {
  "alias": "kimi-code"
}
```

3. 重新应用配置：

```bash
openclaw gateway config.apply --file ~/.openclaw/openclaw.json
```

完整可用模型列表见 [首页](/)。

---

## 异常排查

### 错误：`model not allowed: aytdai/模型名`

**原因**：模型已在 `models.providers` 中定义，但未加入 `agents.defaults.models` 允许列表。

**处理**：
1. 确认 `agents.defaults.models` 对象中存在键 `"aytdai/模型名"`（全限定格式）
2. 确认键名中的模型 ID 与 `models.providers.aytdai.models[].id` 完全一致（区分大小写）
3. 重新执行 `openclaw gateway config.apply`

### 错误：连接失败 / 无响应

**排查清单**：

1. API Key 是否以 `sk-` 开头，无多余空格或换行
2. `baseUrl` 是否为 `https://www.aytdai.com/v1`（含 `/v1` 路径）
3. `api` 字段是否为 `openai-completions`
4. 账户余额是否充足
5. 使用 curl 直接测试 API 连通性：

```bash
curl https://www.aytdai.com/v1/chat/completions \
  -H "Authorization: Bearer sk-你的密钥" \
  -H "Content-Type: application/json" \
  -d '{"model":"deepseek-v4-pro","messages":[{"role":"user","content":"hi"}]}'
```

若 curl 返回正常 JSON 响应，则问题在 OpenClaw 配置层面；若 curl 本身失败，则问题在网络、密钥或平台侧。

### 错误：模型未出现在 `/models` 列表中

**原因**：`models.providers.aytdai.models[]` 数组中缺少该模型的定义。

**处理**：确认同时完成了 Provider 模型定义和允许列表两个步骤。

### 回答截断

**原因**：DeepSeek 系列模型的思考过程（thinking tokens）计入输出 Token，导致实际回答被截断。

**处理**：在模型定义中调大 `maxTokens` 值（如 32000 或更高），重新应用配置。

### 思考模式开关

在对话中输入：
- `/think off`：关闭思考模式
- `/think adaptive`：恢复自适应思考

---

## 配置回滚

若需移除本平台 Provider，恢复使用其他模型服务：

1. 从 `models.providers` 中删除 `aytdai` 对象
2. 从 `agents.defaults.models` 中删除所有 `aytdai/` 前缀的条目
3. 将 `agents.defaults.model.primary` 改为其他 Provider 的模型（或删除该字段使用系统默认）
4. 执行 `openclaw gateway config.apply --file ~/.openclaw/openclaw.json`

【补充说明：也可通过 Control UI → Settings → Model Providers 页面进行可视化的 Provider 添加、替换、移除操作，以及 Default models 卡片管理主模型和回退模型。】
