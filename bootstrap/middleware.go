package bootstrap

import (
	"github.com/gin-gonic/gin"
)

// 中间件编写流程：
// 在 app/Middlewares 里编写，在这里记录中间件函数
var middlewares []gin.HandlerFunc

func UseMiddleware(engine *gin.Engine) *gin.Engine {
	engine.Use(middlewares...)
	return engine
}
