# OpenClaw 接入指南

> 在 OpenClaw 中接入本平台模型，实现本地 AI 助手。

[**OpenClaw**](https://github.com/openclaw/openclaw) 是本地运行的个人 AI 助手，可与多种通讯平台集成实现远程操控。本平台提供 OpenAI 兼容 API，可直接作为 OpenClaw 的模型后端。

---

## 安装 OpenClaw

在终端中运行以下命令安装（老用户同样可以用此命令更新）：

**macOS / Linux**：

```bash
curl -fsSL https://openclaw.bot/install.sh | bash
```

**Windows**：

```powershell
iwr -useb https://openclaw.ai/install.ps1 | iex
```

---

## 配置本平台模型

安装完成后会自动进入配置引导。若没有自动开始，运行：

```bash
openclaw configure
```

### 配置步骤

1. **Where will the Gateway run?** → 选择 **Local (this machine)**
2. **Select sections to configure** → 选择 **Model**
3. **Model/auth provider** → 选择 **OpenAI Compatible**（或 **Custom**）
4. 填写以下信息：

| 配置项 | 值 |
|--------|-----|
| Base URL | `https://www.aytdai.com/v1` |
| API Key | 你的平台 API Key（`sk-` 开头） |
| Model | `deepseek-v4-pro`（推荐）|

> **获取 API Key**：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

### 完成功能配置

按需配置以下选项：

- **Channel**：选择在哪个 App 中对话
- **Skill**：按需安装技能
- **Hooks**（可选）：
  - `session-memory`：执行 `/new` 时自动保存会话上下文
  - `command-logger`：记录所有命令到日志文件
  - `boot-md`：网关启动时运行 BOOT.md

### 验证

输入 `openclaw tui`，能正常对话即配置成功。

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

## 开关思考

运行时用 `/think` 开关思考：`off` 关闭，`adaptive`（默认）让模型自行决定何时思考。

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
