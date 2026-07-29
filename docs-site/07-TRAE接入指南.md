# TRAE 接入指南

> 在 TRAE 中接入本平台模型，进行 AI 编程。

[**TRAE**](https://www.trae.ai) 是字节跳动出品的 AI-native IDE，主打智能协作与 Agent 自动化。本平台提供 OpenAI 兼容 API，可通过自定义模型方式接入。

---

## 安装 TRAE

1. 访问 [TRAE 官网](https://www.trae.cn/) 下载并安装 TRAE
2. 首次启动时，根据指引完成初始设置
3. 登录 TRAE

---

## 配置本平台模型

TRAE 支持通过 **API Key** 接入自定义模型，配置步骤如下：

### 1. 打开设置

在 AI 对话框右上角，点击 **设置** 图标。

### 2. 选择模型页签

进入模型管理页面。

### 3. 添加模型

点击 **+ 添加模型** 按钮，选择服务商为 **OpenAI**（或 **自定义/OpenAI Compatible**）。

### 4. 填写配置

| 配置项 | 值 |
|--------|-----|
| API Key | 你的平台 API Key（`sk-` 开头） |
| Base URL | `https://www.aytdai.com/v1` |
| 模型名称 | `deepseek-v4-pro`（推荐）|

> **获取 API Key**：登录 [控制台](https://www.aytdai.com) → 创建 API 密钥。

### 5. 完成添加

点击 **添加模型** 按钮。

> **注意**：TRAE 将调用接口检测 API Key 是否有效：
> - 若连接成功，该自定义模型会被添加
> - 若连接失败，窗口中会展示错误信息，可参考排查

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

可重复上述步骤添加多个模型，在对话时自由切换。

---

## 常见问题

### 添加模型时连接失败

- 确认 API Key 填写正确（`sk-` 开头）
- 确认 Base URL 为 `https://www.aytdai.com/v1`（注意末尾 `/v1`）
- 确认账户有可用额度（登录控制台查看）
- 检查网络是否能正常访问 `https://www.aytdai.com`

### 响应中断 / 不完整

- DeepSeek 模型带有思考过程（thinking），会消耗输出 token
- 如遇截断，在 TRAE 设置中调大 max_tokens（建议 4096+）

---

## 费用说明

- 按 **Token** 用量计费，输入和输出分别计价
- 新用户注册即送 **¥1 体验额度**
- 详见 [定价与充值说明](https://www.aytdai.com/docs/#/02-定价与充值说明)
