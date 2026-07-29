# Cursor 接入指南

[Cursor](https://cursor.com) 是基于 VS Code 的 AI 代码编辑器。

> **前提**：Cursor 需要 Pro 订阅（付费会员）才能用自定义模型。免费用户配不了。

---

## 配置步骤

1. 打开 Cursor → 左侧栏点 **Models**

2. 展开 **API Keys** 部分：
   - 勾选 **Override OpenAI Base URL**
   - 填入：`https://www.aytdai.com/v1`

3. 在 **OpenAI API Key** 输入框填入你的 `sk-xxx` 密钥

4. 点右侧验证按钮 → 弹窗中点 **Enable OpenAI API Key**

5. 在 Models 板块点 **View All Models → Add Custom Model**，输入模型名：
   - `deepseek-v4-pro`（推荐）
   - `deepseek-v4-flash`（省钱）
   - `kimi-k2.7-code`（代码专精）

6. 启用添加的模型，在聊天面板选择即可使用

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

---

## 注意事项

- **Tab 自动补全不能用自定义模型**。Tab 是 Cursor 自带功能，只用 Cursor 自家模型。你的自定义模型只在 Chat / Composer / Edit 里生效。
- **Override 是全局的**。开了之后 Cursor 自带的 Claude/GPT 也会走你的地址。不用的时候记得关。
- 如果模型没返回内容，试试把 Network 设置改成 HTTP/1.0。

---

## 常见问题

**验证失败？**
- 密钥 `sk-` 开头，没多余空格
- Base URL 是 `https://www.aytdai.com/v1`（有 /v1）
- 余额充足
- 清掉系统里的 `OPENAI_API_KEY` 和 `OPENAI_BASE_URL` 环境变量

**报错 "does not work with your current plan"？**
- 你不是 Cursor Pro 会员，升级就行

完整模型列表看 [首页](/)。
