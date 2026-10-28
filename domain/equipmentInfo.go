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
type EquipmentTypeInfo struct {
	EquipmentTypeName EquipmentTypeName
	Name              string
	Cost              Coal
}

func NewEquipmentTypeInfo(equipmentType EquipmentType) EquipmentTypeInfo {
	return EquipmentTypeInfo{
		EquipmentTypeName: equipmentType.equipmentTypeName,
		Name:              equipmentType.name,
		Cost:              equipmentType.cost,
	}
}
