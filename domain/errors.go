package domain

import "errors"

var (
	// Enterpise
	ErrEquipmentNotFound         = errors.New("equipment not found")
	ErrEnterpriseNotEnoughCoal   = errors.New("not enough coal")
	ErrEquipmentAlreadyPurchased = errors.New("equipment already purchased")

	// Equipment
	ErrInvalidEquipmentTypeName = errors.New("invalid type name equipment")
	ErrInvalidEquipmentName     = errors.New("invalid name equipment")
	ErrInvalidEquipmentCost     = errors.New("invalid cost equipment")

	// Miner
	ErrInvalidMinerTypeName   = errors.New("invalid type name miner")
	ErrInvalidMinerName       = errors.New("invalid name miner")
	ErrInvalidMinerCost       = errors.New("invalid cost miner")
	ErrInvalidMinerEnergy     = errors.New("invalid energy miner")
	ErrInvalidMinerExtraction = errors.New("invalid extraction miner")
	ErrInvalidMinerInterval   = errors.New("invalid interval miner")
	ErrInvalidMinerGrowth     = errors.New("invalid growth miner")
)
