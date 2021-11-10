package agent_offline

import (
	"gameServer/common"
)

var _ common.AgentOnline = &SoftMax{}

type SoftMax struct {
	tau     float64          // 概率
	model   *common.ModelMap // 模型
	encoder common.Encoder
}

func (p *SoftMax) Reset() {}
func (p *SoftMax) String() string {
	return "SoftMax"
}
func (p *SoftMax) Policy(state common.Info, space common.Space) common.ActionEnum {
	var code = p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_SoftMax, p.tau)
	return act
}
func (p *SoftMax) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var code = p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	node.Accum.Add(act, reward)
}

func NewSoftMax(tau float64, mesh common.Encoder) common.AgentOnline {
	var p = &SoftMax{
		tau:     tau,
		model:   common.NewModelMap(common.ModelTypeEnum_Q),
		encoder: mesh,
	}
	return p
}
