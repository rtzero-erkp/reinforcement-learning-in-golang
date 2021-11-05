package envs

import (
	"fmt"
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBandits0(t *testing.T) {
	Convey("TestBandits0", t, func() {
		var (
			state     common.Stater
			reward    float64
			rewardMap = map[common.ActionEnum]float64{}
			countMap  = map[common.ActionEnum]float64{}
		)
		// 构建一个名字叫“CartPole-v0”的Gym场景
		var env = Envs(common.GameEnum_Bandits_v0)
		// 初始化场景
		state = env.Reset()
		var acts = env.ActionSpace().Acts()
		for _, act := range acts {
			rewardMap[act] = 0
			countMap[act] = 0
		}

		for count := 0; count < 10; count++ {
			// 画出当前场景情况
			//t.Logf("env:%v", env)
			// 给环境中Agent一次命令，并让环境演化一步
			var act = env.ActionSpace().Sample()
			state, reward, _ = env.Step(act)
			t.Logf("count:%v, act:%v, reward:%10.7f", count, common.Action2Str(act), reward)
			rewardMap[act] += reward
			countMap[act] += 1
		} // 关闭环境
		env.Close()
		for _, act := range acts {
			var ex = state.GetFloat64(fmt.Sprintf("%v", int(act)))
			t.Logf("act:%v, reward:%10.7f, count:%10.7f, mean:%10.7f, ex:%10.7f", act, rewardMap[act], countMap[act], rewardMap[act]/countMap[act], ex)
		}
	})
}
