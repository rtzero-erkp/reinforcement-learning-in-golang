package agents

import "gameServer/common"

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon   float64                       // 概率
	rewardMap map[common.ActionEnum]float64 // reward accum
	countMap  map[common.ActionEnum]float64 // count accum
	policy    *common.Policy
}

func (p *EpsilonGreed) Policy(*common.Space) *common.Policy {
	return p.policy
}

func (p *EpsilonGreed) Reward(act common.ActionEnum, reward float64) {
}

func NewEpsilonGreed(epsilon float64) Agent {
	var p = &EpsilonGreed{
		epsilon:   epsilon,
		rewardMap: map[common.ActionEnum]float64{},
		countMap:  map[common.ActionEnum]float64{},
		policy:    &common.Policy{},
	}
	return p
}
