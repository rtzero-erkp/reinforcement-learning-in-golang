package envs

import (
	"gameServer/common"
	"math/rand"
)

type Result struct {
	State  common.State
	Reward float64
	Done   bool
	Info   common.Info
}

type Env interface {
	String() string                           // 打印
	ActionSpace() common.Space                // 行动空间
	Seed(seed int64) rand.Source              // 设置随机种子
	Step(act common.ActionEnum) (res *Result) // 执行一步
	Reset() common.State                      // 重置游戏
	Close()                                   // 关闭游戏
	Set(state common.State)                   // 设置为目标状态
}
