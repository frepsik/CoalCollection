package domain

import "strings"

type EquipmentType string

const (
	Pickaxe     EquipmentType = "pickaxe"
	Ventilation EquipmentType = "ventilation"
	Trolleys    EquipmentType = "trolleys"
)

type Equipment struct {
	equipmentType EquipmentType
	name          string
	cost          Coal
	purchased     bool
}

func NewEquipment(equipmentType EquipmentType, equipmentName string, costEquipment Coal) (*Equipment, error) {

	equipment := &Equipment{
		equipmentType: equipmentType,
		name:          equipmentName,
		cost:          costEquipment,
	}

	if err := equipment.validate(); err != nil {
		return nil, err
	}

	return equipment, nil
}

// Небольшой валидатор аргументов оборудования
func (e *Equipment) validate() error {
	if !(e.equipmentType == Pickaxe || e.equipmentType == Ventilation || e.equipmentType == Trolleys) {
		return ErrInvalidEquipmentTypeName
	}
	if strings.TrimSpace(e.name) == "" {
		return ErrInvalidEquipmentName
	}
	if e.cost < 0 {
		return ErrInvalidEquipmentCost
	}
	return nil
}
