package domain

import (
	"errors"
	"sync"
	"time"
)

type Coal int

type Enterprise struct {
	coal       Coal
	mtx        sync.Mutex
	startedAt  time.Time
	finishedAt time.Time
	equipments map[EquipmentType]*Equipment
}

// Конструктор для структуры Enterprise
func NewEnterprise(equipments map[EquipmentType]*Equipment) *Enterprise {
	return &Enterprise{
		equipments: equipments,
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
func (e *Enterprise) WorkTime(now time.Time) time.Duration {
	//На случай, если захотим запросить время работы, во время выполнения самой игры
	if e.finishedAt.IsZero() {
		return now.Sub(e.startedAt)
	}
	return e.finishedAt.Sub(e.startedAt) //Sub - берёт от текущешего времнеи и вычитает то, что передадим в аргументе
}

// Метод который будет собирать текущую информацию по работе предприятия, в последствии предполагаю надо дополнить
func (e *Enterprise) Status() EnterpriseStatus {
	return EnterpriseStatus{
		Coal:     e.coal,
		WorkTime: e.WorkTime(time.Now()),
	}
}

// Метод приобретения оборудования
func (e *Enterprise) BuyEquipment(equipment EquipmentType) error {
	eq, ok := e.equipments[equipment]
	if !ok {
		//Ошибку желательно потом будет вынести в отдельный тип, чтобы возвращать, что то адекватное в http
		return errors.New("Отсутствует данный тип оборудования")
	}

	if eq.Purchased {
		//Ошибку желательно потом будет вынести в отдельный тип, чтобы возвращать, что то адекватное в http
		return errors.New("Оборудование уже приобретено")
	}

	if e.coal < eq.Cost {
		return errors.New("Не хватает угля")
	}
	e.AddCoal(-eq.Cost)
	eq.Purchased = true
	return nil
}

// Метод для получения всего оборудования
func (e *Enterprise) Equipments() []Equipment {
	result := make([]Equipment, 0, len(e.equipments))

	for _, equipment := range e.equipments {
		result = append(result, *equipment)
	}

	return result
}

// Метод для получения приобретённого оборудования
func (e *Enterprise) EquipmentsPurchased() []Equipment {
	result := make([]Equipment, 0, len(e.equipments))

	for _, equipment := range e.equipments {
		if equipment.Purchased {
			result = append(result, *equipment)
		}
	}

	return result
}
