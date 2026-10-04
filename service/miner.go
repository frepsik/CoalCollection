package service

import (
	"CoalCollection/domain"
	"context"
	"time"
)

// Метод необходиый для найма шахтёра и инициализации его как горутины посредствам
// startMinerMine(miner)
func (g *Game) HireMiner(typeMiner string) error {
	minerTypeName := domain.MinerTypeName(typeMiner)

	minerType, exists := g.minerTypes[minerTypeName]
	if !exists {
		return ErrMinerTypeNotFound
	}

	miner, err := g.enterprise.HireMiner(minerType)
	if err != nil {
		return err
	}

	g.startMinerMine(miner)

	return nil
}

// Метод добычи угля определённым шахтёром
func (g *Game) startMinerMine(
	miner *domain.Miner,
) {
	ctxMiner, cancelMine := context.WithCancel(g.ctxGame)

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		defer cancelMine()

		ticker := time.NewTicker(miner.MiningInterval())
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				coal, isExhausted := miner.Mine()
				if coal != 0 {
					g.enterprise.AddCoal(coal)
				}
				if isExhausted {
					return
				}

			case <-ctxMiner.Done():
				return
			}
		}
	}()
}

// Метод для возвращаения списка типов шахтёров, чтобы пользователь мог ознакомиться, что ему можно приобрести
func (g *Game) MinersType() []domain.MinerType {
	result := make([]domain.MinerType, 0, len(g.minerTypes))
	for _, miner := range g.minerTypes {
		result = append(result, miner)
	}
	return result
}

/// У этих трёх методов ниже надо добавить проверку на то, что существует ли вообщем хоть один шахтёр, может список пустой

// Метод для получения работающих сейчас шахтёров
func (g *Game) AvailableMiners() []domain.MinerInfo {
	return g.enterprise.AvailableMiners()
}

// Метод для получения не работающих сейчас шахтёров
func (g *Game) ExhaustedMiners() []domain.MinerInfo {
	return g.enterprise.ExhaustedMiners()
}

// Метод для получения шахтёров по определённому типу
func (g *Game) MinersByType(typeMiner string) ([]domain.MinerInfo, error) {
	typeMinerName := domain.MinerTypeName(typeMiner)
	_, exists := g.minerTypes[typeMinerName]
	if !exists {
		return nil, ErrMinerTypeNotFound
	}
	return g.enterprise.MinersByType(typeMinerName), nil
}
