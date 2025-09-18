package main

import (
	"grep/2.18/utils"

	"grep/2.18/httpsh"
	"grep/2.18/serv"
	"grep/2.18/storage"
)

func main() {
	logger := utils.NewSlogger()
	storage := storage.NewMemoryStorage()
	service := serv.NewServiceImpl(storage)
	handlers := httpsh.NewHandlerEvent(service)
	server := httpsh.NewServer("8081", logger, service, storage, handlers)

	err := server.RunServ()
	if err != nil {
		logger.Error("Ошибка запуска сервера: ", err)
	}
}
