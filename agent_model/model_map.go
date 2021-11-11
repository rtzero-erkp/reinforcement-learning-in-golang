package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.AgentModel = &ModelMap{}

type ModelMap struct {
	model    common.ModelEnum
	update   *common.UpdateParam
	search   *common.SearchParam
	modelMap *common.ModelMap
	encoder  common.Encoder
}

func (p *ModelMap) String() string {
	return fmt.Sprintf("ModelMap:%v, search:%v", p.modelMap, p.search)
}
func (p *ModelMap) Train(env common.Env, trainNum int) interface{} {
	var res *common.Result
	mem := common.NewMemCode()
	p.modelMap.Clear()

	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		mem.Clear()
		var envCrt = env.Clone()
		state := env.State()
		// sim to end
		for {
			act := p.modelMap.Sample(env, p.encoder, p.search)
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
			p.modelMap.Update(memI)
		}
	}
	return p.modelMap
}
func (p *ModelMap) Policy(env common.Env) (act common.ActionEnum) {
	return p.modelMap.Sample(env, p.encoder, common.NewSearchParam(common.SearchEnum_AvgQ))
}
func NewModelMap(model common.ModelEnum, update *common.UpdateParam, search *common.SearchParam, encoder common.Encoder) common.AgentModel {
	var p = &ModelMap{
		model:    model,
		update:   update,
		search:   search,
		modelMap: common.NewModelMap(model, update),
		encoder:  encoder,
	}
	return p
}
