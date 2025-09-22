package main

import (
	"grep/2.18/utils"

	"grep/2.18/httpsh"
	"grep/2.18/serv"
	"grep/2.18/storage"
)

func main() {
	port := utils.Flagos()
	logger := utils.NewSlogger()
	storage := storage.NewMemoryStorage()
	service := serv.NewServiceImpl(storage, logger)
	handlers := httpsh.NewHandlerEvent(service, logger)
	server := httpsh.NewServer(*port, logger, service, storage, handlers)

	err := server.RunServ()
	if err != nil {
		logger.Error("Ошибка запуска сервера: ", err)
	}
}
