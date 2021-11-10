package agents

import (
	"gameServer/common"
)

var _ common.Agent = &EpsilonGreed{}

type EpsilonGreed struct {
	epsilon float64          // 概率
	model   *common.ModelMap // 模型
	encoder common.Encoder
}

func (p *EpsilonGreed) Reset() {}
func (p *EpsilonGreed) String() string {
	return "EpsilonGreed"
}
func (p *EpsilonGreed) Policy(state common.Info, space common.Space) common.ActionEnum {
	var code = p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_EpsilonGreed, p.epsilon)
	return act
}
func (p *EpsilonGreed) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var code = p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	node.Accum.Add(act, reward)
}

func NewEpsilonGreed(epsilon float64, mesh common.Encoder) common.Agent {
	var p = &EpsilonGreed{
		epsilon: epsilon,
		model:   common.NewModelMap(common.ModelTypeEnum_Q),
		encoder: mesh,
	}
	return p
}
