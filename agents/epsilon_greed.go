package agents

import (
	"gameServer/common"
)

var _ common.Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64            // 概率
	model   *common.HashPolicy // 模型
	mesh    *common.Mesh
}

func (p *EpsilonGreed) Reset() {}
func (p *EpsilonGreed) String() string {
	return "EpsilonGreed"
}
func (p *EpsilonGreed) Policy(state common.Info, space common.Space) common.ActionEnum {
	var node = p.model.Find(state, p.mesh)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_EpsilonGreed, p.epsilon)
	return act
}
func (p *EpsilonGreed) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewEpsilonGreed(epsilon float64, mesh *common.Mesh) common.Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		model:   common.NewHashPolicy(),
		mesh:    mesh,
	}
	return p
}
