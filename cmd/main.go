package main

import (
	"log"

	"github.com/AhmAlgiz/marketplace"
	"github.com/AhmAlgiz/marketplace/pkg/handler"
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config file %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(marketplace.Server)
	if err := s.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
