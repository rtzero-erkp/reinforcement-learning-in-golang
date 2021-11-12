package common

import "fmt"

type MemoryCode struct {
	From   Code
	Act    ActEnum
	To     Code
	Reward float64
}

func (p *MemoryCode) String() string {
	var line = fmt.Sprintf("From:%v, Act:%v, To:%v, Reward:%v", p.From, p.Act, p.To, p.Reward)
	return line
}

type MemCode struct {
	mem []*MemoryCode
}

func NewMemCode() *MemCode {
	return &MemCode{
		mem: []*MemoryCode{},
	}
}
func (p *MemCode) Add(from Code, act ActEnum, to Code, reward float64) {
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
