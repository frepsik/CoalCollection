package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Coal mine started")
	ctx, cancell := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	enterprise := Enterprise{
		ctx: ctx,
		wg:  wg,
	}

	fmt.Println(enterprise.GetCoal())
	enterprise.AddCoal(10)

	fmt.Println(enterprise.GetCoal())

	enterprise.StartPassive()

	time.Sleep(time.Second * 10)

	cancell()

	wg.Wait()
	fmt.Println(enterprise.GetCoal())
	fmt.Println("Конец")
}
