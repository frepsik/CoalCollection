package domain

import "errors"

var ErrSearchEquipmentByType = errors.New("equipment not found")
var ErrNotEnoughCoal = errors.New("not enough coal")
var ErrEquipmentAlreadyPurchased = errors.New("equipment already purchased")
