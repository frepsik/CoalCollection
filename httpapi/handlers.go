package httpapi

import (
	"CoalCollection/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	gameService *service.Game
}

func NewHandler(gameService *service.Game) *Handler {
	return &Handler{
		gameService: gameService,
	}
}

// Запуск работы предприятия
func (h *Handler) startGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Start(); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Завершение работы предприятия
func (h *Handler) shutdownGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Shutdown(); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Получение статуса, текущей информации по предприятию в моменте
func (h *Handler) statusEnterprise(w http.ResponseWriter, r *http.Request) {
	enterpriseStatus, err := h.gameService.StatusEnterprise()
	if err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	enterpriseStatusResponse := newEnterPriseStatusResponse(enterpriseStatus)

	writeJSON(w, http.StatusOK, enterpriseStatusResponse)
}

// Получение оборудования, которое можно приобрести (вся информация о нём)
func (h *Handler) equipmentsType(w http.ResponseWriter, r *http.Request) {
	// Здесь без проверок потому что сценария в котором случится так, что данный список будет пуст просто не будет, игра не запустится если при сборке данных не будет
	equipmentsType := h.gameService.Equipments()
	equipmentTypeResponses := make([]equipmentTypeResponse, 0, len(equipmentsType))
	for _, equipmentType := range equipmentsType {
		equipmentTypeResponses = append(equipmentTypeResponses, newEquipmentTypeResponse(equipmentType))
	}
	writeJSON(w, http.StatusOK, equipmentTypeResponses)
}

// Получение оборудования, которое приобретено/не приобретено
func (h *Handler) equipments(w http.ResponseWriter, r *http.Request) {
	// Если челик при отправке запроса, написал вместо state, status /equipments?status=purchased, он залетит в default, т.к не правильно указан
	// параметр, и stateEquipments = "", если просто не правильно указан сам аргумент, чёто типа status /equipments?status=abracadabra
	// тогда там будет abracadabra, что у нас тоже не ожидается в качестве состояния
	stateEquipments := r.URL.Query().Get("state")

	switch stateEquipments {
	case "purchased":
		equipments, err := h.gameService.EquipmentsPurchased()
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		equipmentsResponses := newEquipmentResponses(equipments)
		writeJSON(w, http.StatusOK, equipmentsResponses)
	case "not_purchased":
		equipments, err := h.gameService.EquipmentsNotPurchased()
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		equipmentResponses := newEquipmentResponses(equipments)
		writeJSON(w, http.StatusOK, equipmentResponses)
	default:
		errorResponse := writeError("state query parameter is required or invalid equipment state")
		writeJSON(w, http.StatusBadRequest, errorResponse)
	}
}

// Приобретение оборудования
func (h *Handler) buyEquipment(w http.ResponseWriter, r *http.Request) {

}

// Нанять шахтёра
func (h *Handler) hireMiner(w http.ResponseWriter, r *http.Request) {
	var request hireMinerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorResponse := writeError("invalid request body")
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	if !request.validate() {
		errorResponse := writeError("invalid request body")
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}

	switch request.MinerTypeName {
	case "smallMiner":
		minerTypeInfo, err := h.gameService.HireMiner(request.MinerTypeName)
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		writeJSON(w, http.StatusCreated, newMinerTypeResponse(minerTypeInfo))
	case "normalMiner":
		minerTypeInfo, err := h.gameService.HireMiner(request.MinerTypeName)
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		writeJSON(w, http.StatusCreated, newMinerTypeResponse(minerTypeInfo))
	case "strongMiner":
		minerTypeInfo, err := h.gameService.HireMiner(request.MinerTypeName)
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		writeJSON(w, http.StatusCreated, newMinerTypeResponse(minerTypeInfo))
	}
}

// Получение типов шахтёров, которых можно нанять
func (h *Handler) minersType(w http.ResponseWriter, r *http.Request) {

}

// Получение нанятых шахтёров, которые сейчас работают
func (h *Handler) availableMiners(w http.ResponseWriter, r *http.Request) {

}

// Получение нанятых шахтёров, которые сейчас не работают
func (h *Handler) exhaustedMiners(w http.ResponseWriter, r *http.Request) {

}

// Получение шахтёров по их классу
func (h *Handler) minersByType(w http.ResponseWriter, r *http.Request) {

}
