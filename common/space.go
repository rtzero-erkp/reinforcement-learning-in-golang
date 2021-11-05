package common

import "math/rand"


type Space struct {
	acts []ActionEnum
}

func NewActions(acts ...ActionEnum) *Space {
	var p = &Space{
		acts: []ActionEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}
func NewActionsByNum(num int) *Space {
	var p = &Space{
		acts: []ActionEnum{},
	}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActionEnum(i))
	}
	return p
}

func (p *Space) Contain(act ActionEnum) bool {
	for _, actI := range p.acts {
		if actI == act {
			return true
		}
	}
	return false
}
func (p *Space) Acts() []ActionEnum {
	return p.acts
}
func (p *Space) Sample() ActionEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}
