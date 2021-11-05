package agents

import (
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEpsilonGreed0(t *testing.T) {
	Convey("TestEpsilonGreed0", t, func() {
		var (
			policy common.Policy
			space  common.Space
			state  common.Stater
			act    common.ActionEnum
			reward float64
			rewardCum common.Reward
		)
		var env = envs.NewBanditsEnv(5)
		state = env.Reset()
		space = env.ActionSpace()
		rewardCum = common.NewReward1D(env.ActionSpace())
		var agent = NewEpsilonGreed(space, 0.5)
		for count := 0; count < 100; count++ {
			policy = agent.Policy(space)
			act = policy.Sample()
			state, reward, _ = env.Step(act)
			agent.Reward(act, reward)
			rewardCum.Add(act, reward)
		}
		env.Close()
		t.Log(state)
		t.Log(rewardCum)
	})
}
