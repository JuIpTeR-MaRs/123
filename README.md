# shop_server

基于 Go 语言开发的电商后端服务系统，采用模块化设计，支持用户管理、商品管理、购物车、订单处理、评论等功能，适用于中小型电商平台的后端支撑。

## 项目概述

`shop_server` 是一个基于 Go 语言开发的电商后端服务系统，采用模块化设计，支持用户管理、商品管理、购物车、订单处理、评论等功能。适用于中小型电商平台的后端支撑。

### 核心特性

- 基于 Gin 框架实现高性能 HTTP 路由
- 使用 GORM 操作 MySQL 数据库
- Redis 支持会话或缓存存储
- 日志通过 Zap + Lumberjack 实现滚动切割
- 配置文件使用 YAML 格式，由 Viper 加载
- 时间处理使用 `carbon` 库增强可读性

## 功能模块

- **用户管理**：注册、登录、信息更新、修改密码
- **商品管理**：商品增删改查、分类管理
- **购物车管理**：添加、删除、修改数量、清空购物车
- **订单管理**：创建订单、查询订单、取消订单、支付订单、发货、完成订单、申请退款
- **评论管理**：发表评论、查看评论、更新评论、删除评论
- **后台管理**：管理员登录、用户列表查询、用户详情查询、封禁/解封用户

## 技术架构

### 技术栈

- **后端**: Go 1.23.0, Gin v1.11.0
- **数据库**: MySQL（gorm.io/driver/mysql v1.6.0）
- **缓存**: Redis（github.com/redis/go-redis/v9 v9.16.0）
- **ORM**: GORM v1.31.1
- **配置管理**: Viper v1.21.0
- **日志库**: zap + lumberjack.v2
- **工具库**: carbon（时间处理）、errors（错误堆栈）

### 项目结构

```
.
├── apis                    # API 控制器层
│   ├── admin_ctl           # 管理员控制器
│   ├── cart_ctl            # 购物车控制器
│   ├── order_ctl           # 订单控制器
│   ├── products_ctl        # 商品控制器
│   ├── review_ctl          # 评价控制器
│   └── users_ctl           # 用户控制器
├── cmd
│   └── main.go             # 项目入口文件
├── config                  # 配置文件
│   ├── config.go
│   └── config.yaml
├── internal
│   ├── defs                # 常量定义
│   ├── logics              # 业务逻辑层
│   └── models              # 数据模型层
├── pkg                     # 通用包
│   ├── logs
│   ├── mysqldb
│   └── redisdb
├── requests                # 请求参数结构体
├── routers                 # 路由配置
├── sqls                    # SQL 脚本
├── static                  # 静态资源文件
└── utils                   # 工具函数
```

### 架构模式

- **MVC 分层模式**：apis → logics → models 实现关注点分离
- **单例模式**：数据库连接（MySQL、Redis）在 pkg 层初始化为全局实例
- **配置中心模式**：使用 viper 统一管理配置项
- **中间件模式**：Gin 的中间件机制用于日志、认证等扩展

## API 接口说明

### 用户模块

- `POST /user/register` - 用户注册
- `POST /user/login` - 用户登录
- `GET /user/list` - 获取用户列表
- `POST /user/update` - 更新用户信息
- `POST /user/modify_psw` - 修改用户密码

### 管理员模块

- `POST /admin/login` - 管理员登录
- `POST /admin/userlist` - 管理员获取用户列表
- `POST /admin/queryuser` - 查询用户详情
- `POST /admin/blockuser` - 封禁用户
- `POST /admin/unblockuser` - 解封用户
- `GET /admin/searchuser` - 搜索用户

### 商品模块

- `GET /product/getcategory` - 获取商品分类列表
- `POST /product/addcategory` - 添加商品分类
- `POST /product/updatecategory` - 更新商品分类
- `POST /product/deletecategory` - 删除商品分类
- `GET /product/getproduct` - 获取商品列表
- `POST /product/addproduct` - 添加商品
- `POST /product/updateproduct` - 更新商品
- `POST /product/deleteproduct` - 删除商品

### 购物车模块

