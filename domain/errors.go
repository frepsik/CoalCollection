package domain

import "errors"

var (
	ErrEquipmentNotFound         = errors.New("equipment not found")
	ErrNotEnoughCoal             = errors.New("not enough coal")
	ErrEquipmentAlreadyPurchased = errors.New("equipment already purchased")
)
