package common

type NodeValue struct {
	Value float64
	State State
}

func NewNodeValue(state State) *NodeValue {
	return &NodeValue{State: state, Value: 0}
}

type HashValue struct {
	nodes map[string]*NodeValue // 表结构
}

func NewHashValue() *HashValue {
	return &HashValue{
		nodes: map[string]*NodeValue{}, // 表结构
	}
}
func (p *HashValue) Find(state State, mesh *Mesh) *NodeValue {
	var node, ok = p.nodes[state.Hash(mesh)]
	if !ok {
		node = NewNodeValue(state)
		p.nodes[state.Hash(mesh)] = node
	}
	return node
}
func (p *HashValue) Clear() {
	p.nodes = map[string]*NodeValue{}
}
