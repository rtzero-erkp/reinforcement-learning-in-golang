package agents

import (
	"gameServer/common"
	"math"
)

var _ Agent = &UCB{}

type UCB struct {
	model common.Model // 模型
	mesh  []float64
}

func (p *UCB) String() string {
	return "UCB"
}
func (p *UCB) Policy(state common.State, space common.Space) common.Policy {
	var node = p.model.Find(state.Encode(p.mesh))
	//log.Printf("[accum] state:%v, Accum:%v", state.Encode(p.mesh), node.Accum())
	countSum := node.Accum().Count()
	if countSum == 0 {
		node.Policy().Clean()
		node.Policy().Set(space.Sample(), 1)
		return node.Policy()
	}
	var actsMax []common.ActionEnum
	var ucbMax float64
	for _, act := range space.Acts() {
		count := node.Accum().CountAct(act)
		if count == 0 {
			node.Policy().Clean()
			node.Policy().Set(act, 1)
			return node.Policy()
		}
		upperBound := math.Sqrt((2 * math.Log(countSum)) / count)
		q := node.Accum().Mean(act)
		ucb := upperBound + q
		if (len(actsMax) == 0) || ucb > ucbMax {
			ucbMax = ucb
			actsMax = []common.ActionEnum{act}
		} else
		if ucb == ucbMax {
			actsMax = append(actsMax, act)
		}
	}
	node.Policy().Clean()
	for _, act := range actsMax {
		node.Policy().Set(act, 1)
	}
	return node.Policy()
}
func (p *UCB) Reward(state common.State, act common.ActionEnum, reward float64) {
	var node = p.model.Find(state.Encode(p.mesh))
	node.Accum().Add(act, reward)
}

func NewUCB(mesh []float64) Agent {
	var p = &UCB{
		model: common.NewRootNode(),
		mesh:  mesh,
	}
	return p
}
