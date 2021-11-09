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
func (p *HashPolicy) Find(state State, mesh *Mesh) *NodePolicy {
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

type TreePolicy struct {
	*NodePolicy
	leaves map[int]*TreePolicy // 树结构
}

func NewTree() *TreePolicy {
	return &TreePolicy{
		NodePolicy: &NodePolicy{
			Accum: NewAccum(),
		},
		leaves: map[int]*TreePolicy{},
	}
}
func (p *TreePolicy) Find(state State, mesh *Mesh) *NodePolicy {
	var (
		node = p
		next *TreePolicy
		ok   bool
		code = state.Hash2(mesh)
	)
	for i, s := range code {
		next, ok = node.leaves[s]
		if !ok {
			if i == len(code)-1 {
				next = &TreePolicy{NodePolicy: NewNode(), leaves: map[int]*TreePolicy{}}
			} else {
				next = &TreePolicy{NodePolicy: NewNode(), leaves: map[int]*TreePolicy{}}
			}
			node.leaves[s] = next
		}
		node = next
	}
	return node.NodePolicy
}
func (p *TreePolicy) Clear() {
	p.NodePolicy = &NodePolicy{Accum: NewAccum()}
	p.leaves = map[int]*TreePolicy{}
}
