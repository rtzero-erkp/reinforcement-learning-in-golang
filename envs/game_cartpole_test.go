package envs

import (
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		var (
			reward float64
			count  = 0
			done   = false
		)
		// 构建一个名字叫“CartPole-v0”的Gym场景
		var env = Envs(common.GameEnum_CartPole_v0)
		// 初始化场景
		env.Reset()
		for !done {
			count += 1
			// 画出当前场景情况
			//t.Logf("env:%v", env)
			// 给环境中Agent一次命令，并让环境演化一步
			var act = env.ActionSpace().Sample()
			_, reward, done = env.Step(act)
			t.Logf("count:%v, act:%5v, reward:%10.7f", count, common.Action2Str(act), reward)
		} // 关闭环境
		env.Close()
		t.Logf("count:%v", count)
	})
}
