package agent_offline

import (
	"fmt"
	"gameServer/common"
)

var _ common.AgentOffline = &Search{}

type Search struct {
	model   *common.ModelMap // 模型
	encoder common.Encoder
	env     common.Env
	method  common.SearchMethod
	args    []interface{}
}

func (p *Search) Reset() {
	p.model = common.NewModelMap(common.ModelTypeEnum_Q)
}
func (p *Search) String() string {
	return fmt.Sprintf("Search:%v", p.method)
}
func (p *Search) Train(trainNum int) {
	mem := common.NewMem()
	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		mem.Clear()
		envCrt := p.env.Clone()
		state := envCrt.State()
		// simulate
		for {
			var code = p.encoder.Hash(state)
			var node = p.model.Find(code).(*common.NodeQ)
			var act = node.Accum.Sample(p.env.Space(), p.method, p.args...)
			res := p.env.Step(act)
			mem.Add(state, act, res.Reward[0])
			state = res.State
			if res.Done {
				break
			}
		}
		// train
		var reward float64 = 0
		steps := mem.Get()
		stepsNum := len(steps)
		for i1 := 0; i1 < stepsNum; i1++ {
			step := steps[stepsNum-1-i1]
			reward += step.Reward
			var code = p.encoder.Hash(step.State)
			var node = p.model.Find(code).(*common.NodeQ)
			node.Accum.Add(step.Act, reward)
		}
	}
}
func (p *Search) Policy(state common.Info, space common.Space) common.ActionEnum {
	var code = p.encoder.Hash(state)
	var node = p.model.Find(code).(*common.NodeQ)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_MeanQ)
	return act
}

func NewSearch(env common.Env, mesh common.Encoder, method common.SearchMethod, args ...interface{}) common.AgentOffline {
	var p = &Search{
		env:     env,
		encoder: mesh,
		method:  method,
		args:    args,
		model:   common.NewModelMap(common.ModelTypeEnum_Q),
	}
	return p
}
