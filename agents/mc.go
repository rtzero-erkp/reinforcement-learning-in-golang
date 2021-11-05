package agents

import (
	"gameServer/common"
	"gameServer/envs"
)

var _ Agent = &MC{}

type MC struct {
	env envs.Env
}

func (p *MC) Policy(state common.State, space common.Space) common.Policy {
	//var node = p.model.Find(state.Encode(p.mesh))
	//p.env.Set(state)
	//return node.Policy()
	return nil
}
func (p *MC) Reward(state common.State, act common.ActionEnum, reward float64) {}

func NewMC(env envs.Env) Agent {
	var p = &MC{
		env: env,
	}
	return p
}
