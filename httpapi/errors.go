package httpapi

import "errors"

var (
	ErrInvalidQueryOptions   = errors.New("invalid request query options") //400
	ErrInvalidEquipmentState = errors.New("invalid equipment state")       //400
	ErrInvalidRequestBody    = errors.New("invalid request json body")     //400
)
