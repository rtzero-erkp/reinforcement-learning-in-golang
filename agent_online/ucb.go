package agent_offline

import (
	"gameServer/common"
)

var _ common.AgentOnline = &UCB{}

type UCB struct {
	model   *common.ModelMap // 模型
	encoder common.Encoder
}

func (p *UCB) Reset() {}
func (p *UCB) String() string {
	return "UCB"
}
func (p *UCB) Policy(state common.Info, space common.Space) common.ActionEnum {
	var code =  p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_UCB)
	return act
}
func (p *UCB) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var code =  p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	node.Accum.Add(act, reward)
}

func NewUCB(mesh common.Encoder) common.AgentOnline {
	var p = &UCB{
		model:   common.NewModelMap(common.ModelTypeEnum_Q),
		encoder: mesh,
	}
	return p
}
