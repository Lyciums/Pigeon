# Pigeon
An HTTP service framework based on Gin encapsulation

#### 项目结构
```
├── LICENSE
├── README.md
├── main.go
├── go.mod
├── go.sum
├── app - 应用相关的文件
│   ├── Controllers - 路由控制器
│   │   ├── Index.go - 处理根请求示例
│   │   └── index.nginx.conf
│   ├── Interfaces - 存放应用程序接口
│   │   └── Router.go - 路由控制器接口
│   ├── Middlewares - 存放中间件
│   └── Services - 存放服务相关
├── bootstrap - 启动引导,资源加载
│   ├── middleware.go - 中间件注册
│   ├── router.go - 路由注册
│   └── setup.go
├── config - 存放一些配置项、例如应用程序常量配置
│   ├── consts.go
│   └── database.go
├── database - 数据库
│   └── database.go
├── library
│   ├── env - 环境变量相关
│   │   └── env.go
│   ├── http - 请求库
│   │   ├── config.go
│   │   ├── download.go
│   │   ├── files.go
│   │   ├── header.go
│   │   └── request.go
│   └── utils - 常用工具
│       ├── conver.go
│       ├── crypto.go
│       ├── json.go
│       ├── net.go
│       ├── path.go
│       ├── print.go
│       └── time.go
└── resources - 资源文件夹
    ├── files - 可被下载的文件
    ├── mounter.go - 负责挂载资源文件夹
    ├── statics - 静态文件，比如 css，js
    │   ├── css
    │   │   └── element-ui.css
    │   └── js
    │       ├── element-ui.js
    │       ├── index.js
    │       └── vue.js
    └── views - 资源视图，比如 html 文件之类
        └── index.html

```