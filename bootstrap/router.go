package bootstrap

import (
	"Pigeon/app/Controllers"
	"Pigeon/app/Interfaces"
	"github.com/gin-gonic/gin"
)

// 控制器编写流程：
// 在 app/Controller 里编写，在这里注册
var routers = Interfaces.RouterControllerList{
	new(Controllers.Index),
}

func RegisterRouter(router *gin.Engine) *gin.Engine {
	// 追加 404
	routers = append(routers, new(Controllers.NotFound))
	// 把所有路由都挂载到根路由下
	Interfaces.SetupRouter(router.Group("/"), routers)
	return router
}
