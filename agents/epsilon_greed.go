package agents

import (
	"gameServer/common"
	"math/rand"
)

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon   float64                       // 概率
	rewardCum map[common.ActionEnum]float64 // 累计收益
	countCum  map[common.ActionEnum]float64 // 累计次数
	policy    common.Policy                 // 策略梯度
	space     common.Space                  // 行动空间
}

func (p *EpsilonGreed) Policy(common.Space) common.Policy {
	var rate = rand.Float64()
	if rate < p.epsilon {
		p.policy.Clean()
		p.policy.Set(p.space.Sample(), 1)
	} else {
		var qMax float64
		var actsMax []common.ActionEnum

		for _, act := range p.space.Acts() {
			var q = p.rewardMean(act)
			if len(actsMax) == 0 {
				qMax = q
				actsMax = append(actsMax, act)
			} else
			if q > qMax {
				qMax = q
				actsMax = []common.ActionEnum{act}
			} else
			if q == qMax {
				actsMax = append(actsMax, act)
			}
		}

		p.policy.Clean()
		for _, act := range actsMax {
			p.policy.Set(act, 1)
		}
	}
	return p.policy
}
func (p *EpsilonGreed) Reward(act common.ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
}

func NewEpsilonGreed(space common.Space, epsilon float64) Agent {
	var p = &EpsilonGreed{
		epsilon:   epsilon,
		rewardCum: map[common.ActionEnum]float64{},
		countCum:  map[common.ActionEnum]float64{},
		policy:    common.NewPolicyPlus(),
		space:     space,
	}
	for _, act := range p.space.Acts() {
		p.rewardCum[act] = 0
		p.countCum[act] = 0
	}
	return p
}

func (p *EpsilonGreed) rewardMean(act common.ActionEnum) float64 {
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}
