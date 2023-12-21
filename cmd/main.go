package main

import (
	"log"

	"github.com/spf13/viper"
	webback "github.com/zaim/web_back"
	"github.com/zaim/web_back/pkg/handler"
	"github.com/zaim/web_back/pkg/repository"
	"github.com/zaim/web_back/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while initial cfg: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(webback.Server)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
