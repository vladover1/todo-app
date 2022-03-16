package main

import (
	"github.com/vladover1/todo-app"
	"github.com/vladover1/todo-app/pkg/handler"
	"log"
)

func main()  {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8888", handler.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}


