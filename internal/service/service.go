package service

import (
	"context"
	"log/slog"
	"os"

	"song_app/internal/models"
)

type SongService struct {
	sr  SongRepositorier
	log *slog.Logger
}

func NewSongService(sr SongRepositorier) *SongService {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	return &SongService{sr: sr, log: logger}
}

type SongRepositorier interface {
	CreateGroup(context.Context, *models.Song) (groupId int, err error)
	CreateSong(context.Context, *models.Song) (songId int, err error)
	CreateSongUniqRec(context.Context, int, int) (uniqRecId int, err error)
	AddDetails(context.Context, *models.Song) error
	UpdateSong(context.Context, *models.Song) error
	DeleteSong(context.Context, *models.Song) error
	GetAllSongs(context.Context, *models.FiltAndPagin) ([]*models.Song, error)
	GetSongText(context.Context, *models.FiltAndPagin) (*models.Song, error)
	CheckExistGroup(context.Context, *models.Song) (groupId int, err error)
	CheckExistSong(context.Context, *models.Song) (songId int, err error)
	CheckExistSongUniq(context.Context, int, int) (uniqRecId int, err error)
}
