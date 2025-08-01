# BlogProject

## 项目简介
本项目是一个基于 Go、Gin 和 GORM 的个人博客系统，实现了用户注册/登录、文章管理、评论、JWT 认证等功能，支持基础的博客后台管理。

## 技术栈
- Go 1.x
- Gin
- GORM
- MySQL
- JWT

## 目录结构
```
├── cmd/           # 程序入口
├── config/        # 配置文件和数据库初始化
├── handlers/      # 路由处理器
├── middleware/    # Gin中间件
├── router/        # 路由注册
├── services/      # 业务逻辑
├── structs/       # 结构体定义
├── utils/         # 工具函数
├── go.mod         # Go 依赖管理
├── README.md      # 项目说明
```

## 快速开始
1. 克隆项目
   ```bash
   git clone <your-repo-url>
   cd blogPorject
   ```
2. 安装依赖
   ```bash
   go mod tidy
   ```
3. 配置数据库
   - 修改 `config/config.yaml`，填写你的 MySQL 连接信息。
4. 启动服务
   ```bash
   go run cmd/main.go
   ```

## 主要 API 示例

### 用户相关
- **注册**
  - `POST /user/register`
  - 参数：`username`、`password`、`email`
- **登录**
  - `POST /user/login`
  - 参数：`username`、`password`
  - 返回：`token`

### 文章相关
- **发表文章**
  - `POST /post/createPost`（需认证）
  - 参数：`title`、`content`
- **获取文章列表**
  - `GET /post/all?page=1&pageSize=10`
- **获取单篇文章详情**
  - `GET /post/GetPostForId?postId=1`
- **更新/删除文章**
  - `POST /post/update`（需认证）
  - 参数：`postid`、`content`、`userid`、`opType`（U=更新，D=删除）

### 评论相关
- **发表评论**
  - `POST /comment/create`（需认证）
  - 参数：`postid`、`content`、`userid`
- **获取某篇文章的评论**
  - `GET /comment/list?postid=1`
- **获取用户所有评论**
  - `GET /comment/user?userid=1`
- **删除评论**
  - `POST /comment/delete`（需认证）
  - 参数：`commentid`、`userid`

## 统一响应结构
所有接口返回如下结构：
```json
{
  "code": 200,
  "message": "操作成功",
  "data": { ... }
}
```

## 错误处理与日志记录
- 所有接口返回统一结构，包含 code、message、data 字段。
- 关键操作和错误会自动记录到日志，便于调试和维护。
- GORM SQL 日志和 Gin 请求日志已全局开启。

## 贡献
欢迎提交 issue 和 PR 参与项目改进！
