package common

import (
	"fmt"
)

type ModelTreeNode struct {
	parent   *ModelTreeNode
	children map[ActionEnum]*ModelTreeNode

	reward float64
	count  float64
}

func NewRootNode() *ModelTreeNode {
	return &ModelTreeNode{
		children: map[ActionEnum]*ModelTreeNode{},
	}
}
func NewMidNode(parent *ModelTreeNode) *ModelTreeNode {
	return &ModelTreeNode{
		parent:   parent,
		children: map[ActionEnum]*ModelTreeNode{},
	}
}

func (p *ModelTreeNode) Find(path ...ActionEnum) (node *ModelTreeNode) {
	node = p
	for _, actI := range path {
		node = node.get(actI)
	}
	return node
}
func (p *ModelTreeNode) Sample(env Env, search *SearchParam) (act ActionEnum) {
	accum := NewAccumulate()
	for _, actI := range env.Space().Acts() {
		node := p.get(actI)
		mean := node.reward / node.count
		accum.Add(actI, mean)
	}
	act = accum.Sample(env.Space(), search)
	return act
}
func (p *ModelTreeNode) Update(reward float64) {
	node := p
	for {
		node.reward += reward
		node.count += 1
		node = node.parent
		if node == nil {
			break
		}
		rewardAccum := 0.0
		countAccum := 0.0
		for _, next := range node.children {
			rewardAccum += next.reward
			countAccum += next.count
		}
		if countAccum != 0 {
			reward = rewardAccum / countAccum
		}
	}
}
func (p *ModelTreeNode) get(act ActionEnum) *ModelTreeNode {
	var node, ok = p.children[act]
	if !ok {
		node = NewMidNode(p)
		p.children[act] = node
	}
	return node
}

type ModelTree struct {
	model  ModelEnum
	update *UpdateParam
	root   *ModelTreeNode
}

func NewModelTree(model ModelEnum, update *UpdateParam) *ModelTree {
	var o = &ModelTree{model: model, update: update, root: NewRootNode()}
	return o
}

func (p *ModelTree) Clear() {
	p.root = NewRootNode()
}
func (p *ModelTree) String() string {
	var line = fmt.Sprintf("[model] model:%v, update:%v, root:%p\n", p.model, p.update, p.root)
	return line
}
func (p *ModelTree) Find(path ...ActionEnum) (node *ModelTreeNode) {
	return p.root.Find(path...)
}
func (p *ModelTree) Move(path ...ActionEnum) {
	node := p.root
	for _, actI := range path {
		next := node.get(actI)
		node.parent = nil
		node.children = nil
		next.parent = nil
		node = next
	}
	p.root = node
}
