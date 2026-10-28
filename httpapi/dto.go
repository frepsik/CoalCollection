package httpapi

import (
	"CoalCollection/domain"
	"strings"

	"github.com/google/uuid"
)

// Структура под вывод информации об ошибке
type errorResponse struct {
	Message string `json:"message"`
}

func writeError(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}

// Структура под вывод информации по оборудованию
type equipmentResponse struct {
	EquipmentTypeName string `json:"equipmentTypeName"`
	Name              string `json:"name"`
	Cost              int    `json:"cost"`
	Purchased         string `json:"purchased"`
}

func newEquipmentResponse(equipmentInfo domain.EquipmentInfo) equipmentResponse {
	var purchasedEquipment string
	if equipmentInfo.Purchased {
		purchasedEquipment = "Приобретено"
	}
	purchasedEquipment = "Не приобретено"
	return equipmentResponse{
		EquipmentTypeName: string(equipmentInfo.EquipmentTypeName),
		Name:              equipmentInfo.Name,
		Cost:              int(equipmentInfo.Cost),
		Purchased:         purchasedEquipment,
	}
}

// Структура под вывод информации по типам оборудования
type equipmentTypeResponse struct {
	EquipmentTypeName string `json:"equipmentTypeName"`
	Name              string `json:"name"`
	Cost              int    `json:"cost"`
}

func newEquipmentTypeResponse(equipmentTypeInfo domain.EquipmentTypeInfo) equipmentTypeResponse {
	return equipmentTypeResponse{
		EquipmentTypeName: string(equipmentTypeInfo.EquipmentTypeName),
		Name:              equipmentTypeInfo.Name,
		Cost:              int(equipmentTypeInfo.Cost),
	}
}

// Структура под вывод информации по шахтёрам
type minerResponse struct {
	Id              uuid.UUID `json:"id"`
	TypeName        string    `json:"typeName"`
	Name            string    `json:"name"`
	Cost            int       `json:"cost"`
	LeftEnergy      int       `json:"lefEnergy"`
	Extraction      int       `json:"extraction"`
	IntervalSeconds int       `json:"intervalSeconds"`
	Growth          int       `json:"growth"`
}

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

// Структура под вывод информации по типам шахтёров
type minerTypeResponse struct {
	TypeName        string `json:"typeName"`
	Name            string `json:"name"`
	Cost            int    `json:"cost"`
	Energy          int    `json:"energy"`
	Extraction      int    `json:"extraction"`
	IntervalSeconds int    `json:"intervalSeconds"`
	Growth          int    `json:"growth"`
}

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

// Структура под вывод статистики по предприятию
type enterpriseStatusResponse struct {
	Coal            int                 `json:"coal"`
	WorkTimeSeconds float64             `json:"workTimeSeconds"`
	Equipments      []equipmentResponse `json:"equipments"`
	Miners          []minerResponse     `json:"miners"`
}

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

// DTO - для запроса на создание шахтёра, на текущий момент такая коротенькая версия структуры, в связи с тем
// что имеет потенциал развития
type hireMinerRequest struct {
	MinerTypeName string `json:"minerTypeName"`
}

func (h *hireMinerRequest) validate() bool {
	h.MinerTypeName = strings.TrimSpace(h.MinerTypeName)
	if h.MinerTypeName == "" {
		return false
	}
	if !(h.MinerTypeName == "smallMiner" || h.MinerTypeName == "strongMiner" || h.MinerTypeName == "normalMiner") {
		return false
	}
	return true
}
