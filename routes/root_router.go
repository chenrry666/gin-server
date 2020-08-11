package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine *gin.Engine
}

func (r *Routes) Route() *Routes {
	r.Engine.GET("/", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, gin.H{
			"code": c.DefaultQuery("query", "none"),
		})
	})
	return r
}
func (r *Routes) Static() *Routes {
	r.Engine.Static("/public", "../public")
	return r

}
