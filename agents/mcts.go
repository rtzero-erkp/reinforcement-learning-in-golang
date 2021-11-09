package agents

import (
	"gameServer/common"
	"gameServer/envs"
)

var _ Agent = &MCTS{}

type MCTS struct {
	env   envs.Env
	model common.Model
	mcNum int
}

func (p *MCTS) String() string {
	return "MCTS"
}
func (p *MCTS) Policy(state common.State, space common.Space) common.Policy {
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

	var actsMax []common.ActionEnum
	var valMax float64
	for _, act := range space.Acts() {
		val := accum.Mean(act)
		if (len(actsMax) == 0) || (val > valMax) {
			actsMax = []common.ActionEnum{act}
			valMax = val
		} else
		if val == valMax {
			actsMax = append(actsMax, act)
		}
	}

	var policy = common.NewPolicyPlus()
	for _, act := range actsMax {
		policy.Set(act, 1)
	}

	return policy
}
func (p *MCTS) Reward(state common.State, act common.ActionEnum, reward float64) {}

func NewMCTS(env envs.Env, mcNum int) Agent {
	var p = &MCTS{
		env:   env,
		model: common.NewTree(),
		mcNum: mcNum,
	}
	return p
}
