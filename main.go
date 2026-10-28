package main

import (
	"CoalCollection/httpapi"
	"context"
	"log"
)

func main() {

	rootCtx := context.Background()

	// Создание слоя gameService
	game, err := buildGame(rootCtx)
	if err != nil {
		log.Println("fail build program", err)
		return
	}

	// Создание handlers
	handlers := httpapi.NewHandler(game)

	// Создание http сервера
	httpServer := httpapi.NewServer(handlers)

	if err := httpServer.Start(); err != nil {
		log.Println("fail start server:", err)
		return
	}
}
