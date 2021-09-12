package bootstrap

import (
	"Pigeon/app/Caches/App"
	"Pigeon/resources"
	"github.com/gin-gonic/gin"
)

// Setup 引导加载
func Setup(engine *gin.Engine) {
	App.Engine = engine
	// 资源目录挂载
	resources.Mount(engine)
	// 路由开始前的中间件，可用于拦截请求之类
	UseMiddleware(engine)
	// 注册业务路由
	RegisterRouter(engine)
}
