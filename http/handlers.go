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

// Запуск/завершение работы предприятия
func (h *HttpHandlers) handleGame(w http.ResponseWriter, r *http.Request) {

	// Получаем значение переданное в качестве паттерна
	gameDTO := gameDTO{
		state: r.PathValue("state"),
	}

	// Проверяем, адекватное ли пришло значение
	if err := gameDTO.validate(); err != nil {
		errorDTO := newErrorDTO(err.Error())

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)

		return
	}

	// Запускаем либо завершаем работу предприятия
	if gameDTO.state == "start" {
		h.gameService.Start()
	} else {
		h.gameService.Shutdown()
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
