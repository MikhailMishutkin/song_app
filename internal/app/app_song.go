package app

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"log/slog"
	"net/http"
	"song_app/configs"
	"song_app/internal/song_app/repository"
	"song_app/internal/song_app/service"
	"song_app/internal/song_app/transport"
	"song_app/pkg/middleware"
)

type SongServer struct {
	songRouter *mux.Router
	logger     *slog.Logger
	swagRouter *mux.Router
}

func StartService(conf configs.Config) error {
	s := &SongServer{
		songRouter: mux.NewRouter(),
		logger:     slog.Default(),
		swagRouter: mux.NewRouter(),
	}

	db, err := NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	//httpserver
	repo := repository.NewRepo(db)
	songService := service.NewSongService(repo)
	songHandler := transport.NewHTTPSongHandle(songService)
	songHandler.RegisterSong(s.songRouter)
	s.songRouter.Use(middleware.Logging)
	s.logger.Info("Starting MessageService at port" + fmt.Sprintf("%v", conf.Host))
	return http.ListenAndServe(fmt.Sprintf("%v", conf.Host), s)
}

func NewDB() (*pgx.Conn, error) {
	c, err := configs.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("Can't load db config: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.Conn)

	db, err := pgx.Connect(context.Background(), psqlInfo)

	m, err := migrate.New(
		"file://../song_service/migrations",
		"postgres://"+c.Migrate,
		//root:root@localhost:5444/song_service?sslmode=disable",
	)
	if err != nil {
		log.Println(err)
		return db, fmt.Errorf("can't automigrate: %v\n", err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
		fmt.Errorf("%v\n", err)
	}
	return db, err
}

// ServeHTTP
func (h *SongServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.songRouter.ServeHTTP(w, r)

}
