package service

import "errors"

var (
	ErrMinerTypeNotFound     = errors.New("miner type not found")
	ErrEquipmentTypeNotFound = errors.New("equipment type not found")

	ErrGameCannotStart  = errors.New("game cannot start")
	ErrGameCannotFinish = errors.New("game cannot finish")
	ErrGameIsNotRunning = errors.New("game is not running")
	ErrGameIsNotStart   = errors.New("current game created and not started")

	ErrUnsupportedFilterCombination = errors.New("unsupported filter combination")
	ErrInvalidMinerState            = errors.New("invalid miner state")
)
