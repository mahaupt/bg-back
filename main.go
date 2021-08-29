package main

import (
	"os"

	"github.com/mahaupt/bg-back/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// if failed, assume env variables are set manually
	godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}))
	r.Use(gin.Recovery())

	mc := new(controller.MessageController)
	r.GET("/v1/send", mc.SendMessage)
	r.GET("/health", mc.GetHealth)

	r.Run()
}
