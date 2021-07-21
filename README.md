# Pigeon
An HTTP service framework based on Gin encapsulation

#### 项目结构
```
├── README.md
├── app - 应用相关的文件
│   ├── Controllers - 路由控制器
│   │   ├── BaiDuWenKu
│   │   │   └── downloader.go
│   │   ├── Index.go - 处理根请求
│   │   └── Interface.go - 路由控制器接口
│   ├── Middlewares - 中间件
│   │   └── Resource.go
│   └── Services - 服务相关
│       └── BaiDuWenKu
│           ├── Cookies
│           │   └── 1.cookies
│           └── Downloader.go
├── config - 项目配置
│   ├── consts.go - 项目常量
│   └── database.go - 数据库配置
├── dao
│   └── database.go
├── go.mod
├── go.sum
├── library - 常用类库
│   ├── env - 环境变量
│   │   └── env.go
│   ├── http - 请求
│   │   ├── config.go
│   │   ├── download.go
│   │   ├── files.go
│   │   ├── header.go
│   │   └── request.go
│   └── utils - 工具
│       ├── conver.go
│       ├── path.go
│       └── time.go
├── main.go
├── resources - 资源文件夹
│   ├── files - 可被下载的文件
│   ├── mounter.go - 挂载资源文件夹
│   ├── statics - 静态文件，比如 css，js
│   │   ├── css
│   │   │   └── element-ui.css
│   │   └── js
│   │       ├── element-ui.js
│   │       ├── index.js
│   │       └── vue.js
│   └── views - 资源视图，比如 html 文件之类
│       └── index.html
└── bootstrap - 资源加载类
    ├── middleware.go - 中间件注册
    ├── router.go - 路由注册
    └── setup.go - 资源配置加载


```