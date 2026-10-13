package domain

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

// Структура под вывод данных в Json
type EquipmentTypesInfo struct {
	EquipmentTypeName EquipmentTypeName
	Name              string
	Cost              Coal
}

func NewEquipmentTypesInfo(equipmentType EquipmentType) EquipmentTypesInfo {
	return EquipmentTypesInfo{
		EquipmentTypeName: equipmentType.equipmentTypeName,
		Name:              equipmentType.name,
		Cost:              equipmentType.cost,
	}
}
