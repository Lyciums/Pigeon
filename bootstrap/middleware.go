package bootstrap

import (
	"Pigeon/app/Middlewares/Router"
	"github.com/gin-gonic/gin"
)

// 中间件编写流程：
// 在 app/Middlewares 里编写，在这里记录中间件函数
var middlewares = []gin.HandlerFunc{
	Router.DomainRouter,
}

func UseMiddleware(engine *gin.Engine) *gin.Engine {
	engine.Use(middlewares...)
	return engine
}
