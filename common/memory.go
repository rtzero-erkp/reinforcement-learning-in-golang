package common

type Memory struct {
	Code   []int
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

func (p *Memories) Add(code []int, act ActionEnum, reward float64) {
	p.mem = append(p.mem, &Memory{Code: code, Act: act, Reward: reward})
}

func (p *Memories) Clear() {
	p.mem = []*Memory{}
}

func (p *Memories) Get() []*Memory {
	return p.mem
}
