package service

type gameState int

// Типа enum, у которого константы имеют тип gameState, который число. Здесь iota говорит, что начинаем с нуля, и с каждой следующей строчкой увеличивается на 1
// некоторый счётчик
const (
	gameCreated  gameState = iota // 0
	gameRunning                   // 1
	gameFinished                  // 2
)
