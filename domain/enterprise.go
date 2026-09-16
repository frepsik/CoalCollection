package domain

import (
	"sync"
	"time"
)

type Enterprise struct {
	coal       int
	mtx        sync.Mutex
	startedAt  time.Time
	finishedAt time.Time
}

// Конструктор для структуры Enterprise
func NewEnterprise() *Enterprise {
	return &Enterprise{}
}

// Метод для добавления угля к общему балансу
func (e *Enterprise) AddCoal(amountCoal int) {
	e.mtx.Lock()
	e.coal += amountCoal
	e.mtx.Unlock()
}

// Метод для получения количества текущего угля
func (e *Enterprise) Coal() int {
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
