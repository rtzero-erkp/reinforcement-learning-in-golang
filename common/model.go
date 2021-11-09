package common

type Model interface {
	Find(state State, mesh *Mesh) *Node
}

type Node struct {
	Accum  Accumulate // 累计收益
	Policy Policy     // 策略梯度
}

var _ Model = &Tree{}

type Tree struct {
	*Node
	leaves map[int]*Tree // 树结构
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
				next = &Tree{
					Node: &Node{
						Accum:  NewAccum(),
						Policy: NewPolicyPlus(),
					},
					leaves: map[int]*Tree{},
				}
			} else {
				next = &Tree{
					Node: &Node{
						Accum:  NewAccum(),
						Policy: NewPolicyPlus(),
					},
					leaves: map[int]*Tree{},
				}
			}
			node.leaves[s] = next
		}
		node = next
	}
	return node.Node
}
func NewTree() Model {
	return &Tree{
		Node: &Node{
			Accum:  NewAccum(),
			Policy: NewPolicyPlus(),
		},
		leaves: map[int]*Tree{},
	}
}

var _ Model = &HashMap{}

type HashMap struct {
	nodes map[string]*Node // 树结构
}

func (p *HashMap) Find(state State, mesh *Mesh) *Node {
	return p.nodes[state.Hash(mesh)]
}
func NewHashMap() Model {
	return &HashMap{
		nodes: map[string]*Node{}, // 表结构
	}
}
