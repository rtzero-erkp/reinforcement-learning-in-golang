package agents

import (
	"gameServer/common"
	"gameServer/envs"
)

var _ Agent = &MCTS{}

type MCTS struct {
	env      envs.Env
	mcNum    int
	method   common.SearchMethod
	arg      []interface{}
	model    *common.HashPolicy
	memPath  *common.MemPath
	pathHead string
}

func (p *MCTS) Reset() {
	p.model.Clear()
	p.memPath.Clear()
	p.pathHead = ""
}
func (p *MCTS) String() string {
	return "MCTS:" + p.method.String()
}
func (p *MCTS) Policy(state common.State, space common.Space) common.ActionEnum {
	var res *envs.Result
	var mem *common.Memory2

	for i0 := 0; i0 < p.mcNum; i0++ {
		// reset
		p.memPath.Clear()
		var envCrt = p.env.Clone()
		var path = p.pathHead
		// find leaf point
		for {
			var node = p.model.Find2(path)
			var spaceCrt = p.env.ActionSpace()
			var act = node.Accum.Sample(spaceCrt, p.method, p.arg...)
			res = envCrt.Step(act)
			p.memPath.Add(path, act, res.Reward)
			path += act.String() + " "
			if !res.Done {
				break
			}
			if node.Accum.CountAct(act) == 0 {
				break
			}
		}
		// rand act to end
		var reward float64 = 0
		for !res.Done {
			var spaceCrt = p.env.ActionSpace()
			var act = spaceCrt.Sample()
			res = envCrt.Step(act)
			reward += res.Reward
		}
		// update
		var memories = p.memPath.Get()
		var memoriesLen = len(memories)
		for i1 := 0; i1 < memoriesLen; i1++ {
			mem = memories[memoriesLen-1-i1]
			var node = p.model.Find2(mem.Path)
			node.Accum.Add(mem.Act, reward)
			reward = node.Accum.Mean()
		}
	}

	var node = p.model.Find2(p.pathHead)
	var act = node.Accum.Sample(space, common.SearchMethodEnum_MeanQ)
	p.pathHead += act.String() + " "
	return act
}
func (p *MCTS) Reward(state common.State, act common.ActionEnum, reward float64) {}

func NewMCTS(env envs.Env, mcNum int, method common.SearchMethod, arg ...interface{}) Agent {
	var p = &MCTS{
		env:      env,
		mcNum:    mcNum,
		method:   method,
		arg:      arg,
		model:    common.NewHashPolicy(),
		memPath:  common.NewMemPath(),
		pathHead: "",
	}
	return p
}
