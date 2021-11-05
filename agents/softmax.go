package agents

import (
	"gameServer/common"
	"math"
	"math/rand"
)

var _ Agent = &SoftMax{}

type SoftMax struct {
	tau    float64           // 概率
	accum  common.Accumulate // 累计收益
	policy common.Policy     // 策略梯度
	space  common.Space      // 行动空间
}

func (p *SoftMax) Policy(common.Space) common.Policy {
	var probSum float64 = 0
	for _, act := range p.space.Acts() {
		var q = p.accum.Mean(act)
		probSum += math.Exp(q / p.tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	var probMax float64 = 0
	var actsMax []common.ActionEnum
	for _, act := range p.space.Acts() {
		var q = p.accum.Mean(act)
		var prob = math.Exp(q/p.tau) / probSum
		probCum += prob
		if probCum > rate {
			p.policy.Clean()
			p.policy.Set(act, 1)
			return p.policy
		}
		if prob > probMax {
			actsMax = []common.ActionEnum{act}
			probMax = prob
		} else
		if prob == probMax {
			actsMax = append(actsMax, act)
		}
	}
	p.policy.Clean()
	for _, act := range actsMax {
		p.policy.Set(act, 1)
	}
	return p.policy
}
func (p *SoftMax) Reward(act common.ActionEnum, reward float64) {
	p.accum.Add(act, reward)
}

func NewSoftMax(space common.Space, tau float64) Agent {
	var p = &SoftMax{
		tau:    tau,
		accum:  common.NewReward1D(space),
		policy: common.NewPolicyPlus(),
		space:  space,
	}
	return p
}
