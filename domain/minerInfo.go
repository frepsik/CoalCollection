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
