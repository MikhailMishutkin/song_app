package transport

import (
	"context"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"net/http"
	"song_app/internal/models"
	"song_app/pkg/middleware"
)

type HTTPSongHandle struct {
	sh HTTPSongManager
}

func NewHTTPSongHandle(sh HTTPSongManager) *HTTPSongHandle {
	return &HTTPSongHandle{
		sh: sh,
	}
}

type HTTPSongManager interface {
	CreateSong(context.Context, *models.Song) error
	GetSongText(context.Context, *models.FiltAndPagin) (*models.Song, error)
	UpdateSong(context.Context, *models.Song) error
	DeleteSong(context.Context, *models.Song) error
	GetAllSongs(context.Context, *models.FiltAndPagin) ([]*models.Song, error)
}

func (h *HTTPSongHandle) RegisterSong(router *mux.Router) {
	router.Use(middleware.Logging)
	router.Handle("/create", middleware.Logging(http.HandlerFunc(h.CreateSong))).Methods("POST")
	router.HandleFunc("/edit/{id:[0-9]+}", h.UpdateSong).Methods("PUT")
	router.HandleFunc("/delete/{id:[0-9]+}", h.DeleteSong).Methods("DELETE")
	router.HandleFunc("/gettext/{id:[0-9]+}", h.GetSongText).Methods("GET")
	router.HandleFunc("/getall", h.GetAllSongs).Methods("Get")

}
