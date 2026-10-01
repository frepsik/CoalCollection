package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Coal mine started")

	rootCtx := context.Background()

	game, err := buildGame(rootCtx)
	if err != nil {
		fmt.Println(err)
	}

	game.Start()

	time.Sleep(10 * time.Second)

	game.Shutdown()

	fmt.Println("Статус игры", game.StatusEnterprise())
	fmt.Println("Конец")
}
