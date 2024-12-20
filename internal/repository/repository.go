package repository

import (
	"github.com/jackc/pgx/v5"
	"log/slog"
	"os"
)

type Repo struct {
	DB  *pgx.Conn
	log *slog.Logger
}

func NewRepo(
	db *pgx.Conn,
) *Repo {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	return &Repo{
		DB:  db,
		log: logger,
	}
}