- `POST /carts` - 添加商品到购物车
- `GET /carts/:userId` - 获取用户购物车商品
- `PUT /carts/:cartId` - 更新购物车商品数量
- `DELETE /carts/:cartId` - 从购物车移除商品
- `DELETE /carts/clear/:userId` - 清空用户购物车

### 订单模块

- `POST /orders` - 创建订单
- `GET /orders/:orderNo` - 获取订单详情
- `GET /orders/user/:userId` - 获取用户订单列表
- `PUT /orders/cancel/:orderNo` - 取消订单
- `PUT /orders/pay/:orderNo` - 支付订单
- `PUT /orders/deliver/:orderNo` - 发货
- `PUT /orders/complete/:orderNo` - 完成订单
- `PUT /orders/refund/:orderNo` - 申请退款

### 评论模块

- `POST /reviews` - 添加评价
- `GET /reviews/product/:productId` - 获取商品评价
- `GET /reviews/:reviewId` - 获取单个评价详情
- `PUT /reviews/:reviewId` - 更新评价
- `DELETE /reviews/:reviewId` - 删除评价
- `GET /reviews/user/:userId` - 获取用户的所有评价

## 快速开始

### 环境准备

- Go 1.23.0
- MySQL 5.7+
- Redis 6.0+
- Git

### 安装部署

1. 克隆项目至 `$GOPATH/src/shop_server`

```bash
git clone <repository-url>
```

2. 安装依赖

```bash
go mod download
```

3. 创建数据库并导入 SQL 脚本

```bash
mysql -u root -p
create database shop_db;
exit

mysql -u root -p shop_db < sqls/shop.sql
```

4. 修改配置文件 `config/config.yaml`，配置数据库和 Redis 地址

```yaml
mysql:
  host: "127.0.0.1"
  port: "3306"
  database: "shop_db"
  user: "root"
  Password: "84916325"  # 请修改为您的实际密码
  maxIdleConns: 10
  maxOpenConns: 50

redis:
  host: "127.0.0.1"
  port: "6379"
  database: "1"
  password: ""
```

5. 运行服务

```bash
go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动

## 项目配置

项目使用 Viper 库管理配置，支持多种格式配置文件，当前使用 YAML 格式。

### 配置项说明

- **system.port**: 服务运行端口，默认为 8080
- **system.mode**: 运行模式（debug/release/test）
- **logger.stdout**: 是否在控制台输出日志
- **logger.level**: 日志级别（debug/info/warn/error）
- **logger.dir**: 日志文件存储目录
- **logger.logMaxAge**: 日志文件最大保存时间（天）
- **logger.logTypes**: 日志输出类型（stdout/file）
- **mysql.host**: MySQL 主机地址
- **mysql.port**: MySQL 端口
- **mysql.database**: 数据库名
- **mysql.user**: 用户名
- **mysql.Password**: 密码
- **mysql.maxIdleConns**: 最大空闲连接数
- **mysql.maxOpenConns**: 最大打开连接数
- **redis.host**: Redis 主机地址
- **redis.port**: Redis 端口
- **redis.database**: Redis 数据库编号
- **redis.password**: Redis 密码

## 构建与部署

### 构建命令

```bash
go build -o shop_server cmd/main.go
```

### 部署方式

- 直接运行二进制文件
- 配合 systemd 管理服务
- Docker 容器化部署

### 日志管理

- 日志路径: 默认输出至 logs 目录
- 按大小切分（lumberjack 配置）
- 日志文件最大保存时间由 `logMaxAge` 配置项控制

## 安全说明

- SQL 注入防护由 GORM 参数化查询保障
- 敏感接口需添加 JWT 或 Session 鉴权（当前版本未实现）
- 请在生产环境中使用强密码并配置访问控制

## 已知问题

- 项目包含大量间接依赖（indirect），可能存在冗余或安全隐患，建议定期审计
- [prouct_ctl.go](file:///C:%5CUsers%5C30408%5CGolandProjects%5C%E6%96%B0%E5%BB%BA%E6%96%87%E4%BB%B6%E5%A4%B9%5Capis%5Cproducts_ctl%5Cprouct_ctl.go) 文件名拼写错误（应为 product_ctl）
- 未提供单元测试和接口文档（如 Swagger）
- 安全认证机制未明确展示，存在安全风险

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助我们改进这个项目。

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件