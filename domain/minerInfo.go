package domain

import (
	"time"

	"github.com/google/uuid"
)

type MinerInfo struct {
	id         uuid.UUID
	typeName   MinerTypeName
	name       string
	cost       int
	energy     int
	extraction int
	interval   time.Duration
	growth     int
}

func newMinerInfo(
	id uuid.UUID,
	typeName MinerTypeName,
	name string,
	cost int,
	energy int,
	extraction int,
	interval time.Duration,
	growth int,
) MinerInfo {
	return MinerInfo{
		id:         id,
		typeName:   typeName,
		name:       name,
		cost:       cost,
		energy:     energy,
		extraction: extraction,
		interval:   interval,
		growth:     growth,
	}
}
