# 常见问题 FAQ

## 注册与账号

### Q: 支持哪些注册方式？
A: 三种方式任选：
- 用户名 + 密码
- 邮箱 + 验证码（支持 Gmail、163、QQ、Outlook、Hotmail、iCloud、Yahoo、Foxmail）
- 手机号 + 短信验证码

### Q: 注册后需要充值才能使用吗？
A: 不需要。新用户注册即送 **¥1 体验额度**，可以先体验再决定是否充值。

### Q: 忘记密码怎么办？
A: 在登录页点击「忘记密码」，通过绑定的邮箱或手机号重置。

### Q: 可以修改用户名吗？
A: 用户名注册后不可修改，但可以在个人设置中修改显示名称。

---

## API 使用

### Q: Base URL 应该填什么？
A: 取决于客户端：

| 客户端 | 填写 |
|--------|------|
| ChatGPT-Next-Web | `https://www.aytdai.com`（自动加 /v1） |
| Lobe Chat | `https://www.aytdai.com/v1`（手动加 /v1） |
| OpenAI SDK (Python/Node) | `https://www.aytdai.com/v1` |
| Claude SDK | `https://www.aytdai.com/v1` |

### Q: 支持哪些模型？
A: 目前支持 10 个模型：
- `deepseek-v4-flash` / `deepseek-v4-pro` — DeepSeek 系列
- `kimi-k2.7-code` / `kimi/kimi-k3` — 月之暗面 Kimi 系列
- `qwen3.7-flash` / `qwen3.7-plus` / `qwen3.7-max` — 通义千问 3.7 系列
- `glm-5.2` — 智谱 GLM
- `MiniMax-M2.5` / `MiniMax/MiniMax-M3` — MiniMax 系列

完整列表可通过 `GET /v1/models` 查询，详见 [API 使用文档](01-API使用文档.md)。

### Q: 支持哪些 API 格式？
A: 同时兼容两种格式：
- **OpenAI 格式** — `/v1/chat/completions`
- **Claude 格式** — `/v1/messages`

使用同一个平台 API Key，详见 [API 使用文档](01-API使用文档.md)。

### Q: 为什么调用报错 "model not found"？
A: 请确认模型名称拼写正确（注意大小写和横杠）。

### Q: 可以同时调用多个模型吗？
A: 可以。同一个 API Key 可以调用平台上所有可用模型。

### Q: 接口和 OpenAI 完全一样吗？
A: 聊天补全接口完全兼容 OpenAI 格式，可以直接用 OpenAI 的 SDK，只需修改 Base URL。同时也兼容 Claude 的原生格式。

---

## 充值与费用

### Q: 最低充值多少？
A: 最低 ¥10。

### Q: 充值后多久到账？
A: 支付宝和微信都是即时到账。

### Q: 额度会过期吗？
A: 不会。充值的额度永久有效。

### Q: 怎么查看消费明细？
A: 登录后进入 **控制台** → **日志**，可以查看每次调用的模型、Token 用量和费用。

### Q: 支持退款吗？
A: 如有特殊情况请联系客服邮箱：aytdai@163.com

---

## 令牌 (API Key)

### Q: 令牌和账号密码有什么区别？
A: 令牌（格式 `sk-xxxx`）是给 API 调用用的密钥。账号密码是用来登录控制台的。

### Q: 可以创建多个令牌吗？
A: 可以。你可以为不同项目创建不同的令牌，还可以设置：
- **模型限制** — 令牌只允许调用指定模型
- **IP 白名单** — 令牌只允许从指定 IP 调用
- **额度限制** — 为令牌设置单独的使用额度上限

### Q: 令牌泄露了怎么办？
A: 立即到控制台**禁用**或**删除**该令牌，然后创建新令牌。

---

## 安全

### Q: 平台有什么安全措施？
A:
- **阿里云验证码** — 注册和登录时的人机验证
- **短信验证** — 手机号注册使用阿里云短信服务
- **邮箱验证** — 注册需验证邮箱
- **IP 限制** — 令牌可设置 IP 白名单
- **数据加密** — 全站 HTTPS 加密传输

---

## 客户端

### Q: 推荐用什么客户端？
A: 推荐以下几款：
- **ChatGPT-Next-Web** — 开源 Web 客户端，部署简单
- **Lobe Chat** — 功能丰富，界面美观
- **Cherry Studio** — 桌面端，支持多模型管理
- **Cursor / Windsurf** — AI 编程编辑器

### Q: 控制台里的一键导入链接怎么用？
A: 在控制台令牌页面，点击客户端名称可以一键导入配置到 Cherry Studio、DeepChat、AionUI 等客户端。

---

如有其他问题，可以通过邮箱 aytdai@163.com 联系我们。