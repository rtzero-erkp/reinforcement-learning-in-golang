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
		mesh       = common.NewMesh(1)

		env    = envs.NewBanditsEnv(banditsNum)
		agents = []Agent{
			NewMC(env, mcNum),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_Random),
		}
	)
	for _, agent := range agents {
		Convey(fmt.Sprintf("TestMc0:%v", agent), t, func() {
			t.Logf("TestMc0:%v", agent)
			state := env.Reset()
			space := env.ActionSpace()
			act := agent.Policy(state, space)
			res := env.Step(act)
			agent.Reward(state, act, res.Reward)
			env.Close()
			t.Log(act)
			t.Log(res.Info)
		})
	}
}

func TestMc1(t *testing.T) {
	var (
		mcNum = 100

		xRange     = 2.4
		thetaRange = 12.0
		simulate   = 3
		stepLimit  = 3000
		mesh       = common.NewMesh(50/xRange, 20, 50/thetaRange, 20)

		env    = envs.NewCartPoleEnv(xRange, thetaRange)
		agents = []Agent{
			NewMC(env, mcNum),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_Random),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_MeanQ),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_UCB),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_EpsilonGreed, 0.5),
			NewMCTS(env, mcNum, mesh, common.SearchMethodEnum_SoftMax, 0.5),
		}

		state common.State
		act   common.ActionEnum
		res   *envs.Result
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestMc1:%v", agent), t, func() {
			t.Logf("TestMc1:%v", agent)
			for i0 := 0; i0 < simulate; i0++ {
				// reset
				agent.Reset()
				state = env.Reset()
				res = &envs.Result{}
				var reward float64 = 0
				var step = 0
				// simulate
				for !res.Done {
					act = agent.Policy(state, env.ActionSpace())
					res = env.Step(act)
					reward += res.Reward
					state = res.State
					step += 1
					if step >= stepLimit {
						log.Printf("step:%v >= stepLimit:%v, done", stepLimit, stepLimit)
						break
					}
				}
				log.Printf("i:%v, reward:%v", i0, reward)
			}
			env.Close()
		})
	}
}
