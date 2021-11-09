package common

import "fmt"

//type Memories interface {
//	Add(state State, act ActionEnum, reward float64)
//	Clear()
//	Get() []*Memory
//}

type Memory struct {
	State  State
	Act    ActionEnum
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


type MemoryCircle struct {
	mesh   *Mesh
	code   []string
	bench  []*Memory
	circle []*Memory
}

func NewMemoryCircle(mesh *Mesh) *MemoryCircle {
	return &MemoryCircle{
		bench:  []*Memory{},
		circle: []*Memory{},
		mesh:   mesh,
		code:   []string{},
	}
}

func (p *MemoryCircle) Add(state State, act ActionEnum, reward float64) {
	var mem = &Memory{State: state.Clone(), Act: act, Reward: reward}
	var code = fmt.Sprintf("%v", state.Hash(p.mesh))
	for i, codeI := range p.code {
		if codeI == code {
			p.circle = append(p.circle, p.bench[i:]...)
			p.bench = p.bench[:i]
			p.code = p.code[:i]
			return
		}
	}
	p.bench = append(p.bench, mem)
	p.code = append(p.code, code)
	//log.Printf("[memC] mem:%v", mem)
	//log.Printf("[memC] code:%v", code)
	//log.Printf("[memC] code.len:%v", len(p.code))
	//log.Printf("[memC] bench.len:%v", len(p.bench))
	//log.Printf("[memC] circle.len:%v", len(p.circle))
}
func (p *MemoryCircle) Clear() {
	p.code = []string{}
	p.bench = []*Memory{}
	p.circle = []*Memory{}
}
func (p *MemoryCircle) GetCircle() []*Memory {
	return p.circle
}
func (p *MemoryCircle) GetBench() []*Memory {
	return p.bench
}
