package httpapi

import (
	"CoalCollection/domain"
)

// Метод для сборки списка оборудования под вывод в api
func newEquipmentResponses(equipmentsInfo []domain.EquipmentInfo) []equipmentResponse {
	result := make([]equipmentResponse, 0, len(equipmentsInfo))
	for _, equipmentInfo := range equipmentsInfo {
		result = append(result, newEquipmentResponse(equipmentInfo))
	}
	return result
}

// Метод для сборки списка шахтёров под вывод в api
func newMinerResponses(minersInfo []domain.MinerInfo) []minerResponse {
	result := make([]minerResponse, 0, len(minersInfo))
	for _, minerInfo := range minersInfo {
		result = append(result, newMinerResponse(minerInfo))
	}
	return result
}

// Метод для создания структуры под вывод информации по оборудованию
func newEquipmentResponse(equipmentInfo domain.EquipmentInfo) equipmentResponse {
	return equipmentResponse{
		EquipmentTypeName: string(equipmentInfo.EquipmentTypeName),
		Name:              equipmentInfo.Name,
		Cost:              int(equipmentInfo.Cost),
		Purchased:         equipmentInfo.Purchased,
	}
}

// Метод для создания структуры под вывод информации по типам оборудования
func newEquipmentTypeResponse(equipmentTypeInfo domain.EquipmentTypeInfo) equipmentTypeResponse {
	return equipmentTypeResponse{
		EquipmentTypeName: string(equipmentTypeInfo.EquipmentTypeName),
		Name:              equipmentTypeInfo.Name,
		Cost:              int(equipmentTypeInfo.Cost),
	}
}

// Метод для создания структуры под вывод информации по шахтёрам
func newMinerResponse(minerInfo domain.MinerInfo) minerResponse {
	return minerResponse{
		Id:              minerInfo.Id,
		TypeName:        string(minerInfo.TypeName),
		Name:            minerInfo.Name,
		Cost:            int(minerInfo.Cost),
		LeftEnergy:      minerInfo.LeftEnergy,
		Extraction:      int(minerInfo.Extraction),
		IntervalSeconds: int(minerInfo.Interval.Seconds()),
		Growth:          int(minerInfo.Growth),
	}
}

// Метод для создания структуры под вывод информации по типам шахтёров
func newMinerTypeResponse(minerTypeInfo domain.MinerTypeInfo) minerTypeResponse {
	return minerTypeResponse{
		TypeName:        string(minerTypeInfo.TypeName),
		Name:            minerTypeInfo.Name,
		Cost:            int(minerTypeInfo.Cost),
		Energy:          minerTypeInfo.Energy,
		Extraction:      int(minerTypeInfo.Extraction),
		IntervalSeconds: int(minerTypeInfo.Interval.Seconds()),
		Growth:          int(minerTypeInfo.Growth),
	}
}

// Метод для создания структуры под вывод статистики по предприятию
func newEnterPriseStatusResponse(enterPriseStatus domain.EnterpriseStatus) enterpriseStatusResponse {
	equipmentsResponse := make([]equipmentResponse, 0, len(enterPriseStatus.Equipments))
	minersResponse := make([]minerResponse, 0, len(enterPriseStatus.Miners))

	for _, equipment := range enterPriseStatus.Equipments {
		equipmentsResponse = append(equipmentsResponse, newEquipmentResponse(equipment))
	}

	for _, miner := range enterPriseStatus.Miners {
		minersResponse = append(minersResponse, newMinerResponse(miner))
	}

	return enterpriseStatusResponse{
		Coal:            int(enterPriseStatus.Coal),
		WorkTimeSeconds: enterPriseStatus.WorkTime.Seconds(),
		Equipments:      equipmentsResponse,
		Miners:          minersResponse,
	}
}
