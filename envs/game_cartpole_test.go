package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		var (
			state  common.Stater
			reward float64
			done   = false
			accum  common.Accumulate
		)
		var env = NewCartPoleEnv()
		env.Reset()
		accum = common.NewReward1D(env.ActionSpace())
		for !done {
			var act = env.ActionSpace().Sample()
			state, reward, done = env.Step(act)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(state)
		t.Log(accum)
	})
}
