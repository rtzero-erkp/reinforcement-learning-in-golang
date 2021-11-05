package agents

import (
	"gameServer/common"
	"math/rand"
)

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64     // 概率
	model   common.Tree // 模型
	mesh    []float64
}

func (p *EpsilonGreed) Policy(state common.State, space common.Space) common.Policy {
	var node = p.model.Find(state.Encode(p.mesh))
	var rate = rand.Float64()
	if rate < p.epsilon {
		node.Policy().Clean()
		node.Policy().Set(space.Sample(), 1)
	} else {
		var qMax float64
		var actsMax []common.ActionEnum

		for _, act := range space.Acts() {
			var q = node.Accum().Mean(act)
			if (len(actsMax) == 0) || (q > qMax) {
				qMax = q
				actsMax = append(actsMax, act)
			} else
			if q == qMax {
				actsMax = append(actsMax, act)
			}
		}

		node.Policy().Clean()
		for _, act := range actsMax {
			node.Policy().Set(act, 1)
		}
	}
	return node.Policy()
}
func (p *EpsilonGreed) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state.Encode(p.mesh))
	node.Accum().Add(act, reward)
}

func NewEpsilonGreed(epsilon float64, mesh []float64) Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		model:   common.NewRootNode(),
		mesh:    mesh,
	}
	return p
}
