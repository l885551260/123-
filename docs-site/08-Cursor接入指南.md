# Cursor 接入指南

[Cursor](https://cursor.com) 是基于 VS Code 的 AI 代码编辑器，有 Chat、Composer、Edit 等 AI 功能。

> **前提**：Cursor 需要 **Pro 订阅**（付费会员）才能用自定义模型。免费用户配不了。

---

## 配置步骤

1. 打开 Cursor，点左侧栏 **Models**（或 Settings → Models）

2. 找到 **API Keys** 部分，展开：
   - 勾选 **Override OpenAI Base URL**
   - 在输入框填：`https://www.aytdai.com/v1`

3. 在 **OpenAI API Key** 输入框填入你的 `sk-xxx` 密钥

4. 点旁边的 **Verify**（验证）按钮 → 弹窗中点 **Enable OpenAI API Key**

5. 回到 Models 页面上方，点 **View All Models → Add Custom Model**，输入模型名：
   - `deepseek-v4-pro`（推荐）

6. 再添加几个备选：
   - `deepseek-v4-flash`（省钱）
   - `kimi-k2.7-code`（代码专精）

7. 确保添加的模型开关是**打开**状态

8. 在聊天面板（`Ctrl+L`）顶部的模型下拉里选你添加的模型，开始用

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

---

## 注意事项

- **Tab 自动补全不能用自定义模型**。Tab 是 Cursor 自带功能，只走 Cursor 自家模型。你的自定义模型只在 Chat / Composer / Edit（`Ctrl+K`）里生效。
- **Override 是全局的**。开了之后 Cursor 自带的 GPT 模型也会走你的地址。如果只想用自定义模型，把自带模型关掉就行。
- 如果模型没返回内容或超时，试试 Settings → Network → 把网络模式改成 **HTTP/1.0**。

---

## 常见问题

**验证失败 / 红色报错？**
- 密钥 `sk-` 开头，没多余空格
- Base URL 是 `https://www.aytdai.com/v1`（有 /v1，末尾没斜杠）
- 余额充足
- 清掉系统里的 `OPENAI_API_KEY` 和 `OPENAI_BASE_URL` 环境变量（会覆盖 Cursor 设置）

**报错 "does not work with your current plan"？**
- 你不是 Cursor Pro 会员，升级就行

**Chat 里看不到添加的模型？**
- 确认在 Models 页面里模型开关是开的
- 重启 Cursor 试试

**回答截断？**
- DeepSeek 有思考过程占 Token，在 Settings → Models 里调大 Max Tokens（4096+）

完整模型列表看 [首页](/)。
