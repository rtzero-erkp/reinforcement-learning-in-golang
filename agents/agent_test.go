package agents

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestAgent0(t *testing.T) {
	var (
		tau     = 0.5
		epsilon = 0.5
		mcNum   = 100

		banditsNum = 5
		simulate   = 100

		mesh = common.NewMesh(1)
		env  = envs.NewBanditsEnv(banditsNum)

		agents = []Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
			NewMC(env, mcNum),
		}

		space common.Space
		state common.State
		accum common.Accumulate
		res   = &envs.Result{}
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent0:%v", agent), t, func() {
			t.Logf("TestAgent0:%v", agent)
			state = env.Reset()
			space = env.ActionSpace()
			accum = common.NewAccum()
			for count := 0; count < simulate; count++ {
				act := agent.Policy(state, space)
				res = env.Step(act)
				agent.Reward(state, act, res.Reward)
				accum.Add(act, res.Reward)
				state = res.State
			}
			env.Close()
			t.Log(res.Info)
			t.Log(accum)
		})
	}
}

func TestAgent1(t *testing.T) {
	var (
		tau     = 0.5
		epsilon = 0.5
		mcNum   = 100

		xRange     = 2.4
		thetaRange = 12.0

		mesh = common.NewMesh(50/xRange, 20, 50/thetaRange, 20)
		env  = envs.NewCartPoleEnv(xRange, thetaRange)

		agents = []Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
			NewMC(env, mcNum),
		}

		space common.Space
		state common.State
		res   = &envs.Result{}
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent1:%v", agent), t, func() {
			t.Logf("TestAgent1:%v", agent)
			state = env.Reset()
			space = env.ActionSpace()
			for !res.Done {
				act := agent.Policy(state, space)
				res = env.Step(act)
				state = res.State
			}
			env.Close()
		})
	}
}

func TestAgent2(t *testing.T) {
	var (
		tau     = 0.5
		epsilon = 0.5

		xRange     = 2.4
		thetaRange = 12.0
		simulate   = 1000
		split      = 10

		mesh = common.NewMesh(50/xRange, 20, 50/thetaRange, 20)
		env  = envs.NewCartPoleEnv(xRange, thetaRange)

		agents = []Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
		}

		state  common.State
		act    common.ActionEnum
		res    *envs.Result
		mem    = common.NewMem()
		space  = env.ActionSpace()
		record []float64
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent1:%v", agent), t, func() {
			t.Logf("TestAgent1:%v", agent)
			record = make([]float64, split)
			for i0 := 0; i0 < simulate; i0++ {
				// reset
				state = env.Reset()
				res = &envs.Result{}
				mem.Clear()
				// simulate
				for !res.Done {
					act = agent.Policy(state, space)
					res = env.Step(act)
					mem.Add(state, act, res.Reward)
					state = res.State
				}
				// train
				steps := mem.Get()
				stepsNum := len(steps)
				for i1, step := range steps {
					reward := - float64(i1) / float64(stepsNum)
					agent.Reward(step.State, step.Act, reward)
				}
				// record
				record[i0/(simulate/split)] += float64(stepsNum)
			}
			env.Close()

			for i, stepNum := range record {
				log.Printf("[record] idx:%3d, step:%v, mean:%.3f", i, stepNum, stepNum/float64(simulate/split))
			}
		})
	}
}
