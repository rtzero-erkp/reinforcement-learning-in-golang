package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.Agent = &AgentModelFree{}

type AgentModelFree struct {
	accum  common.Accumulate
	search *common.SearchParam
}

func (p *AgentModelFree) String() string {
	return fmt.Sprintf("AgentModelFree:%v", p.search)
}
func (p *AgentModelFree) Train(env common.Env, trainNum int) interface{} {
	p.accum.Reset()
	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		envCrt := env.Clone()
		reward := 0.0
		// target
		acts := envCrt.Acts()
		target := p.accum.Sample(acts, p.search)
		act := target
		// simulate
		for {
			res := envCrt.Step(act)
			reward += res.Reward[0]
			if res.Done {
				break
			}
			acts = envCrt.Acts()
			act = acts.Sample()
		}
		// train
		p.accum.Add(target, reward)
	}

	return p.accum
}
func (p *AgentModelFree) Policy(env common.Env) (act common.ActEnum) {
	act = p.accum.Sample(env.Acts(), common.NewSearchParam(common.SearchEnum_AvgQ))
	return
}
func NewModelFree(search *common.SearchParam) common.Agent {
	var p = &AgentModelFree{
		search: search,
		accum:  common.NewAccumulate(),
	}
	return p
}
