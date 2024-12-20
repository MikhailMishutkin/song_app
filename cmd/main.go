package main

import (
	"github.com/joho/godotenv"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"song_app/configs"
	"song_app/internal/app"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title			Song Service
//	@version		1.0
//	@description	API for songs info service

//	@host		localhost:8080
//	@BasePath	/

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.StartService(*config); err != nil {
		log.Fatal(err)
	}
}
