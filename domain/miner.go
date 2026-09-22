package domain

import (
	"time"

	"github.com/google/uuid"
)

type MinerType struct {
	typeName   string
	cost       Coal
	energy     int
	extraction int
	interval   time.Duration
	growth     Coal
}

func NewMinerType(
	typeName string,
	cost Coal,
	energy int,
	extraction int,
	interval time.Duration,
	growth Coal,
) MinerType {
	return MinerType{
		typeName:   typeName,
		cost:       cost,
		energy:     energy,
		extraction: extraction,
		interval:   interval,
		growth:     Coal(growth),
	}
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
	coal := m.minerType.extraction + m.actionDone*int(m.minerType.growth)

	m.actionDone++

	finishWorkMiner := m.actionDone >= m.minerType.energy

	return Coal(coal), finishWorkMiner
}

// Метод на получения временного интервала добычи угля
func (m *Miner) IntervalExtraction() time.Duration {
	return m.minerType.interval
}
