package main

import (
	"context"
	"sync"
	"time"
)

type Enterprise struct {
	coal int
	mtx  sync.Mutex
	ctx  context.Context
	wg   *sync.WaitGroup
}

func (e *Enterprise) AddCoal(amountCoal int) {
	e.mtx.Lock()
	e.coal += amountCoal
	e.mtx.Unlock()
}

func (e *Enterprise) GetCoal() int {
	defer e.mtx.Unlock()
	e.mtx.Lock()
	return e.coal

}

func (e *Enterprise) StartPassive() {

	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				e.AddCoal(1)
			case <-e.ctx.Done():
				return
			}
		}
	}()
}
