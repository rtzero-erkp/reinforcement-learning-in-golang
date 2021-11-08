package common

import "math/rand"

var _ Space = &SpaceVec{}

type SpaceVec struct {
	acts []ActionEnum
}

func (p *SpaceVec) Clone() Space {
	var cp = &SpaceVec{
		acts: []ActionEnum{},
	}
	for _, act := range p.acts {
		cp.acts = append(cp.acts, act)
	}
	return cp
}

func NewSpace1DByEnum(acts ...ActionEnum) Space {
	var p = &SpaceVec{
		acts: []ActionEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}
func NewSpace1DByNum(num int) Space {
	var p = &SpaceVec{
		acts: []ActionEnum{},
	}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActionEnum(i))
	}
	return p
}

func (p *SpaceVec) Contain(act ActionEnum) bool {
	for _, actI := range p.acts {
		if actI == act {
			return true
		}
	}
	return false
}
func (p *SpaceVec) Acts() []ActionEnum {
	return p.acts
}
func (p *SpaceVec) Sample() ActionEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}
