package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, gin.H{
			"code": c.DefaultQuery("query", "none"),
		})
	})
}
