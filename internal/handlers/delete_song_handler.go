package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	initT "github.com/skinkvi/onlineBiblMusic/pkg/init"
)

// DeleteSong godoc
// @Summary      Delete song
// @Description  Delete song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Song ID"
// @Success      200  {obj200ect}  gin.H
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /songs/{id} [delete]
func DeleteSong(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		initT.Logger.Error("Error deleting song")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = initT.DbConn.Exec(context.Background(), "DELETE FROM songs WHERE id = $1", id)
	if err != nil {
		initT.Logger.Error("Error deleting song")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	initT.Logger.Info("Succesfully deleted song")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Succesfully deleted song",
	})
}
