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
	lifecycleMu    sync.RWMutex
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
		gameState:      gameCreated,
	}
}

// Метод отвечающий за старт игры, где мы даём начало времени и запускаем пассивное получение угля
func (g *Game) Start() error {
	g.lifecycleMu.Lock()
	defer g.lifecycleMu.Unlock()

	if g.gameState != gameCreated {
		return ErrGameCannotStart
	}
	g.gameState = gameRunning

	g.enterprise.Start(time.Now())
	g.startPassiveIncome()
	return nil
}

// Метод завершения программы, пока временный, далее надо будет чуть более осмысленно сделать, потому что там нужны будут проверки
// на то что, всё ли оборудование куплено, сколько времени там потрачено на всё это
func (g *Game) Shutdown() error {
	g.lifecycleMu.Lock()
	defer g.lifecycleMu.Unlock()

	if g.gameState != gameRunning {
		return ErrGameCannotFinish
	}
	g.gameState = gameFinished

	//Фиксируем окончание игры сначала, то есть тот момент, когда пользователь подал сигнал на завершение
	finishedAt := time.Now()

	g.cancelGame()
	g.wg.Wait()

	g.enterprise.Finish(finishedAt)

	return nil
}
