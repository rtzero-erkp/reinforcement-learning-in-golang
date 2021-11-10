package agents

import (
	"gameServer/common"
)

var _ common.Agent = &DT{}

type DT struct {
	model  *common.HashValue // 模型
	mesh   *common.Mesh
	accum  common.Accumulate
	env    common.Env
	alpha  float64
	lambda float64
	args   []interface{}
	method common.SearchMethod
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
		var nodeCrt = p.model.Find(stateCrt, p.mesh)
		p.accum.Add(act, nodeCrt.Value)
	}
	return p.accum.Sample(space, p.method, p.args...)
}
func (p *DT) Reward(state common.Info, act common.ActionEnum, reward float64) {
	var cp = p.env.Clone()
	cp.Set(state)
	var stateDot = cp.Step(act).State
	var node = p.model.Find(state, p.mesh)
	var nodeDot = p.model.Find(stateDot, p.mesh)
	var vs = node.Value
	var vsDot = nodeDot.Value
	// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
	node.Value = vs + p.alpha*(reward+p.lambda*(vsDot-vs))
}

func NewDT(env common.Env, alpha float64, lambda float64, mesh *common.Mesh, method common.SearchMethod, args ...interface{}) common.Agent {
	var p = &DT{
		alpha:  alpha,  // 0.1
		lambda: lambda, // 0.5
		env:    env,
		model:  common.NewHashValue(),
		mesh:   mesh,
		accum:  common.NewAccum(),
		method: method,
		args:   args,
	}
	return p
}
