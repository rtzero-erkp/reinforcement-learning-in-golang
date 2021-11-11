package common

import (
	"fmt"
	"log"
)

type UpdateParam struct {
	Model UpdateEnum
	Args  []interface{}
}

type SearchParam struct {
	Model SearchEnum
	Args  []interface{}
}

func NewUpdateParam(mode UpdateEnum, args ...interface{}) *UpdateParam {
	return &UpdateParam{
		Model: mode,
		Args:  args,
	}
}
func NewSearchParam(mode SearchEnum, args ...interface{}) *SearchParam {
	return &SearchParam{
		Model: mode,
		Args:  args,
	}
}

type ModelMap struct {
	model  ModelEnum
	update *UpdateParam
	nodes  map[string]interface{}
}

func NewModelMap(model ModelEnum, update *UpdateParam) *ModelMap {
	var o = &ModelMap{model: model, update: update, nodes: map[string]interface{}{}}
	return o
}

func (p *ModelMap) Clear() {
	p.nodes = map[string]interface{}{}
}
func (p *ModelMap) String() string {
	var line = fmt.Sprintf("[model] model:%v, update:%v\n", p.model, p.update)
	for k, v := range p.nodes {
		line += fmt.Sprintf("[model] %v:%v\n", k, v)
	}
	return line
}
func (p *ModelMap) Sample(env Env, encoder Encoder, search *SearchParam) (act ActionEnum) {
	switch p.model {
	case NodeEnum_Value:
		act = p.sampleV(env, encoder, search)
	case NodeEnum_Q:
		act = p.sampleQ(env, encoder, search)
	default:
		act = env.Space().Sample()
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

func (p *ModelMap) sampleV(env Env, encoder Encoder, search *SearchParam) (act ActionEnum) {
	accum := NewAccumulate()
	for _, actI := range env.Space().Acts() {
		envCrt := env.Clone()
		res := envCrt.Step(actI)
		var code = encoder.Hash(res.State)
		var reward, ok = p.nodes[code]
		if !ok {
			reward = 0.0
			p.nodes[code] = reward
		}
		accum.Add(actI, reward.(float64))
	}
	act = accum.Sample(env.Space(), search)
	return
}
func (p *ModelMap) sampleQ(env Env, encoder Encoder, search *SearchParam) (act ActionEnum) {
	var code = encoder.Hash(env.State())
	var accum, ok = p.nodes[code]
	if !ok {
		accum = NewAccumulate()
		p.nodes[code] = accum
	}
	act = accum.(Accumulate).Sample(env.Space(), search)
	return act
}
func (p *ModelMap) updateV(mem *MemoryCode) {
	var codeFrom = mem.From
	var codeTo = mem.To
	var valueFrom, ok0 = p.nodes[codeFrom]
	var valueTo, ok1 = p.nodes[codeTo]
	if !ok0 {
		valueFrom = 0.0
		p.nodes[codeFrom] = valueFrom
	}
	if !ok1 {
		valueTo = 0.0
		p.nodes[codeTo] = valueTo
	}

	switch p.update.Model {
	case UpdateEnum_DT:
		// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
		alpha := p.update.Args[0].(float64)
		lambda := p.update.Args[1].(float64)
		valueFrom = valueFrom.(float64) + alpha*(mem.Reward+lambda*(valueTo.(float64)-valueFrom.(float64)))
		p.nodes[codeFrom] = valueFrom
	case UpdateEnum_SARSA:
		log.Fatal("not impl")
	}
}
func (p *ModelMap) updateQ(mem *MemoryCode) {
	var code = mem.From
	var accum, ok = p.nodes[code]
	if !ok {
		accum = NewAccumulate()
		p.nodes[code] = accum
	}
	accum.(Accumulate).Add(mem.Act, mem.Reward)
}
