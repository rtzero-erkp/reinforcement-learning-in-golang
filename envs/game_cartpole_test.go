package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		var (
			res    = &Result{}
			reward float64
			accum  common.Accumulate
		)
		var env = NewCartPoleEnv(2.4, 12)
		env.Reset()
		accum = common.NewAccum()
		for !res.Done {
			var act = env.ActionSpace().Sample()
			res = env.Step(act)
			t.Log(res.State)
			accum.Add(act, reward)
		}
		env.Close()
		t.Log(accum)
	})
}
