package handler

import (
	"fmt"
	"net/http"

	"github.com/yerlan-tleubekov/go-redis/internal/cache"
)

var (
	tokenCache cache.TokenCache
)

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it works")
}
