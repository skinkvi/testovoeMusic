package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/onlineBiblMusic/internal/models"
	"github.com/skinkvi/onlineBiblMusic/pkg/init"
	"go.uber.org/zap"
)

// AddSong godoc
// @Summary      Add song
// @Description  Add song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Song
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /songs [post]
func AddSong(ctx *gin.Context) {
	var song models.Song

	if err := ctx.ShouldBindJSON(&song); err != nil {
		initT.Logger.Error("Error binding JSON", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	songDetails, err := fetchSongDetails(song.Group, song.Song)
	if err != nil {
		initT.Logger.Error("Error fetching song details", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	song.ReleaseDate = songDetails.ReleaseDate
	song.Text = songDetails.Text
	song.Link = songDetails.Link
	_, err = initT.DbConn.Exec(context.Background(),
		"INSERT INTO songs (song, group_name, text, link, release_date) VALUES ($1, $2, $3, $4, $5)",
		song.Song, song.Group, song.Text, song.Link, song.ReleaseDate)
	if err != nil {
		initT.Logger.Error("Error adding song", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	initT.Logger.Info("Succesfully added song")
	ctx.JSON(http.StatusOK, song)
}
