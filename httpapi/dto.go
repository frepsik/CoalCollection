package httpapi

import (
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
	Purchased         bool   `json:"purchased"`
}

// Структура под вывод информации по типам оборудования
type equipmentTypeResponse struct {
	EquipmentTypeName string `json:"equipmentTypeName"`
	Name              string `json:"name"`
	Cost              int    `json:"cost"`
}

// Структура под вывод информации по шахтёрам
type minerResponse struct {
	Id              uuid.UUID `json:"id"`
	TypeName        string    `json:"typeName"`
	Name            string    `json:"name"`
	Cost            int       `json:"cost"`
	LeftEnergy      int       `json:"leftEnergy"`
	Extraction      int       `json:"extraction"`
	IntervalSeconds int       `json:"intervalSeconds"`
	Growth          int       `json:"growth"`
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

// Структура под вывод статистики по предприятию
type enterpriseStatusResponse struct {
	Coal            int                 `json:"coal"`
	WorkTimeSeconds float64             `json:"workTimeSeconds"`
	Equipments      []equipmentResponse `json:"equipments"`
	Miners          []minerResponse     `json:"miners"`
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
