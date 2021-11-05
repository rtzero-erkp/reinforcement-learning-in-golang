package envs

import (
	"gameServer/common"
	"math/rand"
)

type Env interface {
	String() string
	ActionSpace() common.Space
	Seed(seed int64) rand.Source
	Step(act common.ActionEnum) (state common.Stater, reward float64, done bool)
	Reset() common.Stater
	Close()
}
