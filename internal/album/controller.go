package album

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller for album HTTP endpoints.
type Controller struct {
    service Service
}

// NewController creates a new album controller.
func NewController(service Service) *Controller {
    return &Controller{service}
}

// GetAlbums responds with the list of all albums as JSON.
func (c *Controller) GetAlbums(ctx *gin.Context) {
    albums, err := c.service.GetAllAlbums()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, albums)
}

// GetAlbumByID responds with an album by ID as JSON.
func (c *Controller) GetAlbumByID(ctx *gin.Context) {
    id := ctx.Param("id")
    album, err := c.service.GetAlbumByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
        return
    }
    ctx.JSON(http.StatusOK, album)
}

// PostAlbum creates a new album from JSON received in the request body.
func (c *Controller) PostAlbum(ctx *gin.Context) {
    var newAlbum Album
    if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.CreateAlbum(newAlbum); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, newAlbum)
}

// UpdateAlbum updates an existing album from JSON received in the request body.
// UpdateAlbum updates an existing album from JSON received in the request body.
func (c *Controller) UpdateAlbum(ctx *gin.Context) {
    var updatedAlbum Album

    // Bind JSON to the updatedAlbum variable
    if err := ctx.ShouldBindJSON(&updatedAlbum); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate that the album ID is present
    if updatedAlbum.ID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Album ID is required"})
        return
    }

    // Update the album in the service
    if err := c.service.UpdateAlbum(updatedAlbum); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        return
    }

    // Return the updated album
    ctx.JSON(http.StatusOK, updatedAlbum)
}



// DeleteAlbum deletes an album by ID from the request URL.
func (c *Controller) DeleteAlbum(ctx *gin.Context) {
    // Get the album ID from the URL parameter
    albumID := ctx.Param("id")

    // Validate that the album ID is present
    if albumID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Album ID is required"})
        return
    }

    // Call the service to delete the album
    if err := c.service.DeleteAlbum(albumID); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        
        return
    }

    // Respond with a 204 No Content status (successful delete with no content)
    ctx.Status(http.StatusNoContent)
    
}



