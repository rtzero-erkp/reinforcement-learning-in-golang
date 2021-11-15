package agent_model

import (
	"fmt"
	"gameServer/common"
)

var _ common.Agent = &AgentModelMap{}

type AgentModelMap struct {
	search  *common.SearchMethod
	model   *common.ModelMap
	update  *common.UpdateMethod
	encoder common.Encoder
}

func (p *AgentModelMap) String() string {
	return fmt.Sprintf("AgentModelMap:%v, search:%v", p.model, p.search)
}
func (p *AgentModelMap) Train(env common.Env, trainNum int) interface{} {
	var res *common.Result
	mem := common.NewMemCode()

	for i0 := 0; i0 < trainNum; i0++ {
		//log.Printf("[agent] ===[i0:%v]===", i0)
		// reset
		mem.Clear()
		var envCrt = env.Clone()
		state := envCrt.State()
		// sim to end
		//first := true
		for {
			//log.Println("[agent] ===[step]===")
			//if first {
			//	log.Println("[agent] ===[first]===")
			//}
			act := p.model.Sample(envCrt, p.encoder, p.search)
			//log.Printf("[Mark]")
			//log.Printf("act:%v", act)
			//if first {
			//log.Printf("[agent] act:%10v, acts:%v", act, envCrt.Acts())
			//	log.Println("[agent] ---[first]---")
			//}
			//break
			res = envCrt.Step(act)
			mem.Add(p.encoder.Hash(state), act, p.encoder.Hash(res.State), res.Reward[0])
			//break
			state = res.State
			//first = false
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
			//log.Printf("i1:%v, mem:%v, reward:%v", i1, memI, reward)
			memI.Reward = reward
			p.update.QMap(memI, p.model)
		}
		//log.Println("[agent] ===[test]===")
		//code := p.encoder.Hash(env.State())
		//node := p.model.Find(code).(common.Accumulate)
		//log.Printf("[agent] acts:%v", env.Acts())
		//log.Printf("[agent] node:%v", node)
		//log.Println("[agent] ---[test]---")
	}
	return p.model
}
func (p *AgentModelMap) Policy(env common.Env) (act common.ActEnum) {
	act = p.model.Sample(env, p.encoder, common.NewSearchMethod(common.SearchEnum_AvgQ))
	//code := p.encoder.Hash(env.State())
	//node := p.model.Find(code).(common.Accumulate)
	//log.Printf("code:%v", code)
	//log.Printf("acts:%v", env.Acts())
	//log.Println(node)
	return
}
func NewModelMap(modelMap *common.ModelMap,
	search *common.SearchMethod,
	update *common.UpdateMethod,
	encoder common.Encoder) common.Agent {
	var p = &AgentModelMap{
		model:   modelMap,
		search:  search,
		update:  update,
		encoder: encoder,
	}
	return p
}
