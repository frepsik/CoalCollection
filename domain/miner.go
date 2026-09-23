package domain

import (
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MinerTypeName string

const (
	SmallMiner  MinerTypeName = "smallMiner"
	NormalMiner MinerTypeName = "normalMiner"
	StrongMiner MinerTypeName = "strongMiner"
)

type MinerType struct {
	typeName   MinerTypeName
	name       string
	cost       Coal
	energy     int
	extraction Coal
	interval   time.Duration
	growth     Coal
}

func NewMinerType(
	typeName MinerTypeName,
	name string,
	cost Coal,
	energy int,
	extraction Coal,
	interval time.Duration,
	growth Coal,
) (MinerType, error) {

	minerType := MinerType{
		typeName:   typeName,
		name:       name,
		cost:       cost,
		energy:     energy,
		extraction: extraction,
		interval:   interval,
		growth:     growth,
	}

	if err := minerType.validate(); err != nil {
		return MinerType{}, err
	}

	return minerType, nil
}

func (mt *MinerType) validate() error {
	if !(mt.typeName == SmallMiner || mt.typeName == NormalMiner || mt.typeName == StrongMiner) {
		return ErrInvalidMinerTypeName
	}
	if strings.TrimSpace(mt.name) == "" {
		return ErrInvalidMinerName
	}
	if mt.cost < 0 {
		return ErrInvalidMinerCost
	}
	if mt.energy <= 0 {
		return ErrInvalidMinerEnergy
	}
	if mt.extraction <= 0 {
		return ErrInvalidMinerExtraction
	}
	if mt.interval <= 0 {
		return ErrInvalidMinerInterval
	}
	if mt.growth < 0 {
		return ErrInvalidMinerGrowth
	}

	return nil
}

type Miner struct {
	id         uuid.UUID
	minerType  MinerType
	actionDone int
	mtx        sync.Mutex
}

func newMiner(minerType MinerType) *Miner {
	return &Miner{
		id:        uuid.New(),
		minerType: minerType,
	}
}

// Метод на добычу угля. true - в случае, если шахтёр не может работать (нет энергии более)
func (m *Miner) Mine() (Coal, bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.actionDone >= m.minerType.energy {
		return 0, true
	}

	coal := m.minerType.extraction + Coal(m.actionDone)*m.minerType.growth

	m.actionDone++

	finishWorkMiner := m.actionDone >= m.minerType.energy

	return Coal(coal), finishWorkMiner
}

// Метод на получения временного интервала добычи угля
func (m *Miner) MiningInterval() time.Duration {
	return m.minerType.interval
}

// Метод на проверку типа шахтёра
func (m *Miner) isType(minerTypeName MinerTypeName) bool {
	return m.minerType.typeName == minerTypeName
}

// Метод для того, чтобы определить, шахтёр может работать или нет
func (m *Miner) isExhausted() bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.actionDone >= m.minerType.energy
}

// Метод в котором произведём расчёт того, сколько осталось энергии,
// а также соберём некоторый снимок по текущему шахтёру на вывод, чтобы не возвращать структуру, в рамках которой присутствует mtx
func (m *Miner) info() MinerInfo {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	leftEnergy := m.minerType.energy - m.actionDone
	currentExtraction := m.minerType.extraction + Coal(m.actionDone)*m.minerType.growth
	minerInfo := newMinerInfo(
		m.id,
		m.minerType.typeName,
		m.minerType.name,
		m.minerType.cost,
		leftEnergy,
		currentExtraction,
		m.minerType.interval,
		m.minerType.growth,
	)
	return minerInfo
}
