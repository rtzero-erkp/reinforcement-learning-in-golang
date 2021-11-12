package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.Agent = &AgentModelTreeExpand{}

type AgentModelTreeExpand struct {
	search *common.SearchMethod
	model  *common.ModelTree
}

func (p *AgentModelTreeExpand) String() string {
	return fmt.Sprintf("AgentModelTreeExpand:%v, search:%v", p.model, p.search)
}
func (p *AgentModelTreeExpand) Train(env common.Env, trainNum int) interface{} {
	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		envCrt := env.Clone()
		reward := 0.0
		node := p.model.Find()
		// sim to end
		var endAct common.ActEnum
		var exist bool
		var res = &common.Result{}
		for {
			endAct = node.Sample(envCrt, p.search)
			exist = node.Exist(endAct)
			res = envCrt.Step(endAct)
			node = node.Find(endAct)
			reward += res.Reward[0]
			if res.Done {
				break
			}
			if !exist {
				break
			}
		}
		for !res.Done {
			act := envCrt.Acts().Sample()
			res = envCrt.Step(act)
			reward += res.Reward[0]
		}
		// update
		node.Update(reward)
	}
	return p.model
}
func (p *AgentModelTreeExpand) Policy(env common.Env) (act common.ActEnum) {
	node := p.model.Find()
	act = node.Sample(env, common.NewSearchMethod(common.SearchEnum_AvgQ))
	p.model.Move(act)
	return
}
func NewModelTreeExpand(modelTree *common.ModelTree, search *common.SearchMethod) common.Agent {
	var p = &AgentModelTreeExpand{
		search: search,
		model:  modelTree,
	}
	return p
}
