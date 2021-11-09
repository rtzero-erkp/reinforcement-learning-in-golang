package agents

import (
	"gameServer/common"
)

type Agent interface {
	Reset()
	Policy(state common.State, space common.Space) (act common.ActionEnum)
	Reward(state common.State, act common.ActionEnum, reward float64)
	String() string
}
