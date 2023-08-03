package main

import (
	"log"

	"github.com/AhmAlgiz/marketplace"
	"github.com/AhmAlgiz/marketplace/pkg/handler"
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(marketplace.Server)
	if err := s.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err)
	}
}
