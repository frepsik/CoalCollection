package http

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type gameDTO struct {
	state string
}

// Проверяем, корректное ли значение пришло для запуска/завершения работы предприятия
func (g gameDTO) validate() error {
	state := strings.TrimSpace(g.state)
	if state == "" {
		return errors.New("Argument operation game is empty")
	}
	if !(state == "start" || state == "shutdown") {
		return errors.New("Invalid operation game")
	}
	return nil
}

type errorDTO struct {
	Message string
	Time    time.Time
}

func newErrorDTO(message string) errorDTO {
	return errorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

func (e errorDTO) ToString() string {
	bytes, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(bytes)

}
