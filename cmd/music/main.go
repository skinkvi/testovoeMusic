package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skinkvi/onlineBiblMusic/internal/config"
	"github.com/skinkvi/onlineBiblMusic/internal/handlers"
	"github.com/skinkvi/onlineBiblMusic/pkg/db"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil {
		cfg.Logger.Fatal("Error loading .env file")
	}

	dbC := db.ConnectToDB(cfg)
	if dbC != nil {
		cfg.Logger.Error("Error connecting to DB")
	}
	defer db.CloseConn(dbC)
	cfg.Logger.Info("Connected to DB")

	r := gin.Default()

	r.GET("library", handlers.GetLibrary)
	r.GET()

	port, exists := os.LookupEnv("SERVER_PORT")
	if exists {
		log.Printf("Listening on port %s", port)
		http.ListenAndServe(":"+port, nil)
	}
}
