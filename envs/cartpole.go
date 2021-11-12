package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math"
	"math/rand"
)

var _ common.Env = &CartPoleEnv{}

type CartPoleEnv struct {
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
	state                common.Info
	info                 common.Info
	acts                 common.Acts
}

func NewCartPoleEnv(xRange float64, thetaRange float64) common.Env {
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
		state:                common.NewInfoMap(),
		info:                 common.NewInfoMap(),
		acts: common.NewActsVecByEnum(
			common.ActionEnum_Left,
			common.ActionEnum_Right,
		),
	}

}

func (p *CartPoleEnv) String() string {
	return "CartPole"
}
func (p *CartPoleEnv) Clone() common.Env {
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
		acts:                 p.acts.Clone(),
	}
	return cp
}
func (p *CartPoleEnv) Acts() common.Acts {
	return p.acts
}
func (p *CartPoleEnv) State() common.Info {
	return p.state
}
func (p *CartPoleEnv) Reset() (common.Info, common.Info) {
	p.state.Set("x", rand.Float64()*0.1-0.05)
	p.state.Set("xDot", rand.Float64()*0.1-0.05)
	p.state.Set("theta", rand.Float64()*0.1-0.05)
	p.state.Set("thetaDot", rand.Float64()*0.1-0.05)
	p.info.Set("step", 0)
	return p.state, p.info
}
func (p *CartPoleEnv) Step(act common.ActEnum) (res *common.Result) {
	if !p.acts.Contain(act) {
		log.Fatal(fmt.Sprintf("actions acts not contain act:%v", act))
	}

	var force float64
	switch act {
	case common.ActionEnum_Left:
		force = p.forceMag
	case common.ActionEnum_Right:
		force = -p.forceMag
	}

	var x = p.state.Get("x").(float64)
	var xDot = p.state.Get("xDot").(float64)
	var theta = p.state.Get("theta").(float64)
	var thetaDot = p.state.Get("thetaDot").(float64)

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

	p.state.Set("x", x)
	p.state.Set("xDot", xDot)
	p.state.Set("theta", theta)
	p.state.Set("thetaDot", thetaDot)

	res = &common.Result{}
	res.State = p.state
	res.Done = (x < -p.xRange) ||
		(x > p.xRange) ||
		(theta < -p.thetaRange) ||
		(theta > p.thetaRange)

	if !res.Done {
		var step = p.info.Get("step").(int)
		p.info.Set("step", step+1)
		res.Reward = []float64{1.0}
	} else {
		res.Reward = []float64{0.0}
	}
	res.Info = p.info
	return res
}
