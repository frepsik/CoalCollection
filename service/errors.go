package service

import "errors"

var (
	ErrMinerTypeNotFound     = errors.New("miner type not found")     //400
	ErrEquipmentTypeNotFound = errors.New("equipment type not found") //400

	ErrGameCannotStart  = errors.New("game cannot start")                    //409
	ErrGameCannotFinish = errors.New("game cannot finish")                   //409
	ErrGameIsNotRunning = errors.New("game is not running")                  //409
	ErrGameIsNotStart   = errors.New("current game created and not started") //409

	ErrUnsupportedFilterCombination = errors.New("unsupported filter combination") //400
	ErrInvalidMinerState            = errors.New("invalid miner state")            //400
	ErrUnsupportedCombination       = errors.New("unsupported combination")        //400
)
