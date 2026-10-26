package service

// Структура и методы отвечающие за текущее состояние игры (работает ли предприятие)
type gameState struct {
	gameStarted bool
}

// Возвращает текущее состояние игры
func (gs *gameState) currentState() bool {
	return gs.gameStarted
}

func (gs *gameState) start() {
	gs.gameStarted = true
}

func (gs *gameState) finish() {
	gs.gameStarted = false
}
