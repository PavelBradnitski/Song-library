package handlers

import (
	"Songs/Song-library/internal/models"
	"Songs/Song-library/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	service *services.SongService
}

func NewSongHandler(service *services.SongService) *SongHandler {
	return &SongHandler{service: service}
}

func (h *SongHandler) RegisterRoutes(router *gin.Engine) {
	songGroup := router.Group("/songs")
	{
		songGroup.POST("/", h.CreateSong)
		songGroup.GET("/", h.GetAllSongs)
		songGroup.GET("/:group/:song", h.GetSong)
		songGroup.PUT("/:group/:song", h.UpdateSong)
		songGroup.DELETE("/:group/:song", h.DeleteSong)
	}
}

func (h *SongHandler) CreateSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	if err := h.service.CreateSong(ctx, &song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

func (h *SongHandler) GetAllSongs(c *gin.Context) {
	songs, err := h.service.GetAllSongs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongHandler) GetSong(c *gin.Context) {
	group := c.Param("group")
	song := c.Param("song")

	foundSong, err := h.service.GetSong(c, group, song)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.JSON(http.StatusOK, foundSong)
}

func (h *SongHandler) UpdateSong(c *gin.Context) {
	group := c.Param("group")
	song := c.Param("song")

	var updatedSong models.Song
	if err := c.ShouldBindJSON(&updatedSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingSong, err := h.service.GetSong(c, group, song)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	existingSong.Text = updatedSong.Text
	existingSong.Link = updatedSong.Link
	existingSong.ReleaseDate = updatedSong.ReleaseDate

	if err := h.service.UpdateSong(c, existingSong); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
		return
	}

	c.JSON(http.StatusOK, existingSong)
}

func (h *SongHandler) DeleteSong(c *gin.Context) {
	group := c.Param("group")
	song := c.Param("song")

	if err := h.service.DeleteSong(c, group, song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
