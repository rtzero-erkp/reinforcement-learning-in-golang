package common

type Model interface {
	Find(state []int) Model
	Accum() Accumulate
	Policy() Policy
}

var _ Model = &Node{}

type Node struct {
	accum  Accumulate    // 累计收益
	policy Policy        // 策略梯度
	leaves map[int]*Node // 树结构
}

func (p *Node) Accum() Accumulate {
	return p.accum
}
func (p *Node) Policy() Policy {
	return p.policy
}
func (p *Node) Find(state []int) Model {
	var (
		node = p
		next *Node
		ok   bool
	)

	for i, s := range state {
		next, ok = node.leaves[s]
		if !ok {
			if i == len(state)-1 {
				next = NewEndNode()
			} else {
				next = NewRootNode()
			}
			node.leaves[s] = next
		}
		node = next
	}
	return node
}

func NewRootNode() *Node {
	return &Node{
		leaves: map[int]*Node{},
	}
}
func NewEndNode() *Node {
	return &Node{
		accum:  NewAccum(),
		policy: NewPolicyPlus(),
		leaves: map[int]*Node{},
	}
}
