package agents

import (
	"gameServer/common"
	"math"
	"math/rand"
)

var _ Agent = &SoftMax{}

type SoftMax struct {
	tau   float64     // 概率
	model common.Tree // 模型
}

func (p *SoftMax) Policy(state []int, space common.Space) common.Policy {
	var node = p.model.Find(state)
	var probSum float64 = 0
	for _, act := range space.Acts() {
		var q = node.Accum().Mean(act)
		probSum += math.Exp(q / p.tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	var probMax float64 = 0
	var actsMax []common.ActionEnum
	for _, act := range space.Acts() {
		var q = node.Accum().Mean(act)
		var prob = math.Exp(q/p.tau) / probSum
		probCum += prob
		if probCum > rate {
			node.Policy().Clean()
			node.Policy().Set(act, 1)
			return node.Policy()
		}
		if prob > probMax {
			actsMax = []common.ActionEnum{act}
			probMax = prob
		} else
		if prob == probMax {
			actsMax = append(actsMax, act)
		}
	}
	node.Policy().Clean()
	for _, act := range actsMax {
		node.Policy().Set(act, 1)
	}
	return node.Policy()
}
func (p *SoftMax) Reward(state []int, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state)
	node.Accum().Add(act, reward)
}

func NewSoftMax(tau float64) Agent {
	var p = &SoftMax{
		tau:   tau,
		model: common.NewRootNode(),
	}
	return p
}
