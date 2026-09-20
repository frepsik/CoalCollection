package main

import (
	"CoalCollection/service"
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Coal mine started")

	rootCtx := context.Background()

	gameService := service.NewGame(rootCtx)

	gameService.Start()

	time.Sleep(10 * time.Second)

	gameService.Shutdown()

	fmt.Println("Статус игры", gameService.StatusEnterprise())
	fmt.Println("Конец")
}
