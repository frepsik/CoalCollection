package main

import (
	"context"
	"sync"
	"time"
)

type Enterprise struct {
	coal             int
	mtx              sync.Mutex
	ctxEnterprise    context.Context
	cancelEnterprise context.CancelFunc
	wg               sync.WaitGroup
	startedAt        time.Time
	finishedAt       time.Time
}

// Конструктор для структуры Enterprise
func NewEnterprise(ctxApp context.Context) *Enterprise {

	//Создаём тут контекст предприятия на основании контекста программы всей
	ctxEnterprise, cancelEnterprise := context.WithCancel(ctxApp)

	return &Enterprise{
		ctxEnterprise:    ctxEnterprise,
		cancelEnterprise: cancelEnterprise,
	}
}

// Метод для добавления угля к общему балансу
func (e *Enterprise) AddCoal(amountCoal int) {
	e.mtx.Lock()
	e.coal += amountCoal
	e.mtx.Unlock()
}

// Метод для получения количества текущего угля
func (e *Enterprise) GetCoal() int {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return e.coal
}

// Метод для пассивного получения угля в секунду, запускаем предприятия этим
func (e *Enterprise) startPassive() {

	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				e.AddCoal(1)
			case <-e.ctxEnterprise.Done():
				return
			}
		}
	}()
}

// Метод, на возвращение того, сколько по итогу отработало предприятие
func (e *Enterprise) WorkTime() time.Duration {
	return e.finishedAt.Sub(e.startedAt) //Sub - берёт от текущешего времнеи и вычитает то, что передадим в аргументе
}

// Добавил навсякий случай, вдруг потом захочу расширить логику, например, что-то параллельно будет запускаться вместе с пасивным получением угля в секунду
func (e *Enterprise) Start() {
	e.startedAt = time.Now()
	e.startPassive()
}

// Метод завершения программы, пока временный, далее надо будет чуть более осмысленно сделать, потмоу что там нужны будут проверки
// на то что, всё ли оборудование куплено, сколько времени там затратилось на всё это
func (e *Enterprise) Shutdown() {
	e.finishedAt = time.Now()
	e.cancelEnterprise()
	e.wg.Wait()
}
