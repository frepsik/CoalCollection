package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Coal mine started")

	rootCtx := context.Background()

	enterprise := NewEnterprise(rootCtx)

	enterprise.Start()

	time.Sleep(10 * time.Second)

	enterprise.Shutdown()

	fmt.Println("Количества угля:", enterprise.GetCoal(), "Время работы предприятия:", enterprise.WorkTime())
	fmt.Println("Конец")
}
