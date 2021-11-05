package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBandits0(t *testing.T) {
	Convey("TestBandits0", t, func() {
		var (
			state  common.Stater
			reward float64
			accum  common.Accumulate
		)
		var env = NewBanditsEnv(5)
		state = env.Reset()
		accum = common.NewReward1D(env.ActionSpace())
		for count := 0; count < 10; count++ {
			var act = env.ActionSpace().Sample()
			state, reward, _ = env.Step(act)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(state)
		t.Log(accum)
	})
}
