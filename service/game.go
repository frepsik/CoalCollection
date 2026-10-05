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
}

// Конструктор создания экземпляра самой игры
func NewGame(
	ctx context.Context,
	equipmentTypes map[domain.EquipmentTypeName]domain.EquipmentType,
	minerTypes map[domain.MinerTypeName]domain.MinerType,
) *Game {
	ctxGame, cancelGame := context.WithCancel(ctx)

	equipments := make(map[domain.EquipmentTypeName]*domain.Equipment, len(equipmentTypes))
	for equipmentTypeName, equipmentType := range equipmentTypes {
		equipments[equipmentTypeName] = domain.NewEquipment(equipmentType)
	}

	return &Game{
		ctxGame:        ctxGame,
		cancelGame:     cancelGame,
		minerTypes:     minerTypes,
		equipmentTypes: equipmentTypes,
		enterprise:     domain.NewEnterprise(equipments),
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
