package envs

type Stater interface {
	GetFloat64(key string) float64
	SetFloat64(key string, val float64)
}

var _ Stater = &State{}

type State struct {
	float64s map[string]float64
}

func NewState() *State {
	return &State{
		float64s: map[string]float64{},
	}
}
func (p *State) GetFloat64(key string) float64 {
	return p.float64s[key]
}
func (p *State) SetFloat64(key string, val float64) {
	p.float64s[key] = val
}
