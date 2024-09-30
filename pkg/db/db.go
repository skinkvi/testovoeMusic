package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/skinkvi/onlineBiblMusic/internal/config"
	"go.uber.org/zap"
)

func ConnectToDB(cfg *config.Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		cfg.Logger.Fatal("Error connecting to DB", zap.Error(err))
		return nil
	}
	cfg.Logger.Info("Connected to DB")
	return conn
}

func CloseConn(db *pgx.Conn) {
	db.Close(context.Background())
}
