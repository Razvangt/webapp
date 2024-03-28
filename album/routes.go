package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Tags albums
// @Accept json
// @Produce json
// @Router /albums [get]
func GetAlbumsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Map(all(), Response))
}

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Tags albums
// @Accept json
// @Produce json
// @Param request body newAlbumRequest  true  "query params"
// @Router /albums [post]
func PostAlbumsHandler(c *gin.Context) {
	var newAlbum newAlbumRequest

	// call bindjson to bind the received json to
	// new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	add(validateCreate(&newAlbum))
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Tags albums
// @Accept json
// @Produce json
// @Param id path  string true "id"
// @Router /albums/{id} [get]
func GetAlbumByIDHandler(c *gin.Context) {
	id := c.Param("id")

	album, err := find(&id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, Response(album))
}

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Tags albums
// @Accept json
// @Produce json
// @Param id path  string true "id"
// @Router /albums/{id} [delete]
func DeleteAlbumByIDHandler(c *gin.Context) {
	id := c.Param("id")
	err := delete(&id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted"})
}

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Tags albums
// @Accept json
// @Produce json
// @Param id path  string true "id"
// @Param request body newAlbumRequest true  "query params"
// @Router /albums/{id} [put]
func UpdateAlbumByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum newAlbumRequest
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	if err := update(validateUpdate(&updatedAlbum, id)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, updatedAlbum)
}

func AddGroup(v1 *gin.RouterGroup) *gin.RouterGroup {
	router := v1.Group("/albums")
	{
		router.GET("", GetAlbumsHandler)
		router.GET("/:id", GetAlbumByIDHandler)
		router.POST("", PostAlbumsHandler)
		router.DELETE("/:id", DeleteAlbumByIDHandler)
		router.PUT("/:id", UpdateAlbumByIDHandler)
	}
	return router
}
