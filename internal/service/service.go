package service

import (
	"github.com/yerlan-tleubekov/go-redis/internal/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository}
}
