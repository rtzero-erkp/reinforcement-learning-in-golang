package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math/rand"
)

var _ common.Env = &BanditsEnv{}

type BanditsEnv struct {
	// 常量状态
	banditsNum int
	// 当前状态
	state common.Info
	// 工具类
	info  common.Info
	space common.Space // 可选行动
}

func NewBanditsEnv(banditsNum int) common.Env {
	return &BanditsEnv{
		banditsNum: banditsNum,
		state:      common.NewInfoMap(),
		info:       common.NewInfoMap(),
		space:      common.NewSpaceVecByNum(banditsNum),
	}
}

func (p *BanditsEnv) Clone() common.Env {
	var cp = &BanditsEnv{
		banditsNum: p.banditsNum,
		state:      p.state.Clone(),
		info:       p.info.Clone(),
		space:      p.space.Clone(),
	}

	return cp
}
func (p *BanditsEnv) ActionSpace() common.Space {
	return p.space
}
func (p *BanditsEnv) String() string {
	return "Bandits"
}
func (p *BanditsEnv) Step(act common.ActionEnum) (res *common.Result) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("actions space not contain act:%v", act))
	}

	var key = fmt.Sprintf("ex%v", act)
	var val = p.info.Get(key).(float64)
	//var reward = p.rand.Float64() * val * 2
	var reward = []float64{val}
	return &common.Result{
		State:  p.state,
		Reward: reward,
		Done:   true,
		Info:   p.info,
	}
}
func (p *BanditsEnv) Reset() common.Info {
	for i := 0; i < p.banditsNum; i++ {
		var exI = rand.Float64()
		p.info.Set(fmt.Sprintf("ex%v", i), exI)
	}
	return p.state
}
func (p *BanditsEnv) Set(state common.Info) {
	p.state = state
}
