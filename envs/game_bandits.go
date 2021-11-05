package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math/rand"
	"time"
)

var _ Env = &BanditsEnv{}

type BanditsEnv struct {
	// 常量状态
	banditsNum int
	// 当前状态
	state           common.Stater
	stepsBeyondDone int // step计数
	// 工具类
	space *common.Space // 可选行动
	rand  *rand.Rand    // 随机数生成器
}

func NewBanditsEnv() Env {
	var banditsNum = 5
	return &BanditsEnv{
		banditsNum:      banditsNum,
		state:           common.NewState(),
		stepsBeyondDone: 0,
		space:           common.NewActionsByNum(banditsNum),
		rand:            rand.New(rand.NewSource(time.Now().Unix())),
	}
}
func NewBanditsEnvByNum(banditsNum int) Env {
	return &BanditsEnv{
		banditsNum:      banditsNum,
		state:           common.NewState(),
		stepsBeyondDone: 0,
		space:           common.NewActionsByNum(banditsNum),
		rand:            rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (p *BanditsEnv) ActionSpace() *common.Space {
	return p.space
}
func (p *BanditsEnv) String() string {
	var line = ""
	for i := 0; i < p.banditsNum; i++ {
		var key = fmt.Sprintf("ex%v", i)
		var val = p.state.GetFloat64(key)
		line += fmt.Sprintf("%v:%v ", key, val)
	}
	return line
}
func (p *BanditsEnv) Seed(seed int64) rand.Source {
	var source = rand.NewSource(seed)
	p.rand = rand.New(source)
	return source
}
func (p *BanditsEnv) Step(act common.ActionEnum) (state common.Stater, reward float64, done bool) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("actions space not contain act:%v", act))
	}

	var key = fmt.Sprintf("%v", int(act))
	var val = p.state.GetFloat64(key)

	state = p.state
	reward = p.rand.Float64() * val * 2
	done = false
	return
}
func (p *BanditsEnv) Reset() common.Stater {
	for i := 0; i < p.banditsNum; i++ {
		var exI = p.rand.Float64()
		p.state.SetFloat64(fmt.Sprintf("%v", i), exI)
	}
	p.stepsBeyondDone = 0
	return p.state
}
func (p *BanditsEnv) Close() {
}
