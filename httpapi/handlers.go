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
	w.WriteHeader(http.StatusNoContent)
}

// Завершение работы предприятия
func (h *Handler) shutdownGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Shutdown(); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		return
	}
}

// Приобретение оборудования POST /equipments/{type}/purchase пример
func (h *Handler) buyEquipment(w http.ResponseWriter, r *http.Request) {
	equipmentType := r.PathValue("type")
	if err := h.gameService.BuyEquipment(equipmentType); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Нанять шахтёра (должно приходить json тело) POST /miners/strongMiner пример
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
	// Здесь без проверок потому что сценария в котором случится так, что данный список будет пуст просто не будет, игра не запустится если при сборке данных не будет
	minersTypeInfo := h.gameService.MinersType()
	minerTypeResponses := make([]minerTypeResponse, 0, len(minersTypeInfo))
	for _, miner := range minersTypeInfo {
		minerTypeResponses = append(minerTypeResponses, newMinerTypeResponse(miner))
	}
	writeJSON(w, http.StatusOK, minerTypeResponses)
}

// Получение шахтёров в зависимости от query параметра
// Вообще если вырастить функционал и потребуется различная фильтрация будет лучше использовать следующую концепцию с фильтром
//
//	type MinerFilter struct {
//		Type  *domain.MinerTypeName
//		State *domain.MinerState
//	}
//
// Пример запроса GET /miners?type=strongMiner&state=available
func (h *Handler) miners(w http.ResponseWriter, r *http.Request) {
	stateMiner := r.URL.Query().Get("state")
	typeMiner := r.URL.Query().Get("type")
	switch {
	case stateMiner == "exhausted":
		miners, err := h.gameService.ExhaustedMiners()
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		minerResponses := newMinerResponses(miners)
		writeJSON(w, http.StatusOK, minerResponses)
	case stateMiner == "available":
		miners, err := h.gameService.AvailableMiners()
		if err != nil {
			errorResponse := writeError(err.Error())
			writeJSON(w, http.StatusBadRequest, errorResponse)
			return
		}
		minerResponses := newMinerResponses(miners)
		writeJSON(w, http.StatusOK, minerResponses)
	case typeMiner != "":
		h.minersByType(w, typeMiner)
	default:
		errorResponse := writeError("state query parameter is required or invalid miner state")
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
}

// Получение шахтёров по их классу
func (h *Handler) minersByType(w http.ResponseWriter, minerType string) {
	miners, err := h.gameService.MinersByType(minerType)
	if err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	minerResponses := newMinerResponses(miners)
	writeJSON(w, http.StatusOK, minerResponses)
}
