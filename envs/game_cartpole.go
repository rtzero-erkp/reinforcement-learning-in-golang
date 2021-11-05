package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math"
	"math/rand"
	"time"
)

var _ Env = &CartPoleEnv{}

type CartPoleEnv struct {
	// 常量参数
	gravity               float64 // 重力
	massCart              float64 // 头部质量
	massPole              float64 // 杆部质量
	totalMass             float64 // 质量
	length                float64 // 杆的半长
	poleMassLength        float64
	forceMag              float64 // 推力
	tau                   float64
	kinematicsIntegrator  string // 运动学积分器
	thetaThresholdRadians float64
	xThreshold            float64
	// 当前状态
	state           common.Stater
	stepsBeyondDone int // step计数
	// 工具类
	space *common.Space // 可选行动
	rand  *rand.Rand    // 随机数生成器
}

func NewCartPoleEnv() Env {
	var massCart = 1.0 // 头部质量
	var massPole = 0.1 // 杆部质量
	var length = 0.5   // 杆的半长

	return &CartPoleEnv{
		gravity:               9.8,                 // 重力
		massCart:              massCart,            // 头部质量
		massPole:              massPole,            // 杆部质量
		totalMass:             massPole + massCart, // 质量
		length:                0.5,                 // 杆的半长
		poleMassLength:        massPole * length,
		forceMag:              10.0,    // 推力
		tau:                   0.02,    // seconds between state updates
		kinematicsIntegrator:  "euler", // 运动学积分器
		thetaThresholdRadians: 12 * 2 * math.Pi / 360,
		xThreshold:            2.4,
		state:                 common.NewState(),
		stepsBeyondDone:       0,
		space: common.NewActions(
			common.ActionEnum_Left,
			common.ActionEnum_Right,
		),
		rand: rand.New(rand.NewSource(time.Now().Unix())),
	}

}

func (p *CartPoleEnv) ActionSpace() *common.Space {
	return p.space
}
func (p *CartPoleEnv) String() string {
	var x = p.state.GetFloat64("x")
	var xDot = p.state.GetFloat64("xDot")
	var theta = p.state.GetFloat64("theta")
	var thetaDot = p.state.GetFloat64("thetaDot")
	return fmt.Sprintf("x:%v, xDot:%v, theta:%v, thetaDot:%v", x, xDot, theta, thetaDot)
}
func (p *CartPoleEnv) Seed(seed int64) rand.Source {
	var source = rand.NewSource(seed)
	p.rand = rand.New(source)
	return source
}
func (p *CartPoleEnv) Step(act common.ActionEnum) (state common.Stater, reward float64, done bool) {
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

	var x = p.state.GetFloat64("x")
	var xDot = p.state.GetFloat64("xDot")
	var theta = p.state.GetFloat64("theta")
	var thetaDot = p.state.GetFloat64("thetaDot")

	var cosTheta = math.Cos(theta)
	var sinTheta = math.Sin(theta)

	// For the interested reader:
	// https://coneural.org/florian/papers/05_cart_pole.pdf
	var temp = (force + p.poleMassLength*math.Pow(thetaDot, 2)*sinTheta) / p.totalMass
	var thetaAcc = (p.gravity*sinTheta - cosTheta*temp) / (p.length * (4.0/3.0 - p.massPole*math.Pow(cosTheta, 2)/p.totalMass))
	var xAcc = temp - p.poleMassLength*thetaAcc*cosTheta/p.totalMass

	if p.kinematicsIntegrator == "euler" {
		x = x + p.tau*xDot
		xDot = xDot + p.tau*xAcc
		theta = theta + p.tau*thetaDot
		thetaDot = thetaDot + p.tau*thetaAcc
	} else { // semi-implicit euler
		xDot = xDot + p.tau*xAcc
		x = x + p.tau*xDot
		thetaDot = thetaDot + p.tau*thetaAcc
		theta = theta + p.tau*thetaDot
	}

	p.state.SetFloat64("x", x)
	p.state.SetFloat64("xDot", xDot)
	p.state.SetFloat64("theta", theta)
	p.state.SetFloat64("thetaDot", thetaDot)

	state = p.state
	done = (x < -p.xThreshold) ||
		(x > p.xThreshold) ||
		(theta < -p.thetaThresholdRadians) ||
		(theta > p.thetaThresholdRadians)

	if !done {
		reward = 1.0
	} else {
		p.stepsBeyondDone += 1
		reward = 0.0
	}

	return
}
func (p *CartPoleEnv) Reset() common.Stater {
	p.state.SetFloat64("x", p.rand.Float64()*0.1-0.05)
	p.state.SetFloat64("xDot", p.rand.Float64()*0.1-0.05)
	p.state.SetFloat64("theta", p.rand.Float64()*0.1-0.05)
	p.state.SetFloat64("thetaDot", p.rand.Float64()*0.1-0.05)
	p.stepsBeyondDone = 0
	return p.state
}
func (p *CartPoleEnv) Close() {
}
