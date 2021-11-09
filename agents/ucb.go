package agents

import (
	"gameServer/common"
)

var _ Agent = &UCB{}

type UCB struct {
	model *common.HashPolicy // 模型
	mesh  *common.Mesh
}

func (p *UCB) Reset() {}
func (p *UCB) String() string {
	return "UCB"
}
func (p *UCB) Policy(state common.State, space common.Space) common.ActionEnum {
	var node = p.model.Find(state, p.mesh)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_UCB)
	return act
}
func (p *UCB) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewUCB(mesh *common.Mesh) Agent {
	var p = &UCB{
		model: common.NewHashPolicy(),
		mesh:  mesh,
	}
	return p
}
