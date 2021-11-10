package common

type NodePolicy struct { //
	Accum Accumulate // 累计收益
}

func NewNode() *NodePolicy {
	return &NodePolicy{Accum: NewAccum()}
}

type HashPolicy struct {
	nodes map[string]*NodePolicy // 表结构
}

func NewHashPolicy() *HashPolicy {
	return &HashPolicy{
		nodes: map[string]*NodePolicy{}, // 表结构
	}
}
func (p *HashPolicy) Find(state Info, mesh *Mesh) *NodePolicy {
	var node, ok = p.nodes[state.Hash(mesh)]
	if !ok {
		node = NewNode()
		p.nodes[state.Hash(mesh)] = node
	}
	return node
}
func (p *HashPolicy) Find2(key string) *NodePolicy {
	var node, ok = p.nodes[key]
	if !ok {
		node = NewNode()
		p.nodes[key] = node
	}
	return node
}
func (p *HashPolicy) Clear() {
	p.nodes = map[string]*NodePolicy{}
}

type NodeValue struct {
	Value float64
	State Info
}

func NewNodeValue(state Info) *NodeValue {
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
func (p *HashValue) Find(state Info, mesh *Mesh) *NodeValue {
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
