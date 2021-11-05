package agents

import (
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestSoftMax0(t *testing.T) {
	Convey("TestSoftMax0", t, func() {
		const (
			tau        = 0.5
			banditsNum = 5
			simulate   = 100
		)
		var (
			space  common.Space
			state  common.State
			info   common.Info
			reward float64
			accum  common.Accumulate
			mesh   = []float64{1}
			env    = envs.NewBanditsEnv(banditsNum)
		)

		state = env.Reset()
		space = env.ActionSpace()
		accum = common.NewAccum()
		var agent = NewSoftMax(tau)
		for count := 0; count < simulate; count++ {
			code:=state.Encode(mesh)
			policy := agent.Policy(code, space)
			act := policy.Sample()
			state, reward, _, info = env.Step(act)
			agent.Reward(code, act, reward)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(state)
		t.Log(info)
		t.Log(accum)
	})
}

func TestSoftMax1(t *testing.T) {
	Convey("TestSoftMax1", t, func() {
		const (
			xRange     = 2.4
			thetaRange = 12
			tau        = 0.5
		)
		var (
			mesh  = []float64{50 / xRange, 20, 50 / thetaRange, 20}
			env   = envs.NewCartPoleEnv(xRange, thetaRange)
			agent = NewSoftMax(tau)
		)

		done := false
		state := env.Reset()
		//log.Println(state)
		space := env.ActionSpace()
		for !done {
			code := state.Encode(mesh)
			act := agent.Policy(code, space).Sample()
			state, _, done, _ = env.Step(act)
			//log.Println(code)
			//log.Println(act)
			//log.Println(state)
		}
		env.Close()
	})
}

func TestSoftMax2(t *testing.T) {
	Convey("TestSoftMax2", t, func() {
		const (
			xRange     = 2.4
			thetaRange = 12
			tau        = 0.5
			simulate   = 100000
			split      = 50
		)
		var (
			mesh   = []float64{50 / xRange, 20, 50 / thetaRange, 20}
			mem    = common.NewMemory()
			env    = envs.NewCartPoleEnv(xRange, thetaRange)
			record = make([]float64, split)
			agent  = NewSoftMax(tau)
		)

		for i0 := 0; i0 < simulate; i0++ {
			// reset
			var reward float64
			done := false
			state := env.Reset()
			mem.Clear()
			// sim
			space := env.ActionSpace()
			for !done {
				code := state.Encode(mesh)
				act := agent.Policy(code, space).Sample()
				state, reward, done, _ = env.Step(act)
				mem.Add(code, act, reward)
			}
			// learn
			steps := mem.Get()
			stepsNum := len(steps)
			for i1, step := range steps {
				reward = -float64(i1) / float64(stepsNum)
				agent.Reward(step.Code, step.Act, reward)
			}
			// record
			record[i0/(simulate/split)] += float64(stepsNum)
		}
		//agent.Reward(code, act, reward)
		env.Close()

		for i, stepNum := range record {
			log.Printf("[record] idx:%3d, step:%v, mean:%.3f", i, stepNum, stepNum/simulate*split)
		}
	})
}
