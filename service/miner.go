package service

import (
	"CoalCollection/domain"
	"context"
	"time"
)

// Метод необходимый для найма шахтёра и инициализации его как горутины посредствам
// startMinerMine(miner)
func (g *Game) HireMiner(typeMiner string) (domain.MinerTypeInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()

	if g.gameState != gameRunning {
		return domain.MinerTypeInfo{}, ErrGameIsNotRunning
	}

	minerTypeName := domain.MinerTypeName(typeMiner)

	minerType, exists := g.minerTypes[minerTypeName]
	if !exists {
		return domain.MinerTypeInfo{}, ErrMinerTypeNotFound
	}

	miner, err := g.enterprise.HireMiner(minerType)
	if err != nil {
		return domain.MinerTypeInfo{}, err
	}

	g.startMinerMine(miner)

	return domain.NewMinerTypeInfo(minerType), nil
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

// Метод для возвращения списка типов шахтёров, чтобы пользователь мог ознакомиться, что ему можно приобрести
func (g *Game) MinersType() []domain.MinerType {
	result := make([]domain.MinerType, 0, len(g.minerTypes))
	for _, miner := range g.minerTypes {
		result = append(result, miner)
	}
	return result
}

// Метод для получения работающих сейчас шахтёров
func (g *Game) AvailableMiners() ([]domain.MinerInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return nil, ErrGameIsNotRunning
	}
	return g.enterprise.AvailableMiners(), nil
}

// Метод для получения не работающих сейчас шахтёров
func (g *Game) ExhaustedMiners() ([]domain.MinerInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return nil, ErrGameIsNotRunning
	}

	return g.enterprise.ExhaustedMiners(), nil
}

// Метод для получения шахтёров по определённому типу
func (g *Game) MinersByType(typeMiner string) ([]domain.MinerInfo, error) {
	g.lifecycleMu.RLock()
	defer g.lifecycleMu.RUnlock()
	if g.gameState != gameRunning {
		return nil, ErrGameIsNotRunning
	}

	typeMinerName := domain.MinerTypeName(typeMiner)
	_, exists := g.minerTypes[typeMinerName]
	if !exists {
		return nil, ErrMinerTypeNotFound
	}
	return g.enterprise.MinersByType(typeMinerName), nil
}
