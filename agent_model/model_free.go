package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.AgentModel = &ModelFree{}

type ModelFree struct {
	accum  common.Accumulate
	search *common.SearchParam
}

func (p *ModelFree) String() string {
	return fmt.Sprintf("ModelFree:%v", p.search)
}
func (p *ModelFree) Train(env common.Env, trainNum int) interface{} {
	p.accum.Reset()
	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		envCrt := env.Clone()
		reward := 0.0
		// target
		target := p.accum.Sample(env.Space(), p.search)
		act := target
		// simulate
		for {
			res := envCrt.Step(act)
			reward += res.Reward[0]
			if res.Done {
				break
			}
			act = env.Space().Sample()
		}
		// train
		p.accum.Add(target, reward)
	}
	return p.accum
}
func (p *ModelFree) Policy(env common.Env) (act common.ActionEnum) {
	act = p.accum.Sample(env.Space(), common.NewSearchParam(common.SearchEnum_AvgQ))
	return
}
func NewModelFree(search *common.SearchParam) common.AgentModel {
	var p = &ModelFree{
		search: search,
		accum:  common.NewAccumulate(),
	}
	return p
}
