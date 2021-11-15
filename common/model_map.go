package common

import (
	"fmt"
)

type ModelMap struct {
	model ModelEnum
	nodes map[Code]interface{}
}

func NewModelMap(model ModelEnum) *ModelMap {
	var o = &ModelMap{model: model, nodes: map[Code]interface{}{}}
	return o
}

func (p *ModelMap) Find(code Code) interface{} {
	node := p.nodes[code]
	return node
}
func (p *ModelMap) Exist(code Code) (ok bool) {
	_, ok = p.nodes[code]
	return ok
}
func (p *ModelMap) Reset() {
	p.nodes = map[Code]interface{}{}
}
func (p *ModelMap) String() string {
	var line = fmt.Sprintf("[model] model:%v", p.model)
	return line
}
func (p *ModelMap) Sample(env Env, encoder Encoder, search *SearchMethod) (act ActEnum) {
	switch p.model {
	case NodeEnum_Value:
		act = search.VMap(env, encoder, p)
	case NodeEnum_Q:
		act = search.QMap(env, encoder, p)
	default:
		act = env.Acts().Sample()
	}
	return act
}

func (p *ModelMap) getV(code Code) float64 {
	var reward, ok = p.nodes[code]
	if !ok {
		reward = 0.0
		p.nodes[code] = reward
	}
	return reward.(float64)
}
func (p *ModelMap) getQ(code Code) Accumulate {
	var accum, ok = p.nodes[code]
	if !ok {
		accum = NewAccumulate()
		p.nodes[code] = accum
	}
	return accum.(Accumulate)
}
