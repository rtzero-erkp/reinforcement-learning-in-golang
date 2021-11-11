package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.Agent = &AgentModelMap{}

type AgentModelMap struct {
	search  *common.SearchParam
	model   *common.ModelMap
	encoder common.Encoder
}

func (p *AgentModelMap) String() string {
	return fmt.Sprintf("AgentModelMap:%v, search:%v", p.model, p.search)
}
func (p *AgentModelMap) Train(env common.Env, trainNum int) interface{} {
	var res *common.Result
	mem := common.NewMemCode()
	p.model.Clear()

	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		mem.Clear()
		var envCrt = env.Clone()
		state := env.State()
		// sim to end
		for {
			act := p.model.Sample(env, p.encoder, p.search)
			res = envCrt.Step(act)
			mem.Add(p.encoder.Hash(state), act, p.encoder.Hash(res.State), res.Reward[0])
			state = res.State
			if res.Done {
				break
			}
		}
		// update
		reward := 0.0
		memories := mem.Get()
		memoriesLen := len(memories)
		for i1 := 0; i1 < memoriesLen; i1++ {
			memI := memories[memoriesLen-1-i1]
			reward += memI.Reward
			memI.Reward = reward
			p.model.Update(memI)
		}
	}
	return p.model
}
func (p *AgentModelMap) Policy(env common.Env) (act common.ActionEnum) {
	return p.model.Sample(env, p.encoder, common.NewSearchParam(common.SearchEnum_AvgQ))
}
func NewModelMap(modelMap *common.ModelMap, search *common.SearchParam, encoder common.Encoder) common.Agent {
	var p = &AgentModelMap{
		search:  search,
		model:   modelMap,
		encoder: encoder,
	}
	return p
}
