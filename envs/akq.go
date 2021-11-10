package envs

import (
	"fmt"
	"gameServer/common"
	"log"
)

type Result struct {
	State  common.Info
	Reward float64
	Done   bool
	Info   common.Info
}
type AKQEnv struct {
	state common.Info  // 玩家状态
	info  common.Info  // 游戏信息
	space common.Space // 行动空间
}

func NewAKQEnv() *AKQEnv {
	return &AKQEnv{
		state: common.NewInfoMap(),
		info:  common.NewInfoMap(),
		space: common.NewSpaceVecByEnum(common.ActionEnum_CardA, common.ActionEnum_CardK, common.ActionEnum_CardQ),
	}
}

func (p *AKQEnv) Clone() *AKQEnv {
	var cp = &AKQEnv{
		state: common.NewInfoMap(),
		info:  common.NewInfoMap(),
		space: common.NewSpaceVecByEnum(common.ActionEnum_CardA, common.ActionEnum_CardK, common.ActionEnum_CardQ),
	}

	return cp
}
func (p *AKQEnv) ActionSpace() common.Space {
	return p.space
}
func (p *AKQEnv) String() string {
	return fmt.Sprintf("[AKQ] %v", p.space)
}
func (p *AKQEnv) Step(act common.ActionEnum) (res *Result) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("actions space not contain act:%v", act))
	}

	var key = fmt.Sprintf("ex%v", act)
	var val = p.info.Get(key)
	//var reward = p.rand.Float64() * val * 2
	var reward = val
	return &Result{
		State:  p.state,
		Reward: reward,
		Done:   true,
		Info:   p.info,
	}
}
func (p *AKQEnv) Reset() common.Info {
	return p.state
}
