package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBandits0(t *testing.T) {
	Convey("TestBandits0", t, func() {
		var (
			res    = &Result{}
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
		env.Close()
		t.Log(res.Info)
		t.Log(accum)
	})
}
