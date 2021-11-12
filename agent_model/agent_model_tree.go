package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.Agent = &AgentModelTree{}

type AgentModelTree struct {
	search *common.SearchParam
	model  *common.ModelTree
}

func (p *AgentModelTree) String() string {
	return fmt.Sprintf("AgentModelTree:%v, search:%v", p.model, p.search)
}
func (p *AgentModelTree) Train(env common.Env, trainNum int) interface{} {
	// TODO: expand, total
	for i0 := 0; i0 < trainNum; i0++ {
		// reset
		envCrt := env.Clone()
		reward := 0.0
		node := p.model.Find()
		// sim to end
		for {
			act := node.Sample(envCrt, p.search)
			res := envCrt.Step(act)
			node = node.Find(act)
			reward += res.Reward[0]
			if res.Done {
				break
			}
		}
		// update
		node.Update(reward)
	}
	return p.model
}
func (p *AgentModelTree) Policy(env common.Env) (act common.ActEnum) {
	node := p.model.Find()
	act = node.Sample(env, common.NewSearchParam(common.SearchEnum_AvgQ))
	p.model.Move(act)
	return
}
func NewModelTree(modelTree *common.ModelTree, search *common.SearchParam) common.Agent {
	var p = &AgentModelTree{
		search: search,
		model:  modelTree,
	}
	return p
}
