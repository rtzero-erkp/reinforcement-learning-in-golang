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
	info common.Info
	acts common.Acts // 可选行动
}

func NewBanditsEnv(banditsNum int) common.Env {
	return &BanditsEnv{
		banditsNum: banditsNum,
		state:      common.NewInfoMap(),
		info:       common.NewInfoMap(),
		acts:       common.NewActsVecByNum(banditsNum),
	}
}

func (p *BanditsEnv) String() string {
	return "Bandits"
}
func (p *BanditsEnv) Clone() common.Env {
	var cp = &BanditsEnv{
		banditsNum: p.banditsNum,
		state:      p.state.Clone(),
		info:       p.info.Clone(),
		acts:       p.acts.Clone(),
	}

	return cp
}
func (p *BanditsEnv) Acts() common.Acts {
	return p.acts
}
func (p *BanditsEnv) State() common.Info {
	return p.state
}
func (p *BanditsEnv) Reset() (common.Info, common.Info) {
	p.info.Clear()
	for act := 0; act < p.banditsNum; act++ {
		p.info.Set(fmt.Sprintf("ex%v", act), rand.Float64())
	}
	return p.state, p.info
}
func (p *BanditsEnv) Step(act common.ActEnum) (res *common.Result) {
	if !p.acts.Contain(act) {
		log.Fatal(fmt.Sprintf("acts not contain act:%v", act))
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
