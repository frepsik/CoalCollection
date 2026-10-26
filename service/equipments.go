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
func (g *Game) EquipmentsPurchased() ([]domain.EquipmentInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return nil, ErrGameIsNotRunning
	}
	return g.enterprise.EquipmentsPurchased(), nil
}

// Метод для получения не купленного оборудования
func (g *Game) EquipmentsNotPurchased() ([]domain.EquipmentInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return nil, ErrGameIsNotRunning
	}

	return g.enterprise.EquipmentsNotPurchased(), nil
}

// Метод для покупки оборудования в service
func (g *Game) BuyEquipment(equipmentType string) error {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return ErrGameIsNotRunning
	}

	equipmentTypeName := domain.EquipmentTypeName(equipmentType)
	_, exists := g.equipmentTypes[equipmentTypeName]
	if !exists {
		return ErrEquipmentTypeNotFound
	}

	return g.enterprise.BuyEquipment(equipmentTypeName)
}
