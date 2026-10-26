package main

import (
	"CoalCollection/http"
	"context"
	"fmt"
)

func main() {

	rootCtx := context.Background()

	// Создание слоя gameService
	game, err := buildGame(rootCtx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создание handlers
	handlers := http.NewHttpHandlers(game)

	// Создание http сервера
	httpServer := http.NewHttpServer(handlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("fail start server:", err)
		return
	}
}
