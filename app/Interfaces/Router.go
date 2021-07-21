package Interfaces

import (
	"github.com/gin-gonic/gin"
)

type RouterController interface {
	RegisterRouter(*gin.RouterGroup)
}

type RouterControllerList []RouterController

func SetupRouter(routerGroup *gin.RouterGroup, controllers RouterControllerList) *gin.RouterGroup {
	for _, controller := range controllers {
		controller.RegisterRouter(routerGroup)
	}
	return routerGroup
}
