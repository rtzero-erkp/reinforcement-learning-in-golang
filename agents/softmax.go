package agents

import (
	"gameServer/common"
)

var _ common.Agent = &SoftMax{}

type SoftMax struct {
	tau   float64            // 概率
	model common.ModelPolicy // 模型
	mesh  common.Encoder
}

func (p *SoftMax) Reset() {}
func (p *SoftMax) String() string {
	return "SoftMax"
}
func (p *SoftMax) Policy(state common.Info, space common.Space) common.ActionEnum {
	var node = p.model.Find(state, p.mesh)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_SoftMax, p.tau)
	return act
}
func (p *SoftMax) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewSoftMax(tau float64, mesh common.Encoder) common.Agent {
	var p = &SoftMax{
		tau:   tau,
		model: common.NewHashPolicy(),
		mesh:  mesh,
	}
	return p
}
