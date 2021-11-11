package common

type Memory struct {
	From   Info
	Act    ActionEnum
	To     Info
	Reward float64
}

type MemoryCode struct {
	From   string
	Act    ActionEnum
	To     string
	Reward float64
}

type Mem struct {
	mem []*Memory
}

func NewMem() *Mem {
	return &Mem{
		mem: []*Memory{},
	}
}
func (p *Mem) Add(from Info, act ActionEnum, to Info, reward float64) {
	var mem = &Memory{
		From:   from.Clone(),
		Act:    act,
		To:     to.Clone(),
		Reward: reward}
	p.mem = append(p.mem, mem)
}
func (p *Mem) Clear() {
	p.mem = []*Memory{}
}
func (p *Mem) Get() []*Memory {
	return p.mem
}

type MemCode struct {
	mem []*MemoryCode
}

func NewMemCode() *MemCode {
	return &MemCode{
		mem: []*MemoryCode{},
	}
}
func (p *MemCode) Add(from string, act ActionEnum, to string, reward float64) {
	var mem = &MemoryCode{
		From:   from,
		Act:    act,
		To:     to,
		Reward: reward}
	p.mem = append(p.mem, mem)
}
func (p *MemCode) Clear() {
	p.mem = []*MemoryCode{}
}
func (p *MemCode) Get() []*MemoryCode {
	return p.mem
}
