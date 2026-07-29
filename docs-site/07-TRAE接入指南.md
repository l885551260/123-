# TRAE 接入指南

[TRAE](https://www.trae.ai) 是字节跳动做的 AI IDE（代码编辑器），内置 AI 编程助手。

---

## 安装

去 [TRAE 官网](https://www.trae.cn/) 下载安装，首次启动按指引设置。

---

## 配置

1. 打开 TRAE，在 AI 对话框右上角点 **设置** 图标
2. 进入 **模型** 页签
3. 点 **+ 添加模型**，服务商选 **OpenAI**（或 Custom/OpenAI Compatible）
4. 填写：

| 填什么 | 填多少 |
|--------|--------|
| API Key | 你的 `sk-xxx` 密钥 |
| Base URL | `https://www.aytdai.com/v1` |
| 模型名称 | `deepseek-v4-pro` |

5. 点 **添加模型** 完成

> 密钥在 [控制台](https://www.aytdai.com) → 创建API密钥 获取。

TRAE 会自动验证连接。成功就添加上了，失败会显示错误信息。

---

## 推荐模型

写代码用 `deepseek-v4-pro`，省钱用 `deepseek-v4-flash`。可以添加多个模型，对话时随时切换。

完整列表看 [首页](/)。

---

## 常见问题

**添加时连接失败？**
- 密钥是不是 `sk-` 开头
- Base URL 末尾有 `/v1`（别漏了）
- 余额够不够
- 网络能不能打开 `https://www.aytdai.com`

**回答截断？**
- DeepSeek 有思考过程占 Token，在设置里调大 max_tokens（4096+）
