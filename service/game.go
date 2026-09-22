package service

import (
	"CoalCollection/domain"
	"context"
	"sync"
	"time"
)

type Game struct {
	ctxGame    context.Context
	cancelGame context.CancelFunc
	wg         sync.WaitGroup
	enterprise *domain.Enterprise
}

// Конструктор создания экземпляра самой игры
func NewGame(ctx context.Context) *Game {
	ctxGame, cancelGame := context.WithCancel(ctx)

	equipments := map[domain.EquipmentType]*domain.Equipment{
		domain.Pickaxe: {
			// name: "Кирка",
			// Cost: 3000,
		},
		domain.Ventilation: {
			// Name: "Вентиляция",
			// Cost: 15000,
		},
		domain.Trolleys: {
			// Name: "Вагонетки",
			// Cost: 50000,
		},
	}

	return &Game{
		ctxGame:    ctxGame,
		cancelGame: cancelGame,
		enterprise: domain.NewEnterprise(equipments),
	}
}

// Запуск пассивного получения угля
func (g *Game) startPassiveIncome() {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				g.enterprise.AddCoal(1)
			case <-g.ctxGame.Done():
				return
			}
		}
	}()
}

// Метод отвечающий за старт игры, где мы даём начало времени и запускаем пассивное получение угля
func (g *Game) Start() {
	g.enterprise.Start(time.Now())
	g.startPassiveIncome()
}

// Метод завершения программы, пока временный, далее надо будет чуть более осмысленно сделать, потмоу что там нужны будут проверки
// на то что, всё ли оборудование куплено, сколько времени там затратилось на всё это
func (g *Game) Shutdown() {
	//Фиксируем окончание игры сначал, то есть тот момент, когда пользователь подал сигнал на завершение
	finisheAt := time.Now()

	g.cancelGame()
	g.wg.Wait()

	g.enterprise.Finish(finisheAt)
}

// Метод позволяющий узнать, текущий статус, по тому, сколько идёт игра, и сколько сейчас угля, сюда дальше ещё надо интегрировать шахтёров и оборудование, но пока временно так
func (g *Game) StatusEnterprise() domain.EnterpriseStatus {
	return g.enterprise.Status()
}
