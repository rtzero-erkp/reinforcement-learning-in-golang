package envs

import (
	"gameServer/common"
	"log"
	"math/rand"
)

type Env interface {
	String() string
	ActionSpace() *common.Space
	Seed(seed int64) rand.Source
	Step(act common.ActionEnum) (state common.Stater, reward float64, done bool)
	Reset() common.Stater
	Close()
}

func Envs(game common.GameEnum) Env {
	switch game {
	case common.GameEnum_CartPole_v0:
		return NewCartPoleEnv()
	case common.GameEnum_Bandits_v0:
		return NewBanditsEnv()
	default:
		log.Fatalf("unknown game enum:%v", game)
		return nil
	}
}
