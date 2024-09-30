package initT

import (
	"github.com/jackc/pgx/v5"
	"github.com/skinkvi/onlineBiblMusic/internal/config"
	"github.com/skinkvi/onlineBiblMusic/pkg/db"
	"go.uber.org/zap"
)

var (
	DbConn *pgx.Conn
	Logger *zap.Logger
)

func InitHandlers(cfg *config.Config) {
	DbConn = db.ConnectToDB(cfg)
	Logger = cfg.Logger
}
