package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yerlan-tleubekov/go-redis/internal/handler"

	"github.com/joho/godotenv"
	"github.com/yerlan-tleubekov/go-redis/internal/cache"
	"github.com/yerlan-tleubekov/go-redis/internal/repository"
	"github.com/yerlan-tleubekov/go-redis/internal/service"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load()

	host := "127.0.0.1"
	// host := os.Getenv("HOST")
	// port := os.Getenv("PORT")

	memCache := cache.NewRedisCache(ctx, host+":6379", 0, time.Minute*30)

	repo := repository.NewRepository(0, memCache)
	serv := service.NewService(repo)
	hand := handler.NewHandler(serv)

	if err != nil {
		log.Fatal("error loading .env file")
	}

	server := &http.Server{
		Handler:      hand.InitHandler(),
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server is listening...")
	log.Fatal(server.ListenAndServe())
}
