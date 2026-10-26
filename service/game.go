package service

import (
	"CoalCollection/domain"
	"context"
	"sync"
	"time"
)

type Game struct {
	ctxGame        context.Context
	cancelGame     context.CancelFunc
	wg             sync.WaitGroup
	minerTypes     map[domain.MinerTypeName]domain.MinerType
	equipmentTypes map[domain.EquipmentTypeName]domain.EquipmentType
	enterprise     *domain.Enterprise
	gameState      gameState
	mtx            sync.Mutex
}

// Конструктор создания экземпляра самой игры
func NewGame(
	ctx context.Context,
	equipmentTypes map[domain.EquipmentTypeName]domain.EquipmentType,
	minerTypes map[domain.MinerTypeName]domain.MinerType,
) *Game {
	ctxGame, cancelGame := context.WithCancel(ctx)

	return &Game{
		ctxGame:        ctxGame,
		cancelGame:     cancelGame,
		minerTypes:     minerTypes,
		equipmentTypes: equipmentTypes,
		enterprise:     domain.NewEnterprise(equipmentTypes),
		gameState:      gameState{},
	}
}

// Метод отвечающий за старт игры, где мы даём начало времени и запускаем пассивное получение угля
func (g *Game) Start() error {
	if g.gameState.currentState() {
		return ErrGameIsStarted
	}

	g.gameState.start()

	g.enterprise.Start(time.Now())
	g.startPassiveIncome()
	return nil
}

// Метод завершения программы, пока временный, далее надо будет чуть более осмысленно сделать, потому что там нужны будут проверки
// на то что, всё ли оборудование куплено, сколько времени там потрачено на всё это
func (g *Game) Shutdown() error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if !g.gameState.currentState() {
		return ErrGameIsNotStarted
	}

	g.gameState.finish()

	//Фиксируем окончание игры сначала, то есть тот момент, когда пользователь подал сигнал на завершение
	finishedAt := time.Now()

	g.cancelGame()
	g.wg.Wait()

	g.enterprise.Finish(finishedAt)

	return nil
}
