package common

type ModelPolicy interface {
	Find(state Info, mesh Encoder) *NodePolicy
	Find2(key string) *NodePolicy
	Clear()
}
type ModelValue interface {
	Find(state Info, mesh Encoder) *NodeValue
	Clear()
}
type NodePolicy struct { //
	Accum Accumulate // 累计收益
}
type NodeValue struct {
	Value float64
	State Info
}

func NewNodePolicy() *NodePolicy {
	return &NodePolicy{Accum: NewAccum()}
}
func NewNodeValue(state Info) *NodeValue {
	return &NodeValue{State: state, Value: 0}
}

type hashPolicy struct {
	nodes map[string]*NodePolicy // 表结构
}

func NewHashPolicy() *hashPolicy {
	return &hashPolicy{
		nodes: map[string]*NodePolicy{}, // 表结构
	}
}
func (p *hashPolicy) Find(state Info, encoder Encoder) *NodePolicy {
	var code = encoder.Hash(state)
	var node, ok = p.nodes[code]
	if !ok {
		node = NewNodePolicy()
		p.nodes[code] = node
	}
	return node
}
func (p *hashPolicy) Find2(key string) *NodePolicy {
	var node, ok = p.nodes[key]
	if !ok {
		node = NewNodePolicy()
		p.nodes[key] = node
	}
	return node
}
func (p *hashPolicy) Clear() {
	p.nodes = map[string]*NodePolicy{}
}

type hashValue struct {
	nodes map[string]*NodeValue // 表结构
}

func NewHashValue() *hashValue {
	return &hashValue{
		nodes: map[string]*NodeValue{}, // 表结构
	}
}
func (p *hashValue) Find(state Info, encoder Encoder) *NodeValue {
	var code = encoder.Hash(state)
	var node, ok = p.nodes[code]
	if !ok {
		node = NewNodeValue(state)
		p.nodes[code] = node
	}
	return node
}
func (p *hashValue) Clear() {
	p.nodes = map[string]*NodeValue{}
}
