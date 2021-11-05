package agents

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEpsilonGreed0(t *testing.T) {
	Convey("TestEpsilonGreed0", t, func() {
		var (
			policy *common.Policy
			space  *common.Space
			state  common.Stater
			act    common.ActionEnum
			err    error
			reward float64
		)
		var agent = NewEpsilonGreed(0.5)
		// 构建一个名字叫“CartPole-v0”的Gym场景
		var env = envs.Envs(common.GameEnum_Bandits_v0)
		// 初始化场景
		state = env.Reset()
		space = env.ActionSpace()
		for count := 0; count < 10; count++ {
			// 画出当前场景情况
			//t.Logf("env:%v", env)
			// 给环境中Agent一次命令，并让环境演化一步
			policy = agent.Policy(space)
			act, err = policy.Sample()
			if err != nil {
				t.Fatal(err.Error())
			}
			state, reward, _ = env.Step(act)
			t.Logf("count:%v, act:%v, reward:%10.7f", count, common.Action2Str(act), reward)
		} // 关闭环境
		env.Close()
		for _, act = range space.Acts() {
			var ex = state.GetFloat64(fmt.Sprintf("%v", int(act)))
			t.Logf("act:%v, ex:%10.7f", act, ex)
		}
	})
}
