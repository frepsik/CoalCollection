package service

import "CoalCollection/domain"

// Метод для получения оборудования, которое можно приобрести
func (g *Game) Equipments() []domain.EquipmentType {
	result := make([]domain.EquipmentType, 0, len(g.equipmentTypes))

	for _, equipmentType := range g.equipmentTypes {
		result = append(result, equipmentType)
	}

	return result
}

// Метод для получения купленного на текущий момент оборудования
func (g *Game) EquipmentsPurchased() []domain.EquipmentInfo {
	return g.enterprise.EquipmentsPurchased()
}

// Метод для получения не купленного оборудования
func (g *Game) EquipmentsNotPurchased() []domain.EquipmentInfo {
	return g.enterprise.EquipmentsNotPurchased()
}

// Метод для покупки оборудования в service
func (g *Game) BuyEquipment(equipmentType string) error {
	equipmentTypeName := domain.EquipmentTypeName(equipmentType)
	_, exists := g.equipmentTypes[equipmentTypeName]
	if !exists {
		return ErrEquipmentTypeNotFound
	}

	return g.enterprise.BuyEquipment(equipmentTypeName)
}
