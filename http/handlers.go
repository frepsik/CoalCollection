package http

import (
	"CoalCollection/service"
	"net/http"
)

type HttpHandlers struct {
	gameService *service.Game
}

func NewHttpHandlers(gameService *service.Game) *HttpHandlers {
	return &HttpHandlers{
		gameService: gameService,
	}
}

// Запуск работы предприятия
func (h *HttpHandlers) handleStartGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Start(); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Завершение работы предприятия
func (h *HttpHandlers) handleShutdownGame(w http.ResponseWriter, r *http.Request) {
	if err := h.gameService.Shutdown(); err != nil {
		errorResponse := writeError(err.Error())
		writeJSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Получение статуса, текущей информации по предприятию в моменте
func (h *HttpHandlers) handleStatusEnterprise(w http.ResponseWriter, r *http.Request) {

}

// Получение оборудования, которое можно приобрести (вся информация о нём)
func (h *HttpHandlers) handleEquipments(w http.ResponseWriter, r *http.Request) {

}

// Получение оборудования, которое приобретено
func (h *HttpHandlers) handleEquipmentPurchased(w http.ResponseWriter, r *http.Request) {

}

// Получение оборудования, которое не приобретено
func (h *HttpHandlers) handleEquipmentNotPurchased(w http.ResponseWriter, r *http.Request) {

}

// Приобретение оборудования
func (h *HttpHandlers) handleBuyEquipment(w http.ResponseWriter, r *http.Request) {

}

// Нанять шахтёра
func (h *HttpHandlers) handleHireMiner(w http.ResponseWriter, r *http.Request) {

}

// Получение типов шахтёров, которых можно нанять
func (h *HttpHandlers) handleMinersType(w http.ResponseWriter, r *http.Request) {

}

// Получение нанятых шахтёров, которые сейчас работают
func (h *HttpHandlers) handleAvailableMiners(w http.ResponseWriter, r *http.Request) {

}

// Получение нанятых шахтёров, которые сейчас не работают
func (h *HttpHandlers) handleExhaustedMiners(w http.ResponseWriter, r *http.Request) {

}

// Получение шахтёров по их классу
func (h *HttpHandlers) handleMinersByType(w http.ResponseWriter, r *http.Request) {

}
