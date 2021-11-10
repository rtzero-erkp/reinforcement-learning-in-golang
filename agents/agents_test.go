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

		mesh = common.NewMesh([]string{}, []float64{})
		env  = envs.NewBanditsEnv(banditsNum)

		agents = []common.Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
			NewMC(env, mcNum),
			NewDT(env, 0.1, 0.5, mesh, common.SearchMethodEnum_UCB),
		}

		space common.Space
		state common.Info
		accum common.Accumulate
		res   = &common.Result{}
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent0:%v", agent), t, func() {
			t.Logf("TestAgent0:%v", agent)
			agent.Reset()
			state = env.Reset()
			space = env.ActionSpace()
			accum = common.NewAccum()
			for count := 0; count < simulate; count++ {
				act := agent.Policy(state, space)
				res = env.Step(act)
				agent.Reward(state, act, res.Reward[0])
				accum.Add(act, res.Reward[0])
				state = res.Info
			}
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

		mesh = common.NewMesh(
			[]string{"x", "xDot", "theta", "thetaDot"},
			[]float64{50 / xRange, 20, 50 / thetaRange, 20},
		)
		env = envs.NewCartPoleEnv(xRange, thetaRange)

		agents = []common.Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
			NewMC(env, mcNum),
			NewDT(env, 0.1, 0.5, mesh, common.SearchMethodEnum_UCB),
		}

		space common.Space
		state common.Info
		res   = &common.Result{}
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent1:%v", agent), t, func() {
			t.Logf("TestAgent1:%v", agent)
			agent.Reset()
			state = env.Reset()
			space = env.ActionSpace()
			for !res.Done {
				act := agent.Policy(state, space)
				res = env.Step(act)
				state = res.Info
			}
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

		mesh = common.NewMesh(
			[]string{"x", "xDot", "theta", "thetaDot"},
			[]float64{50 / xRange, 20, 50 / thetaRange, 20},
		)
		env = envs.NewCartPoleEnv(xRange, thetaRange)

		agents = []common.Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
			NewDT(env, 0.1, 0.5, mesh, common.SearchMethodEnum_UCB),
		}

		state  common.Info
		act    common.ActionEnum
		res    *common.Result
		mem    = common.NewMem()
		space  = env.ActionSpace()
		record []float64
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent1:%v", agent), t, func() {
			t.Logf("TestAgent1:%v", agent)
			agent.Reset()
			record = make([]float64, split)
			for i0 := 0; i0 < simulate; i0++ {
				// reset
				state = env.Reset()
				res = &common.Result{}
				mem.Clear()
				// simulate
				for !res.Done {
					act = agent.Policy(state, space)
					res = env.Step(act)
					mem.Add(state, act, res.Reward[0])
					state = res.Info
				}
				// train
				steps := mem.Get()
				stepsNum := len(steps)
				for i1, step := range steps {
					reward := - float64(i1) / float64(stepsNum)
					agent.Reward(step.Info, step.Act, reward)
				}
				// record
				record[i0/(simulate/split)] += float64(stepsNum)
			}

			for i, stepNum := range record {
				log.Printf("[record] idx:%3d, step:%v, mean:%.3f", i, stepNum, stepNum/float64(simulate/split))
			}
		})
	}
}

func TestAgent3(t *testing.T) {
	var (
		mcNum      = 100
		banditsNum = 5

		env    = envs.NewBanditsEnv(banditsNum)
		agents = []common.Agent{
			NewMC(env, mcNum),
			NewMCTS(env, mcNum, common.SearchMethodEnum_Random),
			NewMCTS(env, mcNum, common.SearchMethodEnum_MeanQ),
			NewMCTS(env, mcNum, common.SearchMethodEnum_UCB),
			NewMCTS(env, mcNum, common.SearchMethodEnum_EpsilonGreed, 0.5),
			NewMCTS(env, mcNum, common.SearchMethodEnum_SoftMax, 0.5),
		}
	)
	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent3:%v", agent), t, func() {
			t.Logf("TestAgent3:%v", agent)
			agent.Reset()
			state := env.Reset()
			space := env.ActionSpace()
			act := agent.Policy(state, space)
			res := env.Step(act)
			agent.Reward(state, act, res.Reward[0])
			t.Log(act)
			t.Log(res.Info)
		})
	}
}

func TestAgent4(t *testing.T) {
	var (
		mcNum      = 100
		xRange     = 2.4
		thetaRange = 12.0
		simulate   = 3
		stepLimit  = 3000

		env    = envs.NewCartPoleEnv(xRange, thetaRange)
		agents = []common.Agent{
			NewMC(env, mcNum),
			NewMCTS(env, mcNum, common.SearchMethodEnum_Random),
			NewMCTS(env, mcNum, common.SearchMethodEnum_MeanQ),
			NewMCTS(env, mcNum, common.SearchMethodEnum_UCB),
			NewMCTS(env, mcNum, common.SearchMethodEnum_EpsilonGreed, 0.5),
			NewMCTS(env, mcNum, common.SearchMethodEnum_SoftMax, 0.5),
		}

		state common.Info
		act   common.ActionEnum
		res   *common.Result
	)

	for _, agent := range agents {
		Convey(fmt.Sprintf("TestAgent4:%v", agent), t, func() {
			t.Logf("TestAgent4:%v", agent)
			for i0 := 0; i0 < simulate; i0++ {
				// reset
				agent.Reset()
				state = env.Reset()
				res = &common.Result{}
				var reward float64 = 0
				var step = 0
				// simulate
				for !res.Done {
					act = agent.Policy(state, env.ActionSpace())
					res = env.Step(act)
					reward += res.Reward[0]
					state = res.Info
					step += 1
					if step >= stepLimit {
						log.Printf("step:%v >= stepLimit:%v, done", stepLimit, stepLimit)
						break
					}
				}
				log.Printf("i:%v, reward:%v", i0, reward)
			}
		})
	}
}
