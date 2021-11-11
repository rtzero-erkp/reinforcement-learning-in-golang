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
func (p *BanditsEnv) Space() common.Space {
	return p.space
}
func (p *BanditsEnv) String() string {
	return "Bandits"
}
func (p *BanditsEnv) Step(act common.ActionEnum) (res *common.Result) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("space not contain act:%v", act))
	}
	var val = p.info.Get(fmt.Sprintf("ex%v", act)).(float64)
	var reward = []float64{val}
	return &common.Result{
		State:  p.state,
		Reward: reward,
		Done:   true,
		Info:   p.info,
	}
}
func (p *BanditsEnv) Reset() (common.Info, common.Info) {
	for act := 0; act < p.banditsNum; act++ {
		p.info.Set(fmt.Sprintf("ex%v", act), rand.Float64())
	}
	return p.state, p.info
}
func (p *BanditsEnv) Set(state common.Info) {
	p.state = state
}
func (p *BanditsEnv) State() common.Info {
	return p.state
}
