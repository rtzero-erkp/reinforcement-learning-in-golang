package agent_offline

import (
	"gameServer/common"
)

var _ common.AgentOnline = &DT{}

type DT struct {
	model   *common.ModelMap // 模型
	encoder common.Encoder
	accum   common.Accumulate
	env     common.Env
	alpha   float64
	lambda  float64
	args    []interface{}
	method  common.SearchMethod
}

func (p *DT) Reset() {
	p.model.Clear()
}
func (p *DT) String() string {
	return "DT:" + p.method.String()
}
func (p *DT) Policy(state common.Info, space common.Space) common.ActionEnum {
	p.accum.Reset()
	for _, act := range space.Acts() {
		var cp = p.env.Clone()
		cp.Set(state)
		var stateCrt = cp.Step(act).State
		var code = p.encoder.Hash(stateCrt)
		var nodeCrt = p.model.Find(code).(*common.NodeValue)
		p.accum.Add(act, nodeCrt.Value)
	}
	return p.accum.Sample(space, p.method, p.args...)
}
func (p *DT) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var cp = p.env.Clone()
	cp.Set(state)
	var stateDot = cp.Step(act).State
	var code = p.encoder.Hash(state)
	var codeDot = p.encoder.Hash(stateDot)
	var node = p.model.Find(code).(*common.NodeValue)
	var nodeDot = p.model.Find(codeDot).(*common.NodeValue)
	var vs = node.Value
	var vsDot = nodeDot.Value
	// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
	node.Value = vs + p.alpha*(reward+p.lambda*(vsDot-vs))
}

func NewDT(env common.Env, alpha float64, lambda float64, mesh common.Encoder, method common.SearchMethod, args ...interface{}) common.AgentOnline {
	var p = &DT{
		alpha:   alpha,  // 0.1
		lambda:  lambda, // 0.5
		env:     env,
		model:   common.NewModelMap(common.ModelTypeEnum_Value),
		encoder: mesh,
		accum:   common.NewAccum(),
		method:  method,
		args:    args,
	}
	return p
}
