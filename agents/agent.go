package agents

import (
	"gameServer/common"
)

type Agent interface {
	Policy(state common.State, space common.Space) (policy common.Policy)
	Reward(state common.State, act common.ActionEnum, reward float64)
	String() string
}
