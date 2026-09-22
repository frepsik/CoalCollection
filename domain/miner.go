package domain

import (
	"time"

	"github.com/google/uuid"
)

type MinerTypeName string

var (
	SmallMiner  MinerTypeName = "smallMiner"
	NormalMiner MinerTypeName = "normalMiner"
	StrongMiner MinerTypeName = "strongMiner"
)

type MinerType struct {
	typeName   MinerTypeName
	cost       Coal
	energy     int
	extraction Coal
	interval   time.Duration
	growth     Coal
}

func NewMinerType(
	typeName MinerTypeName,
	cost Coal,
	energy int,
	extraction Coal,
	interval time.Duration,
	growth Coal,
) (MinerType, error) {
	if typeName != SmallMiner || typeName != NormalMiner || typeName != StrongMiner {
		return MinerType{}, ErrInvalidMinerTypeName
	} else if cost <= 0 {
		return MinerType{}, ErrInvalidMinerCost
	} else if energy <= 0 {
		return MinerType{}, ErrInvalidMinerEnergy
	} else if extraction <= 0 {
		return MinerType{}, ErrInvalidMinerExtraction
	} else if interval <= 0 {
		return MinerType{}, ErrInvalidMinerInterval
	} else if growth < 0 {
		return MinerType{}, ErrInvalidMinerGrowth
	}

	return MinerType{
		typeName:   typeName,
		cost:       cost,
		energy:     energy,
		extraction: extraction,
		interval:   interval,
		growth:     Coal(growth),
	}, nil
}

type Miner struct {
	id         uuid.UUID
	minerType  MinerType
	actionDone int
}

func NewMiner(minerType MinerType) *Miner {
	return &Miner{
		id:        uuid.New(),
		minerType: minerType,
	}
}

// Метод на добычу угля
func (m *Miner) Mine() (Coal, bool) {

	if m.actionDone >= m.minerType.energy {
		return 0, true
	}

	coal := int(m.minerType.extraction) + m.actionDone*int(m.minerType.growth)

	m.actionDone++

	finishWorkMiner := m.actionDone >= m.minerType.energy

	return Coal(coal), finishWorkMiner
}

// Метод на получения временного интервала добычи угля
func (m *Miner) MiningInterval() time.Duration {
	return m.minerType.interval
}
