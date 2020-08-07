package main

import (
	"gin-server/routes"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	routes.Route(r)
	routes.Static(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
