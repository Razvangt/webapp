package main

import (
	"log"
	"os"
	"webapp/album"
	"webapp/common"
	docs "webapp/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Router
	db := common.Init()
	db.AutoMigrate(&album.Album{})
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		album.AddGroup(v1)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api_port := ":" + os.Getenv("API_PORT")
	router.Run(api_port) // listen and serve localhost:8080
}
