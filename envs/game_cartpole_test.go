package envs

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCartPole0(t *testing.T) {
	Convey("TestCartPole0", t, func() {
		// 构建一个名字叫“CartPole-v0”的Gym场景
		var env = Make(GameEnum_CartPole_v0)
		// 初始化场景
		env.reset()
		for i := 0; i < 200; i++ {
			// 画出当前场景情况
			t.Logf("env:%v", env)
			// 给环境中Agent一次命令，并让环境演化一步
			var act = env.ActionSpace().Sample()
			t.Logf("act:%v", Action2Str(act))
			env.Step(act)
		} // 关闭环境
		env.Close()
	})
}
