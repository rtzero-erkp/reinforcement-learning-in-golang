package agents

import (
	"gameServer/common"
	"math/rand"
)

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64           // 概率
	accum   common.Accumulate // 累计收益
	policy  common.Policy     // 策略梯度
	space   common.Space      // 行动空间
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
			var q = p.accum.Mean(act)
			if (len(actsMax) == 0) || (q > qMax) {
				qMax = q
				actsMax = append(actsMax, act)
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
	p.accum.Add(act, reward)
}

func NewEpsilonGreed(space common.Space, epsilon float64) Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		accum:   common.NewReward1D(space),
		policy:  common.NewPolicyPlus(),
		space:   space,
	}
	return p
}
