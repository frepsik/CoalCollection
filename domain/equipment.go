package domain

import "strings"

type EquipmentTypeName string

const (
	Pickaxe     EquipmentTypeName = "pickaxe"
	Ventilation EquipmentTypeName = "ventilation"
	Trolleys    EquipmentTypeName = "trolleys"
)

type EquipmentType struct {
	equipmentTypeName EquipmentTypeName
	name              string
	cost              Coal
}

func NewEquipmentType(equipmentTypeName EquipmentTypeName, equipmentName string, costEquipment Coal) (EquipmentType, error) {

	equipment := EquipmentType{
		equipmentTypeName: equipmentTypeName,
		name:              equipmentName,
		cost:              costEquipment,
	}

	if err := equipment.validate(); err != nil {
		return EquipmentType{}, err
	}

	return equipment, nil
}

// Небольшой валидатор аргументов оборудования
func (e *EquipmentType) validate() error {
	if !(e.equipmentTypeName == Pickaxe || e.equipmentTypeName == Ventilation || e.equipmentTypeName == Trolleys) {
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

// Структура для состояния в рамках игры
type Equipment struct {
	equipmentType EquipmentType
	purchased     bool
}

// Конструктор осуществляющий создание экземпляра оборудования по указателю, в связи с необходимость отслеживания его сотсояния
func NewEquipment(equipmentType EquipmentType) *Equipment {
	return &Equipment{
		equipmentType: equipmentType,
	}
}

// Структура для вывода состояния купленного/не купленного оборудования на момент игры
type EquipmentInfo struct {
	EquipmentTypeName EquipmentTypeName
	Name              string
	Cost              Coal
	Purchased         bool
}

func newEquipmentInfo(equipment Equipment) EquipmentInfo {
	return EquipmentInfo{
		EquipmentTypeName: equipment.equipmentType.equipmentTypeName,
		Name:              equipment.equipmentType.name,
		Cost:              equipment.equipmentType.cost,
		Purchased:         equipment.purchased,
	}
}
