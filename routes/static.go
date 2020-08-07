package routes

import "github.com/gin-gonic/gin"

func Static(r *gin.Engine) {
	r.Static("/public", "../public")
}
