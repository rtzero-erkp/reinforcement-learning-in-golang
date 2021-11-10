package common

type Memory struct {
	Info   Info
	Act    ActionEnum
	Reward float64
}
type Memory2 struct {
	Path   []ActionEnum
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
func (p *Mem) Add(state Info, act ActionEnum, reward float64) {
	var mem = &Memory{Info: state.Clone(), Act: act, Reward: reward}
	p.mem = append(p.mem, mem)
}
func (p *Mem) Clear() {
	p.mem = []*Memory{}
}
func (p *Mem) Get() []*Memory {
	return p.mem
}

type MemoryCircle struct {
	encoder Encoder
	code    []string
	bench   []*Memory
	circle  []*Memory
}

func NewMemoryCircle(encoder Encoder) *MemoryCircle {
	return &MemoryCircle{
		bench:   []*Memory{},
		circle:  []*Memory{},
		encoder: encoder,
		code:    []string{},
	}
}
func (p *MemoryCircle) Add(state Info, act ActionEnum, reward float64) {
	var mem = &Memory{Info: state.Clone(), Act: act, Reward: reward}
	var code = p.encoder.Hash(state)
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

type MemPath struct {
	mem []*Memory2
}

func NewMemPath() *MemPath {
	return &MemPath{
		mem: []*Memory2{},
	}
}
func (p *MemPath) Add(path []ActionEnum, act ActionEnum, reward float64) {
	var mem = &Memory2{Path: path, Act: act, Reward: reward}
	p.mem = append(p.mem, mem)
}
func (p *MemPath) Clear() {
	p.mem = []*Memory2{}
}
func (p *MemPath) Get() []*Memory2 {
	return p.mem
}
