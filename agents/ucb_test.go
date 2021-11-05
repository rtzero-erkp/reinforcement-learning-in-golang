package agents

import (
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUCB0(t *testing.T) {
	Convey("TestUCB0", t, func() {
		var (
			state  common.Stater
			reward float64
			accum  common.Accumulate
		)
		var env = envs.NewBanditsEnv(5)
		state = env.Reset()
		space := env.ActionSpace()
		accum = common.NewReward1D(env.ActionSpace())
		var agent = NewUCB(space, 0.5)
		for count := 0; count < 100; count++ {
			policy := agent.Policy(space)
			act := policy.Sample()
			state, reward, _ = env.Step(act)
			agent.Reward(act, reward)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(state)
		t.Log(accum)
	})
}
