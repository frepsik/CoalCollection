package service

import (
	"CoalCollection/domain"
	"time"
)

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

// Метод позволяющий узнать, текущий статус, по тому, сколько идёт игра, и сколько сейчас угля, сюда дальше ещё надо интегрировать шахтёров и оборудование, но пока временно так
func (g *Game) StatusEnterprise() domain.EnterpriseStatus {
	return g.enterprise.Status()
}
