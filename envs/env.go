package envs

import (
	"log"
	"math/rand"
)

type Env interface {
	String() string
	ActionSpace() AuctionSpacer
	Seed(seed int64) rand.Source
	Step(act ActionEnum) (state Stater, reward float64, done bool)
	reset() Stater
	Close()
}

func Make(game GameEnum) Env {
	switch game {
	case GameEnum_CartPole_v0:
		return NewCartPoleEnv()
	default:
		log.Fatalf("unknown game enum:%v", game)
		return nil
	}
}
