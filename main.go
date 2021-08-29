package main

import (
	"log"
	"os"

	"github.com/mahaupt/bg-back/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	{
		mc := new(controller.MessageController)
		v1.GET("/send", mc.SendMessage)
	}

	r.Run()
}
