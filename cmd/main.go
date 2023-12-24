package main

import (
	"log"

	_ "github.com/lib/pq"
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

	db, err := repository.NewPosgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
