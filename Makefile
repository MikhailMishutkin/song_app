.PHONY: build
db:
	docker-compose up -d --build

migrateup:
	migrate -path migrations -database "postgres://root:root@localhost:5444/song_app?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://root:root@localhost:5444/song_app?sslmode=disable" -verbose down

http:
	go run ./cmd/service/main.go


build: db migrateup http


.DEFAULT_GOAL := build