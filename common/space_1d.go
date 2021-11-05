package common

import "math/rand"

var _ Space = &Space1D{}

type Space1D struct {
	acts []ActionEnum
}

func NewSpace1DByEnum(acts ...ActionEnum) Space {
	var p = &Space1D{
		acts: []ActionEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}
func NewSpace1DByNum(num int) Space {
	var p = &Space1D{
		acts: []ActionEnum{},
	}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActionEnum(i))
	}
	return p
}

func (p *Space1D) Contain(act ActionEnum) bool {
	for _, actI := range p.acts {
		if actI == act {
			return true
		}
	}
	return false
}
func (p *Space1D) Acts() []ActionEnum {
	return p.acts
}
func (p *Space1D) Sample() ActionEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}
