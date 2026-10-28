package httpapi

import "CoalCollection/domain"

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
