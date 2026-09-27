package storage

import "errors"

var (
	ErrStorage                = errors.New("storage error")
	ErrDuplicateEquipmentType = errors.New("duplicate equipment type")
	ErrDuplicateMinerType     = errors.New("duplicate miner type")
)
