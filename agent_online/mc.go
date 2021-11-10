package agent_offline

import (
	"gameServer/common"
)

var _ common.AgentOnline = &MC{}

type MC struct {
	env   common.Env
	mcNum int
}

func (p *MC) Reset() {}
func (p *MC) String() string {
	return "MC"
}
func (p *MC) Policy(state common.Info, space common.Space) common.ActionEnum {
	var accum = common.NewAccum()
	var res *common.Result

	for i := 0; i < p.mcNum; i++ {
		var envCrt = p.env.Clone()
		var target = space.Sample()

		var act = target
		res = envCrt.Step(act)
		var reward = res.Reward[0]
		for !res.Done {
			act = envCrt.Space().Sample()
			res = envCrt.Step(act)
			reward += res.Reward[0]
		}
		accum.Add(target, reward)
	}

	return accum.Sample(space, common.SearchMethodEnum_MeanQ)
}
func (p *MC) Reward(state common.Info, act common.ActionEnum, reward float64) {}

func NewMC(env common.Env, mcNum int) common.AgentOnline {
	var p = &MC{
		env:   env,
		mcNum: mcNum,
	}
	return p
}
