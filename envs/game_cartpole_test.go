package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		var (
			state  common.State
			reward float64
			done   = false
			accum  common.Accumulate
		)
		var env = NewCartPoleEnv(2.4, 12)
		env.Reset()
		accum = common.NewAccum()
		for !done {
			var act = env.ActionSpace().Sample()
			state, reward, done, _ = env.Step(act)
			t.Log(state)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(accum)
	})
}
