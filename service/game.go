package service

import (
	"CoalCollection/domain"
	"context"
	"errors"
	"sync"
	"time"
)

type Game struct {
	ctxGame    context.Context
	cancelGame context.CancelFunc
	wg         sync.WaitGroup
	enterprise *domain.Enterprise
	minerTypes map[domain.MinerTypeName]domain.MinerType
}

// Конструктор создания экземпляра самой игры
func NewGame(
	ctx context.Context,
	equipments map[domain.EquipmentType]*domain.Equipment,
	minerTypes map[domain.MinerTypeName]domain.MinerType,
) *Game {
	ctxGame, cancelGame := context.WithCancel(ctx)

	return &Game{
		ctxGame:    ctxGame,
		cancelGame: cancelGame,
		minerTypes: minerTypes,
		enterprise: domain.NewEnterprise(equipments),
	}
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

func (g *Game) HireMiner(typeMiner string) error {
	minerTypeName := domain.MinerTypeName(typeMiner)

	minerType, exists := g.minerTypes[minerTypeName]
	if !exists {
		return errors.New("Invalid type miner name")
	}

	miner, err := g.enterprise.HireMiner(minerType)
	if err != nil {
		return err
	}
	ctxMiner, cancleMine := context.WithCancel(g.ctxGame)

	g.startMinerMine(miner, ctxMiner, cancleMine)
	return nil
}
