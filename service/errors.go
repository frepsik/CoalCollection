package service

import "errors"

var (
	ErrMinerTypeNotFound     = errors.New("miner type not found")
	ErrEquipmentTypeNotFound = errors.New("equipment type not found")

	ErrGameCannotStart  = errors.New("game cannot start")
	ErrGameCannotFinish = errors.New("game cannot finish")
	ErrGameIsFinished   = errors.New("game is already finished")
	ErrGameIsNotRunning = errors.New("game is not running")
)
