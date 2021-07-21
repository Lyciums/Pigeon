package Controllers

import (
	"github.com/gin-gonic/gin"
)

type Index struct{}

func (i *Index) RegisterRouter(r *gin.RouterGroup) {
	r.GET("", i.Index)
}

func (i *Index) Index(c *gin.Context) {
	c.String(200, "index.")
}
