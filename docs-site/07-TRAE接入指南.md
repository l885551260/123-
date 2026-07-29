# TRAE 接入指南

[TRAE](https://www.trae.ai) 是字节跳动的 AI IDE（代码编辑器），内置 Builder（Agent）、Chat、Inline Chat 三种 AI 模式。

---

## 第一步：安装

- 国际版：去 [trae.ai](https://www.trae.ai) 下载（macOS / Windows / Linux）
- 国内版：去 [trae.cn](https://www.trae.cn) 下载（macOS / Windows）

两个版本都支持自定义模型，配置方法一样。

---

## 第二步：添加自定义模型

1. 打开 TRAE，点右上角 **⚙️ 设置** 图标
2. 左侧选 **Models**（模型）
3. 点 **Add model**（添加模型）
4. Provider（服务商）选 **OpenAI**
5. Model 下拉选 **Custom Model**（自定义模型，在列表最下面）
6. 填写：

| 字段 | 填什么 |
|------|--------|
| Model ID | `deepseek-v4-pro` |
| API Key | 你的 `sk-xxx` 密钥 |
| Custom Request URL | `https://www.aytdai.com/v1/chat/completions` |

7. 点确认完成添加

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

**⚠️ 关键：URL 必须写到 `/v1/chat/completions`**

TRAE v3.3.51 之后，Custom Request URL 是**原样使用**的，不会自动补路径。所以：
- ✅ 正确：`https://www.aytdai.com/v1/chat/completions`
- ❌ 错误：`https://www.aytdai.com/v1`（会 404）
- ❌ 错误：`https://www.aytdai.com`（会 404）

---

## 第三步：使用

回到编辑器，点顶部（或对话框里）的**模型下拉菜单**，选你刚添加的 `deepseek-v4-pro`。

然后就可以：
- **Builder 模式**：AI 直接帮你改代码、跑命令（类似 Agent）
- **Chat 模式**：侧边栏问答
- **Inline Chat**：`Ctrl+I` 在代码里直接对话

---

## 推荐模型

可以添加多个模型，随时切换：

| Model ID | 适合 |
|----------|------|
| `deepseek-v4-pro` | 写代码首选（推荐） |
| `deepseek-v4-flash` | 简单任务，省钱 |
| `kimi-k2.7-code` | 代码专精 |
| `qwen3.7-plus` | 日常对话 |

每个模型都要单独添加一条（重复第二步，改 Model ID 就行）。

完整列表看 [首页](/)。

---

## 常见问题

**添加后连接失败 / 404？**
- 99% 是 URL 没写全。必须是 `https://www.aytdai.com/v1/chat/completions`，不能只写到 `/v1`
- 密钥是不是 `sk-` 开头，没有多余空格
- 余额够不够

**Builder 模式卡住 / 工具调用失败？**
- 换 `deepseek-v4-pro` 试试（小模型在 Agent 模式下容易卡）
- 检查网络能不能打开 `https://www.aytdai.com`

**回答截断？**
- DeepSeek 有思考过程占 Token，在设置里调大 max_tokens（4096+）

**找不到 Custom Model 选项？**
- 确认 TRAE 是最新版本（Help → Check for Updates）
- 国际版和国内版都有这个选项，但位置可能略有不同
