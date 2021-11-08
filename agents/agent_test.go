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

		mesh = []float64{1}
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
				policy := agent.Policy(state, space)
				act := policy.Sample()
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

		mesh = []float64{50 / xRange, 20, 50 / thetaRange, 20}
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
				act := agent.Policy(state, space).Sample()
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
		simulate   = 10000
		split      = 50

		mesh = []float64{100 / xRange, 100, 100 / thetaRange, 100}
		env  = envs.NewCartPoleEnv(xRange, thetaRange)

		agents = []Agent{
			NewEpsilonGreed(epsilon, mesh),
			NewSoftMax(tau, mesh),
			NewUCB(mesh),
		}

		state  common.State
		policy common.Policy
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
					policy = agent.Policy(state, space)
					act = policy.Sample()
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

func TestMc0(t *testing.T) {
	var (
		mcNum = 10000

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
