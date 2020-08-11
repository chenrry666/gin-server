package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chenrry666/gin-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	f, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_RDWR, os.ModeAppend)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	r := gin.Default()
	{
		route := routes.Routes{r}
		route.Static().Route()
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
