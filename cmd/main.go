package main

import (
	_ "github.com/jmoiron/sqlx"
	songLibs "github.com/kahuri1/song-libraries"
	"github.com/kahuri1/song-libraries/pkg/handler"
	"github.com/kahuri1/song-libraries/pkg/model"
	"github.com/kahuri1/song-libraries/pkg/repository"
	"github.com/kahuri1/song-libraries/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(model.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Errorf("Failed initialization db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.Newhandler(service)

	srv := new(songLibs.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs") //исправить на нормальный путь
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
