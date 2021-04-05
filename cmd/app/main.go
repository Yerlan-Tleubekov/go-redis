package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yerlan-tleubekov/go-redis/internal/repository"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"github.com/yerlan-tleubekov/go-redis/internal/cache"
	"github.com/yerlan-tleubekov/go-redis/internal/handler"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	memCache := cache.NewRedisCache(host + ":" + port)

	repo := repository.NewRepository(0, memCache)

	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Main)

	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
