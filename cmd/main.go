package main

import (
	"github.com/spf13/viper"
	"github.com/vladover1/todo-app"
	"github.com/vladover1/todo-app/pkg/handler"
	"github.com/vladover1/todo-app/pkg/repository"
	"github.com/vladover1/todo-app/pkg/service"
	"log"
)

func main()  {
	if err:= initConfig(); err !=nil {
		log.Fatalf("error initializating: %s", err.Error())
	}
	
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}



