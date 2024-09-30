package models

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongDetails struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
