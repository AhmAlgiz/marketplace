package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AhmAlgiz/marketplace"
	"github.com/AhmAlgiz/marketplace/pkg/handler"
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
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
		logrus.Fatalf("failed to initialize DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(marketplace.Server)
	go func() {
		if err := s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Server running error: %s", err.Error())
		}
	}()

	logrus.Print("Marketplace started.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Marketplace shutting down.")

	if err := s.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error closing database: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
