package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/onlineBiblMusic/internal/models"
	initT "github.com/skinkvi/onlineBiblMusic/pkg/init"
	"go.uber.org/zap"
)

// GetLibrary godoc
// @Summary      Get library of songs
// @Description  Get all songs from library
// @Tags         library
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Song
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Failure      500  {object} gin.H
// @Router       /library [get]
func GetLibrary(ctx *gin.Context) {
	var songs []models.Song
	rows, err := initT.DbConn.Query(context.Background(), "SELECT * FROM songs")
	if err != nil {
		initT.Logger.Error("Error getting songs from library", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Song, &song.Group, &song.Text, &song.Link, &song.ReleaseDate)
		if err != nil {
			initT.Logger.Error("Error getting songs from library", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		initT.Logger.Error("Error getting songs from library", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	initT.Logger.Info("Successfully got songs from library")
	ctx.JSON(http.StatusOK, songs)
}
