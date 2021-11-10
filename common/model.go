package common

import "log"

type Node interface{}
type NodePolicy struct {
	Accum Accumulate
}
type NodeQ struct {
	Accum Accumulate
}
type NodeValue struct {
	Value float64
}
type NodeCfr struct{}

func NewNode(mode ModelType) Node {
	var node Node
	switch mode {
	case ModelTypeEnum_Value:
		node = &NodeValue{Value: 0}
	case ModelTypeEnum_Policy:
		node = &NodePolicy{Accum: NewAccum()}
	case ModelTypeEnum_Q:
		node = &NodeQ{Accum: NewAccum()}
	case ModelTypeEnum_Cfr:
		node = &NodeCfr{}
	default:
		log.Fatalf("unknown mode %v", mode)
	}
	return node
}

type ModelMap struct {
	mode  ModelType
	nodes map[string]Node
}

func NewModelMap(mode ModelType) *ModelMap {
	return &ModelMap{
		mode:  mode,
		nodes: map[string]Node{},
	}
}
func (p *ModelMap) Find(code string) Node {
	var node, ok = p.nodes[code]
	if !ok {
		node = NewNode(p.mode)
		p.nodes[code] = node
	}
	return node
}
func (p *ModelMap) Clear() {
	p.nodes = map[string]Node{}
}

type ModelTree struct {
	mode  ModelType
	node  Node
	child map[ActionEnum]*ModelTree
}

func NewModelTree(mode ModelType) *ModelTree {
	return &ModelTree{
		mode:  mode,
		node:  NewNode(mode),
		child: map[ActionEnum]*ModelTree{},
	}
}
func (p *ModelTree) Find(history ...ActionEnum) Node {
	var target = p
	for _, act := range history {
		var next, ok = target.child[act]
		if !ok {
			next = &ModelTree{
				mode:  p.mode,
				node:  NewNode(p.mode),
				child: map[ActionEnum]*ModelTree{},
			}
			target.child[act] = next
		}
		target = next
	}
	return target.node
}
func (p *ModelTree) Move(history ...ActionEnum) *ModelTree {
	var target = p
	for _, act := range history {
		var next, ok = target.child[act]
		if !ok {
			next = &ModelTree{
				mode:  p.mode,
				node:  NewNode(p.mode),
				child: map[ActionEnum]*ModelTree{},
			}
			target.child[act] = next
		}
		target = next
	}
	return target
}
func (p *ModelTree) Clear() {
	p.node = NewNode(p.mode)
	p.child = map[ActionEnum]*ModelTree{}
}
