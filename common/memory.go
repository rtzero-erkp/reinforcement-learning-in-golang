package common

type Memory struct {
	State  State
	Act    ActionEnum
	Reward float64
}

type Memories struct {
	mem []*Memory
}

func NewMemory() *Memories {
	return &Memories{
		mem: []*Memory{},
	}
}

func (p *Memories) Add(state State, act ActionEnum, reward float64) {
	var mem = &Memory{State: state.Clone(), Act: act, Reward: reward}
	p.mem = append(p.mem, mem)
}

func (p *Memories) Clear() {
	p.mem = []*Memory{}
}

func (p *Memories) Get() []*Memory {
	return p.mem
}
