package common

import (
	"fmt"
	"math/rand"
)

type Space interface {
	Contain(act ActionEnum) bool
	Acts() []ActionEnum
	Sample() ActionEnum
	Shuffle()
	String() string
	Clone() Space
	SetByEnum(acts ...ActionEnum)
	SetByNum(num int)
}

var _ Space = &spaceVec{}

type spaceVec struct {
	acts []ActionEnum
}

func NewSpaceVecByEnum(acts ...ActionEnum) Space {
	var p = &spaceVec{
		acts: []ActionEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}
func NewSpaceVecByNum(num int) Space {
	var p = &spaceVec{
		acts: []ActionEnum{},
	}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActionEnum(fmt.Sprint(i)))
	}
	return p
}

func (p *spaceVec) Contain(act ActionEnum) bool {
	for _, actI := range p.acts {
		if actI == act {
			return true
		}
	}
	return false
}
func (p *spaceVec) Acts() []ActionEnum {
	return p.acts
}
func (p *spaceVec) Sample() ActionEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}
func (p *spaceVec) String() string {
	var line = ""
	for _, act := range p.acts {
		line += act.String() + " "
	}
	return line
}
func (p *spaceVec) Shuffle() {
	var num = len(p.acts)
	for i := 0; i < num; i++ {
		var idx = rand.Intn(num)
		var tmp = p.acts[idx]
		p.acts[idx] = p.acts[i]
		p.acts[i] = tmp
	}
}
func (p *spaceVec) Clone() Space {
	var cp = &spaceVec{
		acts: []ActionEnum{},
	}
	for _, act := range p.acts {
		cp.acts = append(cp.acts, act)
	}
	return cp
}
func (p *spaceVec) SetByEnum(acts ...ActionEnum) {
	p.acts = []ActionEnum{}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
}
func (p *spaceVec) SetByNum(num int) {
	p.acts = []ActionEnum{}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActionEnum(fmt.Sprint(i)))
	}
}
