package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/skinkvi/onlineBiblMusic/internal/models"
)

func fetchSongDetails(group, song string) (*models.SongDetails, error) {
	url := "http://example.com/info?group=" + group + "&song=" + song
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var songDetail models.SongDetails
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return &songDetail, nil
}
