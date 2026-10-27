package httpapi

import (
	"CoalCollection/service"
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

}

// Получение оборудования, которое можно приобрести (вся информация о нём)
func (h *Handler) equipments(w http.ResponseWriter, r *http.Request) {

}

// Получение оборудования, которое приобретено
func (h *Handler) equipmentPurchased(w http.ResponseWriter, r *http.Request) {

}

// Получение оборудования, которое не приобретено
func (h *Handler) equipmentNotPurchased(w http.ResponseWriter, r *http.Request) {

}

// Приобретение оборудования
func (h *Handler) buyEquipment(w http.ResponseWriter, r *http.Request) {

}

// Нанять шахтёра
func (h *Handler) hireMiner(w http.ResponseWriter, r *http.Request) {

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
