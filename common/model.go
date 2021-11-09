package common

type Node struct {
	Accum Accumulate // 累计收益
}

func NewNode() *Node {
	return &Node{Accum: NewAccum()}
}

type HashMap struct {
	nodes map[string]*Node // 表结构
}

func NewHashMap() *HashMap {
	return &HashMap{
		nodes: map[string]*Node{}, // 表结构
	}
}
func (p *HashMap) Find(state State, mesh *Mesh) *Node {
	var node, ok = p.nodes[state.Hash(mesh)]
	if !ok {
		node = NewNode()
		p.nodes[state.Hash(mesh)] = node
	}
	return node
}
func (p *HashMap) Find2(key string) *Node {
	var node, ok = p.nodes[key]
	if !ok {
		node = NewNode()
		p.nodes[key] = node
	}
	return node
}
func (p *HashMap) Clear() {
	p.nodes = map[string]*Node{}
}

type Tree struct {
	*Node
	leaves map[int]*Tree // 树结构
}

func NewTree() *Tree {
	return &Tree{
		Node: &Node{
			Accum: NewAccum(),
		},
		leaves: map[int]*Tree{},
	}
}
func (p *Tree) Find(state State, mesh *Mesh) *Node {
	var (
		node = p
		next *Tree
		ok   bool
		code = state.Hash2(mesh)
	)
	for i, s := range code {
		next, ok = node.leaves[s]
		if !ok {
			if i == len(code)-1 {
				next = &Tree{Node: NewNode(), leaves: map[int]*Tree{}}
			} else {
				next = &Tree{Node: NewNode(), leaves: map[int]*Tree{}}
			}
			node.leaves[s] = next
		}
		node = next
	}
	return node.Node
}
func (p *Tree) Clear() {
	p.Node = &Node{Accum: NewAccum()}
	p.leaves = map[int]*Tree{}
}
