package agents

import (
	"gameServer/common"
	"gameServer/envs"
)

var _ Agent = &MC{}

type MC struct {
	env   envs.Env
	mcNum int
}

func (p *MC) Reset() {}
func (p *MC) String() string {
	return "MC"
}
func (p *MC) Policy(state common.State, space common.Space) common.ActionEnum {
	var accum = common.NewAccum()
	var res *envs.Result

	for i := 0; i < p.mcNum; i++ {
		var envCrt = p.env.Clone()
		var target = space.Sample()

		var act = target
		res = envCrt.Step(act)
		var reward = res.Reward
		for !res.Done {
			act = envCrt.ActionSpace().Sample()
			res = envCrt.Step(act)
			reward += res.Reward
		}
		accum.Add(target, reward)
	}

	return accum.Sample(space, common.SearchMethodEnum_MeanQ)
}
func (p *MC) Reward(state common.State, act common.ActionEnum, reward float64) {}

func NewMC(env envs.Env, mcNum int) Agent {
	var p = &MC{
		env:   env,
		mcNum: mcNum,
	}
	return p
}
