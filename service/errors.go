package service

import "errors"

var (
	ErrMinerTypeNotFound     = errors.New("miner type not found")
	ErrEquipmentTypeNotFound = errors.New("equipment type not found")
)
