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
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Завершение работы предприятия
func (h *Handler) finishGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Shutdown(); err != nil {
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Получение статуса, текущей информации по предприятию в моменте
func (h *Handler) statusEnterprise(w http.ResponseWriter, r *http.Request) {
	enterpriseStatus, err := h.gameService.StatusEnterprise()
	if err != nil {
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
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
	query := r.URL.Query()
	for key := range query {
		if key != "state" {
			statusCodeError := statusCodeFromError(ErrInvalidQueryOptions)
			errorResponse := writeError(ErrInvalidQueryOptions.Error())
			writeJSON(w, statusCodeError, errorResponse)
			return
		}
	}
	// Если ключа state вообще нет
	if !query.Has("state") {
		statusCodeError := statusCodeFromError(ErrInvalidQueryOptions)
		errorResponse := writeError(ErrInvalidQueryOptions.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}

	stateEquipments := query.Get("state")

	switch stateEquipments {
	case "purchased":
		equipments, err := h.gameService.EquipmentsPurchased()
		if err != nil {
			statusCodeError := statusCodeFromError(err)
			errorResponse := writeError(err.Error())
			writeJSON(w, statusCodeError, errorResponse)
			return
		}
		equipmentsResponses := newEquipmentResponses(equipments)
		writeJSON(w, http.StatusOK, equipmentsResponses)
	case "not_purchased":
		equipments, err := h.gameService.EquipmentsNotPurchased()
		if err != nil {
			statusCodeError := statusCodeFromError(err)
			errorResponse := writeError(err.Error())
			writeJSON(w, statusCodeError, errorResponse)
			return
		}
		equipmentResponses := newEquipmentResponses(equipments)
		writeJSON(w, http.StatusOK, equipmentResponses)
	default:
		statusCodeError := statusCodeFromError(ErrInvalidEquipmentState)
		errorResponse := writeError(ErrInvalidEquipmentState.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
}

// Приобретение оборудования POST /equipments/{type}/purchase пример
func (h *Handler) buyEquipment(w http.ResponseWriter, r *http.Request) {
	equipmentType := r.PathValue("type")
	if err := h.gameService.BuyEquipment(equipmentType); err != nil {
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Нанять шахтёра (должно приходить json тело) POST /miners/strongMiner пример
func (h *Handler) hireMiner(w http.ResponseWriter, r *http.Request) {
	var request hireMinerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		statusCodeError := statusCodeFromError(ErrInvalidRequestBody)
		errorResponse := writeError(ErrInvalidRequestBody.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
	if !request.validate() {
		statusCodeError := statusCodeFromError(ErrInvalidRequestBody)
		errorResponse := writeError(ErrInvalidRequestBody.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}

	err := h.gameService.HireMiner(request.MinerTypeName)
	if err != nil {
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
// Пример такого запроса GET /miners?type=strongMiner&state=available
// Но пока необходимости нет, оставлю как пример так, но в случае чего нужно будет переписать методы в Service, с фильтрацией
func (h *Handler) miners(w http.ResponseWriter, r *http.Request) {
	minerFilter := service.MinerFilter{}
	// Получаем map с ключами, и значениям, где ключ это первое значение после вопроса, например type, а значенеи, это уже значение после =, наприммер available
	query := r.URL.Query()
	//Проверяем текущие ключи, если они не совпадают (если ключей вообще не передали, то query будет пустой как map, там не будет ничего, цикл не пройдёт)
	for key := range query {
		if key != "type" && key != "state" {
			statusCodeError := statusCodeFromError(ErrInvalidQueryOptions)
			errorResponse := writeError(ErrInvalidQueryOptions.Error())
			writeJSON(w, statusCodeError, errorResponse)
			return
		}
	}

	//Проверяем, передали ли хоть один из этих, ключей, а если передали, конкретно нужный вернёт true и какой из них
	if query.Has("state") {
		state := query.Get("state")
		minerFilter.State = &state
	}
	if query.Has("type") {
		minerType := query.Get("type")
		minerFilter.Type = &minerType
	}

	miners, err := h.gameService.Miners(minerFilter)
	if err != nil {
		statusCodeError := statusCodeFromError(err)
		errorResponse := writeError(err.Error())
		writeJSON(w, statusCodeError, errorResponse)
		return
	}

	minerResponses := newMinerResponses(miners)
	writeJSON(w, http.StatusOK, minerResponses)
}
