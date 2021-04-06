package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yerlan-tleubekov/go-redis/internal/service"
)

type Authorization interface {
	SignIn(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services}
}

func (h *Handler) InitHandler() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/auth/sign-up", h.SignUp)
	r.HandleFunc("/auth/sign-in", h.SignIn)

	return r
}
