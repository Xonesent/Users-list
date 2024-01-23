package main

import (
	"users-list/internal/handler"
	"users-list/internal/repository"
	"users-list/internal/service"
	"users-list/server"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	repos := repository.New_Repository()
	services := service.New_Service(repos)
	handlers := handler.New_Handler(services)

	server := new(server.Server)
	server.Run("8080", handlers.Init_Routes())
}
