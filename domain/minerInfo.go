package domain

import (
	"time"

	"github.com/google/uuid"
)

type MinerInfo struct {
	Id         uuid.UUID
	TypeName   MinerTypeName
	Name       string
	Cost       Coal
	LeftEnergy int
	Extraction Coal
	Interval   time.Duration
	Growth     Coal
}

func newMinerInfo(
	id uuid.UUID,
	typeName MinerTypeName,
	name string,
	cost Coal,
	leftEnergy int,
	extraction Coal,
	interval time.Duration,
	growth Coal,
) MinerInfo {
	return MinerInfo{
		Id:         id,
		TypeName:   typeName,
		Name:       name,
		Cost:       cost,
		LeftEnergy: leftEnergy,
		Extraction: extraction,
		Interval:   interval,
		Growth:     growth,
	}
}

// Структура под вывод данных в Json
type MinerTypeInfo struct {
	TypeName   MinerTypeName
	Name       string
	Cost       Coal
	Energy     int
	Extraction Coal
	Interval   time.Duration
	Growth     Coal
}

func NewMinerTypeInfo(minerType MinerType) MinerTypeInfo {
	return MinerTypeInfo{
		TypeName:   minerType.typeName,
		Name:       minerType.name,
		Cost:       minerType.cost,
		Energy:     minerType.energy,
		Extraction: minerType.extraction,
		Interval:   minerType.interval,
		Growth:     minerType.growth,
	}
}
