package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbumsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbumsHandler(c *gin.Context) {
	var newAlbum album

	// call bindjson to bind the received json to
	// new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	add(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByIDHandler(c *gin.Context) {
	id := c.Param("id")

	album, err := find(&id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbumByIDHandler(c *gin.Context) {
	id := c.Param("id")
	err := delete(&id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted"})
}

func UpdateAlbumByIDHandler(c *gin.Context) {
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	if err := update(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, updatedAlbum)
}
