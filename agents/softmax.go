package agents

import (
	"gameServer/common"
)

var _ Agent = &SoftMax{}

type SoftMax struct {
	tau   float64      // 概率
	model *common.HashMap // 模型
	mesh  *common.Mesh
}

func (p *SoftMax) Reset() {}
func (p *SoftMax) String() string {
	return "SoftMax"
}
func (p *SoftMax) Policy(state common.State, space common.Space) common.ActionEnum {
	var node = p.model.Find(state, p.mesh)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_SoftMax, p.tau)
	return act
}
func (p *SoftMax) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewSoftMax(tau float64, mesh *common.Mesh) Agent {
	var p = &SoftMax{
		tau:   tau,
		model: common.NewHashMap(),
		mesh:  mesh,
	}
	return p
}
