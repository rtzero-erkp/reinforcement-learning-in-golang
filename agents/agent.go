package agents

import (
	"gameServer/common"
)

type Agent interface {
	Policy(*common.Space) *common.Policy
	Reward(act common.ActionEnum, reward float64)
}
