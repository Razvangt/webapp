package main

import (
	"github.com/Razvangt/webapp/album"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbumsHandler)
	router.GET("/albums/:id", album.GetAlbumByIDHandler)
	router.POST("/albums", album.PostAlbumsHandler)
	router.DELETE("/albums:id", album.DeleteAlbumByIDHandler)
	router.PUT("albums:id", album.UpdateAlbumByIDHandler)
	router.Run() // listen and serve localhost:8080
}
