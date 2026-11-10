package domain

import "errors"

var (
	// Enterprise
	ErrEquipmentNotFound         = errors.New("equipment not found")         //404
	ErrEnterpriseNotEnoughCoal   = errors.New("not enough coal")             // 409
	ErrEquipmentAlreadyPurchased = errors.New("equipment already purchased") //409

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
