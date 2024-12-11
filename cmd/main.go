package main

import (
	"Songs/Song-library/internal/handlers"
	"Songs/Song-library/internal/repositories"
	"Songs/Song-library/internal/services"
	"Songs/Song-library/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "postgresql://users:1234@localhost:5432/library"
	database := db.ConnectDB(dsn)
	defer database.Close()

	songRepo := repositories.NewSongRepository(database)
	songService := services.NewSongService(songRepo)
	songHandler := handlers.NewSongHandler(songService)

	router := gin.Default()
	songHandler.RegisterRoutes(router)

	router.Run(":8080")
}
