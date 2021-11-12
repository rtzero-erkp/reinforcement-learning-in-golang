package common

import (
	"fmt"
	"log"
	"math/rand"
)

type Acts interface {
	Contain(act ActEnum) bool
	All() []ActEnum
	Sample() ActEnum
	Shuffle()
	String() string
	Clone() Acts
	Clear()
	AddEnum(acts ...ActEnum)
	AddInt(acts ...int)
	AddNum(num int)
}

var _ Acts = &ActsVec{}

type ActsVec struct {
	acts []ActEnum
}

func NewActsVec() Acts {
	var p = &ActsVec{
		acts: []ActEnum{},
	}
	return p
}
func NewActsVecByEnum(acts ...ActEnum) Acts {
	var p = &ActsVec{
		acts: []ActEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}
func NewActsVecByInt(acts ...int) Acts {
	var p = &ActsVec{
		acts: []ActEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, ActEnum(fmt.Sprint(act)))
	}
	return p
}
func NewActsVecByNum(num int) Acts {
	var p = &ActsVec{
		acts: []ActEnum{},
	}
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActEnum(fmt.Sprint(i)))
	}
	return p
}

func (p *ActsVec) Contain(act ActEnum) bool {
	for _, actI := range p.acts {
		if actI.Same(act) {
			return true
		}
	}
	return false
}
func (p *ActsVec) All() []ActEnum {
	return p.acts
}
func (p *ActsVec) Sample() ActEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}
func (p *ActsVec) String() string {
	var line = ""
	for _, act := range p.acts {
		line += act.String() + " "
	}
	return line
}
func (p *ActsVec) Shuffle() {
	var num = len(p.acts)
	for i := 0; i < num; i++ {
		var idx = rand.Intn(num)
		var tmp = p.acts[idx]
		p.acts[idx] = p.acts[i]
		p.acts[i] = tmp
	}
}
func (p *ActsVec) Clone() Acts {
	var cp = &ActsVec{
		acts: []ActEnum{},
	}
	for _, act := range p.acts {
		cp.acts = append(cp.acts, act)
	}
	return cp
}
func (p *ActsVec) AddEnum(acts ...ActEnum) {
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
}
func (p *ActsVec) AddInt(acts ...int) {
	for _, act := range acts {
		p.acts = append(p.acts, ActEnum(fmt.Sprint(act)))
	}
}
func (p *ActsVec) AddNum(num int) {
	for i := 0; i < num; i++ {
		p.acts = append(p.acts, ActEnum(fmt.Sprint(i)))
	}
}
func (p *ActsVec) Clear() {
	p.acts = []ActEnum{}
}

type ActsMax struct {
	acts []ActEnum
	max  float64
}

func NewActsMax() *ActsMax {
	return &ActsMax{
		acts: []ActEnum{},
		max:  0,
	}
}
func (p *ActsMax) Add(act ActEnum, val float64) {
	if (len(p.acts) == 0) || (val > p.max) {
		p.acts = []ActEnum{act}
		p.max = val
	} else
	if val == p.max {
		p.acts = append(p.acts, act)
	}
}
func (p *ActsMax) Sample() (act ActEnum) {
	if len(p.acts) == 0 {
		log.Fatalln("acts is nil")
	}
	var idx = rand.Intn(len(p.acts))
	act = p.acts[idx]
	return
}
