package agents

import (
	"gameServer/common"
	"gameServer/envs"
	"log"
	"math/rand"
)

var _ Agent = &DT{}

type DT struct {
	model  *common.ValueMap // 模型
	mesh   *common.Mesh
	env    envs.Env
	alpha  float64
	lambda float64
}

func (p *DT) Reset() {}
func (p *DT) String() string {
	return "DT"
}
func (p *DT) Policy(state common.State, space common.Space) common.ActionEnum {
	var valueMax float64
	var actsMax []common.ActionEnum
	for _, act := range space.Acts() {
		var cp = p.env.Clone()
		cp.Set(state)
		var stateCrt = cp.Step(act).State
		var nodeCrt = p.model.Find(stateCrt, p.mesh)
		if (len(actsMax) == 0) || (nodeCrt.Value > valueMax) {
			actsMax = []common.ActionEnum{act}
		} else
		if nodeCrt.Value == valueMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("space is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
func (p *DT) Reward(state common.State, act common.ActionEnum, reward float64) {
	var cp = p.env.Clone()
	cp.Set(state)
	var stateDot = cp.Step(act).State
	var node = p.model.Find(state, p.mesh)
	var nodeDot = p.model.Find(stateDot, p.mesh)
	var vs = node.Value
	var vsDot = nodeDot.Value
	// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
	vs = vs + p.alpha*(reward+p.lambda*(vsDot-vs))
	node.Value = vs
}

func NewDT(env envs.Env, alpha float64, lambda float64, mesh *common.Mesh) Agent {
	var p = &DT{
		alpha:  alpha,  // 0.1
		lambda: lambda, // 0.5
		env:    env,
		model:  common.NewValueMap(),
		mesh:   mesh,
	}
	return p
}
