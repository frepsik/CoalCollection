package httpapi

import (
	"CoalCollection/domain"
	"CoalCollection/service"
	"errors"
	"net/http"
)

var (
	ErrInvalidQueryOptions   = errors.New("invalid request query options") //400
	ErrInvalidEquipmentState = errors.New("invalid equipment state")       //400
	ErrInvalidRequestBody    = errors.New("invalid request json body")     //400
)

// helper для определения ошибки на выход
func statusCodeFromError(err error) int {
	switch {
	// bad-request 400
	case errors.Is(err, service.ErrMinerTypeNotFound),
		errors.Is(err, service.ErrEquipmentTypeNotFound),
		errors.Is(err, service.ErrUnsupportedFilterCombination),
		errors.Is(err, service.ErrInvalidMinerState),
		errors.Is(err, service.ErrUnsupportedCombination),
		errors.Is(err, ErrInvalidQueryOptions),
		errors.Is(err, ErrInvalidEquipmentState),
		errors.Is(err, ErrInvalidRequestBody):

		return http.StatusBadRequest

	// conflict 409
	case errors.Is(err, service.ErrGameCannotStart),
		errors.Is(err, service.ErrGameCannotFinish),
		errors.Is(err, service.ErrGameIsNotRunning),
		errors.Is(err, service.ErrGameIsNotStart),
		errors.Is(err, domain.ErrEnterpriseNotEnoughCoal),
		errors.Is(err, domain.ErrEquipmentAlreadyPurchased):

		return http.StatusConflict
	// not found 404
	case errors.Is(err, domain.ErrEquipmentNotFound):
		return http.StatusNotFound
	// internal server 500
	default:
		return http.StatusInternalServerError
	}
}
