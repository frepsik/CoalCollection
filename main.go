package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	rootCtx := context.Background()

	game, err := buildGame(rootCtx)
	if err != nil {
		fmt.Println(err)
		return
	}

	game.Start()
	fmt.Println("Coal mine started")
	time.Sleep(10 * time.Second)

	game.Shutdown()

	fmt.Println("Статус игры", game.StatusEnterprise())
	fmt.Println("Конец")
}
