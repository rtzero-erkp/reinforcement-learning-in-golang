package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBandits0(t *testing.T) {
	Convey("TestBandits0", t, func() {
		var (
			info   common.Info
			reward float64
			accum  common.Accumulate
		)
		var env = NewBanditsEnv(5)
		env.Reset()
		accum = common.NewAccum()
		for count := 0; count < 10; count++ {
			var act = env.ActionSpace().Sample()
			_, reward, _, info = env.Step(act)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(info)
		t.Log(accum)
	})
}
