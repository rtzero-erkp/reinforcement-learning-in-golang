package agents

import (
	"gameServer/common"
	"math"
)

var _ Agent = &UCB{}

type UCB struct {
	epsilon float64           // 概率
	accum   common.Accumulate // 累计收益
	policy  common.Policy     // 策略梯度
	space   common.Space      // 行动空间
}

func (p *UCB) Policy(common.Space) common.Policy {
	countSum := p.accum.Count()
	if countSum == 0 {
		p.policy.Clean()
		p.policy.Set(p.space.Sample(), 1)
		return p.policy
	}
	var actsMax []common.ActionEnum
	var ucbMax float64
	for _, act := range p.space.Acts() {
		count := p.accum.CountAct(act)
		if count == 0 {
			p.policy.Clean()
			p.policy.Set(act, 1)
			return p.policy
		}
		upperBound := math.Sqrt((2 * math.Log(countSum)) / count)
		q := p.accum.Mean(act)
		ucb := upperBound + q
		if (len(actsMax) == 0) || ucb > ucbMax {
			ucbMax = ucb
			actsMax = []common.ActionEnum{act}
		} else
		if ucb == ucbMax {
			actsMax = append(actsMax, act)
		}
	}
	p.policy.Clean()
	for _, act := range actsMax {
		p.policy.Set(act, 1)
	}
	return p.policy
}
func (p *UCB) Reward(act common.ActionEnum, reward float64) {
	p.accum.Add(act, reward)
}

func NewUCB(space common.Space, epsilon float64) Agent {
	var p = &UCB{
		epsilon: epsilon,
		accum:   common.NewReward1D(space),
		policy:  common.NewPolicyPlus(),
		space:   space,
	}
	return p
}
