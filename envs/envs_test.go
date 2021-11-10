package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBandits0(t *testing.T) {
	Convey("TestBandits0", t, func() {
		var (
			res    = &common.Result{}
			reward float64
			accum  common.Accumulate
		)
		var env = NewBanditsEnv(5)
		env.Reset()
		accum = common.NewAccum()
		for count := 0; count < 10; count++ {
			var act = env.ActionSpace().Sample()
			res = env.Step(act)
			accum.Add(act, reward)
		}
		t.Log(res.Info)
		t.Log(accum)
	})
}

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		var (
			res    = &common.Result{}
			reward float64
			accum  common.Accumulate
		)
		var env = NewCartPoleEnv(2.4, 12)
		env.Reset()
		accum = common.NewAccum()
		for !res.Done {
			var act = env.ActionSpace().Sample()
			res = env.Step(act)
			//t.Log(res.State)
			accum.Add(act, reward)
		}
		t.Log(accum)
	})
}
