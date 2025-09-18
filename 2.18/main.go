package main

import (
	"grep/2.18/httpsh"
	"grep/2.18/serv"
	"grep/2.18/storage"
)

func main() {
	storage := storage.NewMemoryStorage()
	service := serv.NewServiceImpl(storage)
	handlers := httpsh.NewHandlerEvent(service)
}
