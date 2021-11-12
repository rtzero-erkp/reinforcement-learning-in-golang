package common

import (
	"fmt"
	"log"
)

type ModelMap struct {
	model  ModelEnum
	update *UpdateParam
	nodes  map[Code]interface{}
}

func NewModelMap(model ModelEnum, update *UpdateParam) *ModelMap {
	var o = &ModelMap{model: model, update: update, nodes: map[Code]interface{}{}}
	return o
}

func (p *ModelMap) Reset() {
	p.nodes = map[Code]interface{}{}
}
func (p *ModelMap) String() string {
	var line = fmt.Sprintf("[model] model:%v, update:%v\n", p.model, p.update)
	return line
}
func (p *ModelMap) Sample(env Env, encoder Encoder, search *SearchParam) (act ActEnum) {
	switch p.model {
	case NodeEnum_Value:
		act = p.sampleV(env, encoder, search)
	case NodeEnum_Q:
		act = p.sampleQ(env, encoder, search)
	default:
		act = env.Acts().Sample()
	}
	return act
}
func (p *ModelMap) Update(mem *MemoryCode) {
	switch p.model {
	case NodeEnum_Value:
		p.updateV(mem)
	case NodeEnum_Q:
		p.updateQ(mem)
	}
}

func (p *ModelMap) sampleV(env Env, encoder Encoder, search *SearchParam) (act ActEnum) {
	accum := NewAccumulate()
	for _, actI := range env.Acts().All() {
		envCrt := env.Clone()
		res := envCrt.Step(actI)
		code := encoder.Hash(res.State)
		reward := p.getV(code)
		accum.Add(actI, reward)
	}
	act = accum.Sample(env.Acts(), search)
	return
}
func (p *ModelMap) sampleQ(env Env, encoder Encoder, search *SearchParam) (act ActEnum) {
	//log.Printf("[model] acts:%v", env.Acts())
	//log.Printf("[model] state:%v", env.State())
	//log.Printf("[model] encoder:%v", encoder)
	code := encoder.Hash(env.State())
	//log.Printf("[model] code:%v", code)
	accum := p.getQ(code)
	//log.Printf("[model] accum:%v", accum)
	act = accum.Sample(env.Acts(), search)
	//log.Printf("[model] act:%v", act)
	return act
}
func (p *ModelMap) updateV(mem *MemoryCode) {
	codeFrom := mem.From
	codeTo := mem.To
	valueFrom := p.getV(codeFrom)
	valueTo := p.getV(codeTo)
	switch p.update.Model {
	case UpdateEnum_DT:
		// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
		alpha := p.update.Args[0].(float64)
		lambda := p.update.Args[1].(float64)
		valueFrom = valueFrom + alpha*(mem.Reward+lambda*(valueTo-valueFrom))
		p.nodes[codeFrom] = valueFrom
	case UpdateEnum_SARSA:
		log.Fatal("not impl")
	}
}
func (p *ModelMap) updateQ(mem *MemoryCode) {
	code := mem.From
	accum := p.getQ(code)
	accum.Add(mem.Act, mem.Reward)
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
