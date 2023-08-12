package main

import (
	"log"
	"os"

	"github.com/AhmAlgiz/marketplace"
	"github.com/AhmAlgiz/marketplace/pkg/handler"
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(marketplace.Server)
	if err := s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
