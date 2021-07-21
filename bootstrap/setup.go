package bootstrap

import (
	"Pigeon/resources"
	"github.com/gin-gonic/gin"
)

// Setup 引导加载
func Setup(engine *gin.Engine) {
	// 注册业务路由
	RegisterRouter(
		// 资源目录挂载
		resources.Mount(
			// 路由开始前的中间件，可用于拦截请求之类
			UseMiddleware(engine)))
}
