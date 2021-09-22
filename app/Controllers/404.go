package Controllers

import (
	"github.com/gin-gonic/gin"
)

type NotFound struct{}

func (i *NotFound) RegisterRouter(r *gin.RouterGroup) {
	r.GET(":route", i.NotFoundRoute)
	r.GET(":route/*option", i.NotFoundRoutePath)
}

func (i *NotFound) NotFoundRoute(c *gin.Context) {
	c.String(200, "404 not found:"+c.Param("route"))
}

func (i *NotFound) NotFoundRoutePath(c *gin.Context) {
	c.String(200, "404 not found:"+c.Param("route")+"/"+c.Param("option"))
}
