package agents

import (
	"gameServer/common"
)

var _ Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64      // 概率
	model   *common.HashMap // 模型
	mesh    *common.Mesh
}

func (p *EpsilonGreed) Reset() {}
func (p *EpsilonGreed) String() string {
	return "EpsilonGreed"
}
func (p *EpsilonGreed) Policy(state common.State, space common.Space) common.ActionEnum {
	var node = p.model.Find(state, p.mesh)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_EpsilonGreed, p.epsilon)
	return act
}
func (p *EpsilonGreed) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state, p.mesh)
	node.Accum.Add(act, reward)
}

func NewEpsilonGreed(epsilon float64, mesh *common.Mesh) Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		model:   common.NewHashMap(),
		mesh:    mesh,
	}
	return p
}
