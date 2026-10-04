package service

import (
	"CoalCollection/domain"
	"context"
	"time"
)

func (g *Game) startMinerMine(
	miner *domain.Miner,
	ctx context.Context,
	cancelMine context.CancelFunc,
) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		ticker := time.NewTicker(miner.MiningInterval())
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				coal, isExhausted := miner.Mine()
				if isExhausted {
					cancelMine()
					return
				}
				g.enterprise.AddCoal(coal)
			case <-g.ctxGame.Done():
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}
