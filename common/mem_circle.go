package common

import (
	"fmt"
)

//var _ Memories = &MemoryCircle{}

type MemoryCircle struct {
	mesh   []float64
	code   []string
	bench  []*Memory
	circle []*Memory
}

func NewMemoryCircle(mesh []float64) *MemoryCircle {
	return &MemoryCircle{
		bench:  []*Memory{},
		circle: []*Memory{},
		mesh:   mesh,
		code:   []string{},
	}
}

func (p *MemoryCircle) Add(state State, act ActionEnum, reward float64) {
	var mem = &Memory{State: state.Clone(), Act: act, Reward: reward}
	var code = fmt.Sprintf("%v", state.Encode(p.mesh))
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
