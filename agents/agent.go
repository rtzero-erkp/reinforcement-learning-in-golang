package agents

import (
	"gameServer/common"
)

type Agent interface {
	Policy(state []int, space common.Space) (policy common.Policy)
	Reward(state []int, act common.ActionEnum, reward float64)
}
