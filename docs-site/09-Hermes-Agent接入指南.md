# Hermes Agent 接入指南

> 在 Hermes Agent 中接入本平台模型，实现自主 AI 编程。

[**Hermes Agent**](https://github.com/NousResearch/hermes-agent) 是 [Nous Research](https://nousresearch.com) 出品的开源自我进化 AI Agent 框架。具备跨会话持久记忆、内置自改进学习能力、40+ 集成工具以及多平台支持（CLI、Telegram、Discord、Slack、WhatsApp）。

---

## 前置条件

| 条件 | 要求 |
|------|------|
| API Key | 本平台签发的 `sk-` 前缀密钥（[控制台](https://www.aytdai.com) → 创建 API 密钥） |
| 操作系统 | macOS、Linux 或 Windows WSL2 |
| 终端 | 可正常执行 bash 命令 |

---

## 安装 Hermes Agent

在终端中运行一键安装命令：

```bash
curl -fsSL https://raw.githubusercontent.com/NousResearch/hermes-agent/main/scripts/install.sh | bash
```

验证安装：

```bash
hermes doctor
```

**预期结果**：输出环境检查信息，各项显示通过。

更多信息参考 [Hermes Agent 官方文档](https://hermes-agent.nousresearch.com/docs/)。

---

## 配置本平台模型

Hermes Agent 通过 `hermes model` 命令进行 Provider 和模型的交互式配置。

### 方法一：交互式配置（推荐）

运行模型选择器：

```bash
hermes model
```

**步骤 1**：从 Provider 列表中选择 **"Custom endpoint (self-hosted / VLLM / etc.)"**。

【补充说明：MiniMax 等部分服务商在 Hermes Agent 中拥有内置 Provider（如 "MiniMax China (mainland China endpoint)"），可直接选用。本平台暂无 Hermes Agent 内置 Provider，因此需通过 Custom endpoint 方式接入。该选项支持任何 OpenAI 兼容 API 端点。】

**步骤 2**：按提示依次填写：

| 提示项 | 填写内容 |
|--------|----------|
| API base URL | `https://www.aytdai.com/v1` |
| API key | `sk-你的密钥` |
| Model name | `deepseek-v4-pro` |

**步骤 3**：确认保存。

### 方法二：手动编辑配置文件

编辑 `~/.hermes/config.yaml`（Windows WSL2 路径同 Linux）：

```yaml
model:
  default: deepseek-v4-pro
  provider: custom
  base_url: https://www.aytdai.com/v1
  api_key: sk-你的密钥
```

【补充说明：`config.yaml` 中 `model.default` 与 `model.model` 两个键名均可用于指定模型 ID，功能等价。即 `model: { default: my-model }` 与 `model: { model: my-model }` 效果相同。】

### 方法三：环境变量

编辑 `~/.hermes/.env` 文件：

```bash
OPENAI_API_KEY=sk-你的密钥
OPENAI_BASE_URL=https://www.aytdai.com/v1
```

【补充说明：此方式通过 openai-api Provider 路径生效。设置后需通过 `hermes model` 选择 openai-api Provider 并指定模型。与方法一（Custom endpoint）为两条独立路径，选择其一即可。】

---

## 开始使用

运行：

```bash
hermes
```

开始与由本平台模型驱动的 Hermes Agent 对话。

【补充说明：Hermes Agent 的持久记忆和自改进技能系统意味着使用越多效果越好——编程模式、项目上下文和偏好会自动跨会话记忆。】

---

## 模型切换

### 会话内切换

在对话中输入：

```
/model
```

可在已配置好的 Provider 和模型之间快速切换。

### 添加新 Provider

若需添加新的 Provider（如切换到其他模型服务），需先退出当前会话（`Ctrl+C` 或 `/quit`），然后在终端运行：

```bash
hermes model
```

完成新 Provider 的配置后，重新启动会话。

【补充说明：`/model`（会话内命令）仅能在已配置好的 Provider 之间切换。若要添加全新的 Provider，必须使用 `hermes model`（终端命令）。两者用途不同，不可混淆。】

---

## 推荐模型

| 模型 ID | 适用场景 |
|---------|----------|
| `deepseek-v4-pro` | 编程主力模型（推荐） |
| `deepseek-v4-flash` | 轻量任务、降低成本 |
| `kimi-k2.7-code` | 代码审查 |
| `qwen3.7-plus` | 日常对话 |

完整可用模型列表见 [首页](/)。

---

## 开关思考

运行时用 `/reasoning` 命令控制思考模式：

| 参数 | 效果 |
|------|------|
| `/reasoning none` | 关闭思考 |
| `/reasoning minimal` | 最小思考 |
| `/reasoning low` | 低档思考 |
| `/reasoning medium` | 中档思考 |
| `/reasoning high` | 高档思考 |
| `/reasoning xhigh` | 最高档思考 |

【补充说明：`none` 关闭思考，其余档位（minimal / low / medium / high / xhigh）均为开启状态，区别在于思考深度和 Token 消耗。档位越高，推理越深入但 Token 消耗越大。】

---

## 异常排查

### 连接失败

**排查清单**：

1. API Key 是否以 `sk-` 开头，无多余空格
2. `base_url` 是否为 `https://www.aytdai.com/v1`（含 `/v1`）
3. 账户余额是否充足
4. 运行 `hermes doctor` 检查环境是否存在问题

### hermes model 中找不到 Custom endpoint 选项

**处理**：
1. 更新到最新版本：`hermes update`
2. 若仍无该选项，使用方法二（手动编辑 config.yaml）或方法三（环境变量）完成配置

### 回答截断

**原因**：DeepSeek 系列模型的思考过程（thinking tokens）计入输出 Token。

**处理**：调大 max_tokens（建议 4096+），或使用 `/reasoning none` 关闭思考以减少 Token 消耗。

---

## 配置回滚

若需恢复使用其他模型服务：

1. 运行 `hermes model`，选择其他 Provider 并完成配置
2. 或手动编辑 `~/.hermes/config.yaml`，将 `base_url` 和 `api_key` 改为目标服务的值
3. 若使用环境变量方式，编辑 `~/.hermes/.env` 修改对应值

完整可用模型列表见 [首页](/)。
