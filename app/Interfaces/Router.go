package Interfaces

import (
	"github.com/gin-gonic/gin"
)

type RouterController interface {
	RegisterRouter(*gin.RouterGroup)
}

type RoutePathControllerMap map[string]gin.RouteInfo

func MakeRouterMap(routes gin.RoutesInfo) (rm RoutePathControllerMap) {
	rm = make(RoutePathControllerMap, len(routes))
	for _, route := range routes {
		rm[route.Path] = route
	}
	return
}

type DomainRouterMap map[string]RouterController

type RouterControllerList []RouterController

func SetupRouter(routerGroup *gin.RouterGroup, controllers RouterControllerList) *gin.RouterGroup {
	for _, controller := range controllers {
		controller.RegisterRouter(routerGroup)
	}
	return routerGroup
}
