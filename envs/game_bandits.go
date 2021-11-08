package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math/rand"
)

var _ Env = &BanditsEnv{}

type BanditsEnv struct {
	// 常量状态
	banditsNum int
	// 当前状态
	state common.State
	// 工具类
	info  common.Info
	space common.Space // 可选行动
	rand  *rand.Rand   // 随机数生成器
}

func NewBanditsEnv(banditsNum int) Env {
	return &BanditsEnv{
		banditsNum: banditsNum,
		state:      []float64{0},
		info:       common.NewInfoMap(),
		space:      common.NewSpace1DByNum(banditsNum),
		rand:       rand.New(rand.NewSource(rand.Int63())),
	}
}

func (p *BanditsEnv) Clone() Env {
	var cp = &BanditsEnv{
		banditsNum: p.banditsNum,
		state:      p.state.Clone(),
		info:       p.info.Clone(),
		space:      p.space.Clone(),
		rand:       rand.New(rand.NewSource(rand.Int63())),
	}

	return cp
}
func (p *BanditsEnv) ActionSpace() common.Space {
	return p.space
}
func (p *BanditsEnv) String() string {
	var line = ""
	for i := 0; i < p.banditsNum; i++ {
		var key = fmt.Sprintf("ex%v", i)
		var val = p.info.Get(key)
		line += fmt.Sprintf("%v:%v ", key, val)
	}
	return line
}
func (p *BanditsEnv) Seed(seed int64) rand.Source {
	var source = rand.NewSource(seed)
	p.rand = rand.New(source)
	return source
}
func (p *BanditsEnv) Step(act common.ActionEnum) (res *Result) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("actions space not contain act:%v", act))
	}

	var key = fmt.Sprintf("ex%v", int(act))
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
func (p *BanditsEnv) Reset() common.State {
	for i := 0; i < p.banditsNum; i++ {
		var exI = p.rand.Float64()
		p.info.Set(fmt.Sprintf("ex%v", i), exI)
	}
	return p.state
}
func (p *BanditsEnv) Set(state common.State) {
	p.state = state
}
func (p *BanditsEnv) Close() {
}
