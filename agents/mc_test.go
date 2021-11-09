package agents

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestMc0(t *testing.T) {
	var (
		mcNum = 100

		banditsNum = 5
		env        = envs.NewBanditsEnv(banditsNum)
		agent      = NewMC(env, mcNum)
	)

	Convey(fmt.Sprintf("TestMc0:%v", agent), t, func() {
		t.Logf("TestMc0:%v", agent)
		state := env.Reset()
		space := env.ActionSpace()
		policy := agent.Policy(state, space)
		log.Println(policy)
		act := policy.Sample()
		res := env.Step(act)
		agent.Reward(state, act, res.Reward)
		env.Close()
		t.Log(state)
		t.Log(res.Info)
	})
}

func TestMc1(t *testing.T) {
	var (
		mcNum = 100

		xRange     = 2.4
		thetaRange = 12.0
		simulate   = 10

		env   = envs.NewCartPoleEnv(xRange, thetaRange)
		agent = NewMC(env, mcNum)

		state  common.State
		policy common.Policy
		act    common.ActionEnum
		res    *envs.Result
	)

	Convey(fmt.Sprintf("TestMc1:%v", agent), t, func() {
		t.Logf("TestMc1:%v", agent)
		for i0 := 0; i0 < simulate; i0++ {
			// reset
			state = env.Reset()
			res = &envs.Result{}
			var reward float64 = 0
			// simulate
			for !res.Done {
				policy = agent.Policy(state, env.ActionSpace())
				act = policy.Sample()
				res = env.Step(act)
				reward += res.Reward
				state = res.State
			}
			log.Printf("i:%v, reward:%v", i0, reward)
		}
		env.Close()
	})
}
