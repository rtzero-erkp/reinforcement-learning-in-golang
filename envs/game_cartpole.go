package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math"
	"math/rand"
)

var _ Env = &CartPoleEnv{}

type CartPoleEnv struct {
	// 常量参数
	gravity              float64 // 重力
	massCart             float64 // 头部质量
	massPole             float64 // 杆部质量
	totalMass            float64 // 质量
	length               float64 // 杆的半长
	poleMassLength       float64
	forceMag             float64 // 推力
	tau                  float64
	kinematicsIntegrator string // 运动学积分器
	thetaRange           float64
	xRange               float64
	// 当前状态
	state common.State
	// 工具类
	info  common.Info
	space common.Space // 可选行动
	rand  *rand.Rand   // 随机数生成器
}

func (p *CartPoleEnv) Clone() Env {
	var cp = &CartPoleEnv{
		gravity:              p.gravity,
		massCart:             p.massCart,
		massPole:             p.massPole,
		totalMass:            p.totalMass,
		length:               p.length,
		poleMassLength:       p.poleMassLength,
		forceMag:             p.forceMag,
		tau:                  p.tau,
		kinematicsIntegrator: p.kinematicsIntegrator,
		thetaRange:           p.thetaRange,
		xRange:               p.xRange,
		state:                p.state.Clone(),
		info:                 p.info.Clone(),
		space:                p.space.Clone(),
		rand:                 rand.New(rand.NewSource(rand.Int63())),
	}
	return cp
}

func NewCartPoleEnv(xRange float64, thetaRange float64) Env {
	var massCart = 1.0 // 头部质量
	var massPole = 0.1 // 杆部质量
	var length = 0.5   // 杆的半长

	return &CartPoleEnv{
		gravity:              9.8,                 // 重力
		massCart:             massCart,            // 头部质量
		massPole:             massPole,            // 杆部质量
		totalMass:            massPole + massCart, // 质量
		length:               0.5,                 // 杆的半长
		poleMassLength:       massPole * length,
		forceMag:             10.0,    // 推力
		tau:                  0.02,    // 更新的单位时间s
		kinematicsIntegrator: "euler", // 运动学积分器
		thetaRange:           thetaRange,
		xRange:               xRange,
		state:                make([]float64, 4),
		info:                 common.NewInfoMap(),
		space: common.NewSpace1DByEnum(
			common.ActionEnum_Left,
			common.ActionEnum_Right,
		),
		rand: rand.New(rand.NewSource(rand.Int63())),
	}

}

func (p *CartPoleEnv) ActionSpace() common.Space {
	return p.space
}
func (p *CartPoleEnv) String() string {
	var x = p.state[0]
	var xDot = p.state[1]
	var theta = p.state[2]
	var thetaDot = p.state[3]
	return fmt.Sprintf("x:%v, xDot:%v, theta:%v, thetaDot:%v", x, xDot, theta, thetaDot)
}
func (p *CartPoleEnv) Seed(seed int64) rand.Source {
	var source = rand.NewSource(seed)
	p.rand = rand.New(source)
	return source
}
func (p *CartPoleEnv) Step(act common.ActionEnum) (res *Result) {
	if !p.space.Contain(act) {
		log.Fatal(fmt.Sprintf("actions space not contain act:%v", act))
	}

	var force float64
	switch act {
	case common.ActionEnum_Left:
		force = p.forceMag
	case common.ActionEnum_Right:
		force = -p.forceMag
	}

	var x = p.state[0]
	var xDot = p.state[1]
	var theta = p.state[2]
	var thetaDot = p.state[3]

	theta = theta * 2 * math.Pi / 360

	var cosTheta = math.Cos(theta)
	var sinTheta = math.Sin(theta)

	var temp = (force + p.poleMassLength*thetaDot*thetaDot*sinTheta) / p.totalMass
	var thetaAcc = (p.gravity*sinTheta - cosTheta*temp) / (p.length * (4.0/3.0 - p.massPole*cosTheta*cosTheta/p.totalMass))
	var xAcc = temp - p.poleMassLength*thetaAcc*cosTheta/p.totalMass

	if p.kinematicsIntegrator == "euler" {
		x = x + p.tau*xDot
		xDot = xDot + p.tau*xAcc
		theta = theta + p.tau*thetaDot
		thetaDot = thetaDot + p.tau*thetaAcc
	} else {
		xDot = xDot + p.tau*xAcc
		x = x + p.tau*xDot
		thetaDot = thetaDot + p.tau*thetaAcc
		theta = theta + p.tau*thetaDot
	}
	theta = theta / 2 / math.Pi * 360

	p.state[0] = x
	p.state[1] = xDot
	p.state[2] = theta
	p.state[3] = thetaDot

	res = &Result{}
	res.State = p.state
	res.Done = (x < -p.xRange) ||
		(x > p.xRange) ||
		(theta < -p.thetaRange) ||
		(theta > p.thetaRange)

	if !res.Done {
		p.info.Add("step", 1)
		res.Reward = 1.0
	} else {
		res.Reward = 0.0
	}
	res.Info = p.info
	return res
}
func (p *CartPoleEnv) Reset() common.State {
	p.state = []float64{
		p.rand.Float64()*0.1 - 0.05,
		p.rand.Float64()*0.1 - 0.05,
		p.rand.Float64()*0.1 - 0.05,
		p.rand.Float64()*0.1 - 0.05,
	}
	p.info.Set("step", 0)
	return p.state
}
func (p *CartPoleEnv) Set(state common.State) {
	p.state = state
}
func (p *CartPoleEnv) Close() {
}
