package common

//var _ Memories = &Mem{}

type Mem struct {
	mem []*Memory
}

func NewMem() *Mem {
	return &Mem{
		mem: []*Memory{},
	}
}

func (p *Mem) Add(state State, act ActionEnum, reward float64) {
	var mem = &Memory{State: state.Clone(), Act: act, Reward: reward}
	p.mem = append(p.mem, mem)
}

func (p *Mem) Clear() {
	p.mem = []*Memory{}
}

func (p *Mem) Get() []*Memory {
	return p.mem
}
