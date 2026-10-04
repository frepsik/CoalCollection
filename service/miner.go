package service

import (
	"CoalCollection/domain"
	"context"
	"time"
)

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
