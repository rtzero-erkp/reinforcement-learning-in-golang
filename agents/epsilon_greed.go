package agents

import (
	"gameServer/common"
	"math/rand"
)

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64      // 概率
	model   common.Model // 模型
	mesh    *common.Mesh
}

func (p *EpsilonGreed) String() string {
	return "EpsilonGreed"
}
func (p *EpsilonGreed) Policy(state common.State, space common.Space) common.Policy {
	var node = p.model.Find(state, p.mesh)
	var rate = rand.Float64()
	if rate < p.epsilon {
		node.Policy.Clean()
		node.Policy.Set(space.Sample(), 1)
	} else {
		var qMax float64
		var actsMax []common.ActionEnum

		for _, act := range space.Acts() {
			var q = node.Accum.Mean(act)
			if (len(actsMax) == 0) || (q > qMax) {
				qMax = q
				actsMax = append(actsMax, act)
			} else
			if q == qMax {
				actsMax = append(actsMax, act)
			}
		}

		node.Policy.Clean()
		for _, act := range actsMax {
			node.Policy.Set(act, 1)
		}
	}
	return node.Policy
}
func (p *EpsilonGreed) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewEpsilonGreed(epsilon float64, mesh *common.Mesh) Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		model:   common.NewTree(),
		mesh:    mesh,
	}
	return p
}
