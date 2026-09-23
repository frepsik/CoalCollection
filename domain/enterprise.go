package domain

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Coal int

type Enterprise struct {
	coal       Coal
	mtx        sync.Mutex
	startedAt  time.Time
	finishedAt time.Time
	equipments map[EquipmentType]*Equipment
	miners     map[uuid.UUID]*Miner
}

// Конструктор для структуры Enterprise
func NewEnterprise(equipments map[EquipmentType]*Equipment) *Enterprise {
	miners := make(map[uuid.UUID]*Miner)
	return &Enterprise{
		equipments: equipments,
		miners:     miners,
	}
}

// Метод для добавления угля к общему балансу
func (e *Enterprise) AddCoal(amountCoal Coal) {
	e.mtx.Lock()
	e.coal += amountCoal
	e.mtx.Unlock()
}

// Метод для получения количества текущего угля
func (e *Enterprise) Coal() Coal {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return e.coal
}

// Добавил навсякий случай, вдруг потом захочу расширить логику, например, что-то параллельно будет запускаться вместе с пасивным получением угля в секунду
func (e *Enterprise) Start(at time.Time) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.startedAt = at
}

func (e *Enterprise) Finish(at time.Time) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.finishedAt = at
}

// Метод, на возвращение того, сколько по итогу отработало предприятие
func (e *Enterprise) workTime(now time.Time) time.Duration {

	//На случай, если захотим запросить время работы, во время выполнения самой игры
	if e.finishedAt.IsZero() {
		return now.Sub(e.startedAt)
	}
	return e.finishedAt.Sub(e.startedAt) //Sub - берёт от текущешего времнеи и вычитает то, что передадим в аргументе
}

// Метод который будет собирать текущую информацию по работе предприятия, в последствии предполагаю надо дополнить
func (e *Enterprise) Status() EnterpriseStatus {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return EnterpriseStatus{
		Coal: e.coal, WorkTime: e.workTime(time.Now()),
	}
}

// Метод приобретения оборудования
func (e *Enterprise) BuyEquipment(equipment EquipmentType) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	eq, ok := e.equipments[equipment]
	if !ok {
		return ErrEquipmentNotFound
	}

	if eq.purchased {
		return ErrEquipmentAlreadyPurchased
	}

	if e.coal < eq.cost {
		return ErrEnterpriseNotEnoughCoal
	}

	e.coal -= eq.cost

	eq.purchased = true
	return nil
}

// Метод для получения всего оборудования
func (e *Enterprise) Equipments() []Equipment {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	result := make([]Equipment, 0, len(e.equipments))

	for _, equipment := range e.equipments {
		result = append(result, *equipment)
	}

	return result
}

// Метод для получения приобретённого оборудования
func (e *Enterprise) EquipmentsPurchased() []Equipment {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	result := make([]Equipment, 0, len(e.equipments))

	for _, equipment := range e.equipments {
		if equipment.purchased {
			result = append(result, *equipment)
		}
	}

	return result
}

// Метод для найма шахтёра
func (e *Enterprise) HireMiner(minerType MinerType) (*Miner, error) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	if e.coal < minerType.cost {
		return nil, ErrEnterpriseNotEnoughCoal
	}

	e.coal -= minerType.cost

	miner := newMiner(minerType)

	e.miners[miner.id] = miner

	return miner, nil
}

// Метод для получения шахтёров определённого типа
func (e *Enterprise) MinersByType(minerTypeName MinerTypeName) []MinerInfo {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	var result []MinerInfo

	for _, miner := range e.miners {
		if miner.isType(minerTypeName) {
			result = append(result, miner.info())
		}
	}

	return result
}

// Метод для получения не работающих шахтёров
// тут возможна логическа гонка данных: между miner.isExhausted() и miner.info() - может вклинитьс miner.Mine()
// Надо подумать как тут реализовать атомарную операцию
func (e *Enterprise) ExhaustedMiners() []MinerInfo {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	var result []MinerInfo

	for _, miner := range e.miners {
		if miner.isExhausted() {
			result = append(result, miner.info())
		}
	}

	return result
}

// Метод для получения работающих шахтёров
// тут возможна логическа гонка данных: между miner.isExhausted() и miner.info() - может вклинитьс miner.Mine()
// Надо подумать как тут реализовать атомарную операцию
func (e *Enterprise) AvailableMiners() []MinerInfo {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	var result []MinerInfo

	for _, miner := range e.miners {
		if !(miner.isExhausted()) {
			result = append(result, miner.info())
		}
	}

	return result
}
