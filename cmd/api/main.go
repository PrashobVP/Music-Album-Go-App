package main

import (
    "log"
    "album-app/internal/album"
    "album-app/internal/db"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    database, err := db.Connect()
    if err != nil {
        log.Fatal("Could not connect to the database: ", err)
    }
    defer func() {
        sqlDB, err := database.DB()
        if err != nil {
            log.Fatal(err)
        }
        sqlDB.Close()
    }()

    // Auto migrate the Album model to create the table
    if err := database.AutoMigrate(&album.Album{}); err != nil {
        log.Fatal("Could not migrate database: ", err)
    }

    // Set up repository and service
    albumRepo := album.NewGormRepository(database)
    albumService := album.NewAlbumService(albumRepo)
    albumController := album.NewController(albumService)

    // Set up Gin router
    router := gin.Default()
    router.GET("/albums", albumController.GetAlbums)
    router.GET("/albums/:id", albumController.GetAlbumByID)
    router.POST("/albums", albumController.PostAlbum)
	router.PUT("/albums/:id", albumController.UpdateAlbum)
	router.DELETE("/albums/:id", albumController.DeleteAlbum)

    // Run server
    if err := router.Run("localhost:3001"); err != nil {
        log.Fatal("Could not run server: ", err)
    }
}
