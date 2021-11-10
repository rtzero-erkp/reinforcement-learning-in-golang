package agents

import (
	"gameServer/common"
)

var _ common.Agent = &Cfr{}

type Cfr struct {
	model *common.ModelTree // 模型
}

func (p *Cfr) Reset() {
	p.model.Clear()
}
func (p *Cfr) String() string {
	return "Cfr"
}
func (p *Cfr) Policy(state common.Info, space common.Space) common.ActionEnum {
	//p.accum.Reset()
	//for _, act := range space.Acts() {
	//	var cp = p.env.Clone()
	//	cp.Set(state)
	//	var stateCrt = cp.Step(act).State
	//	var nodeCrt = p.model.Find(stateCrt, p.encoder)
	//	p.accum.Add(act, nodeCrt.Value)
	//}
	//return p.accum.Sample(space, p.method, p.args...)
	return space.Sample()
}
func (p *Cfr) Reward(state common.Info, act common.ActionEnum, reward float64) {
	//var cp = p.env.Clone()
	//cp.Set(state)
	//var stateDot = cp.Step(act).State
	//var node = p.model.Find(state, p.encoder)
	//var nodeDot = p.model.Find(stateDot, p.encoder)
	//var vs = node.Value
	//var vsDot = nodeDot.Value
	//// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
	//node.Value = vs + p.alpha*(reward+p.lambda*(vsDot-vs))
}

func NewCFR() common.Agent {
	var p = &Cfr{
		model: common.NewModelTree(common.ModelTypeEnum_Cfr),
	}
	return p
}
