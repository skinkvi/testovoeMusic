package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/onlineBiblMusic/internal/models"
	initT "github.com/skinkvi/onlineBiblMusic/pkg/init"
)

// GetSongById godoc
// @Summary      Get song by id
// @Description  Get song by id
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Song ID"
// @Success      200  {object}  models.Song
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /songs/{id} [get]
func GetSongById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		initT.Logger.Error("Error getting song by id")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var song models.Song
	err = initT.DbConn.QueryRow(context.Background(),
		"SELECT * FROM songs WHERE id = $1", id).Scan(
		&song.ID, &song.Song, &song.Group, &song.Text, &song.Link, &song.ReleaseDate)
	if err != nil {
		initT.Logger.Error("Error getting song by id")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	initT.Logger.Info("Succesfully got song by id")
	ctx.JSON(http.StatusOK, song)
}
