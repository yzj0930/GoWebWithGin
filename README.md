# GoWebWithGin
## 说明
基于Gin框架进行二次开发的web功能

## 前置条件
go环境可用

## 执行步骤
``` bash
go mod init "github.com/yzj0930/GoWebWithGin"
go mod tidy
go run main.go
```

## 框架结构
```markdown
GoWebWithGin/
├── 📁 cmd/                    # 应用程序入口
│   └── 📁 server/
│       └── main.go           # 主程序入口
├── 📁 internal/              # 内部应用程序代码
│   ├── 📁 controllers/       # HTTP 处理器 (Controllers)
│   │   ├── base_controller.go  # 路由基类，业务类继承该类方法，用于routes层获取路由信息以便注册路由
│   │   ├── registry.go.go      # 维护静态controller注册表
│   │   ├── user_controller.go
│   │   ├── product_controller.go
│   │   └── auth_controller.go
│   ├── 📁 routes/            # 路由定义
│   │   └── routes_manager.go
│   ├── 📁 services/          # 业务逻辑层
│   │   ├── user_service.go
│   │   ├── product_service.go
│   │   └── auth_service.go
│   ├── 📁 repositories/      # 数据访问层
│   │   ├── user_repo.go
│   │   ├── product_repo.go
│   │   └── base_repo.go
│   ├── 📁 models/            # 数据模型
│   │   ├── user.go
│   │   ├── product.go
│   │   └── base_model.go
│   ├── 📁 middleware/        # 中间件
│   │   ├── auth.go
│   │   ├── logger.go
│   │   └── cors.go
│   └── 📁 dto/               # 数据传输对象
│       ├── requests/         # 请求DTO
│       └── responses/        # 响应DTO
├── 📁 pkg/                   # 可公开使用的包
│   ├── 📁 config/            # 配置管理
│   ├── 📁 database/          # 数据库连接
│   ├── 📁 cache/             # 缓存管理
│   ├── 📁 utils/             # 工具函数
│   └── 📁 logger/            # 日志管理
├── 📁 api/                   # API 定义
│   └── 📁 docs/              # API 文档
├── 📁 web/                   # 前端资源
│   ├── 📁 static/            # 静态文件
│   └── 📁 templates/         # 模板文件
├── 📁 scripts/               # 脚本文件
├── 📁 deployments/           # 部署配置
├── 📁 tests/                 # 测试文件
├── go.mod
├── go.sum
├── Makefile
└── README.md