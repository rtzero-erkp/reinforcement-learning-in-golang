package common

import (
	"fmt"
	"math/rand"
)

type ActionEnum string

const (
	ActionEnum_Unknown ActionEnum = "unknown"

	ActionEnum_Up    ActionEnum = "Up"
	ActionEnum_Down  ActionEnum = "Down"
	ActionEnum_Right ActionEnum = "Right"
	ActionEnum_Left  ActionEnum = "Left"

	ActionEnum_Card2 ActionEnum = "Card2"
	ActionEnum_Card3 ActionEnum = "Card3"
	ActionEnum_Card4 ActionEnum = "Card4"
	ActionEnum_Card5 ActionEnum = "Card5"
	ActionEnum_Card6 ActionEnum = "Card6"
	ActionEnum_Card7 ActionEnum = "Card7"
	ActionEnum_Card8 ActionEnum = "Card8"
	ActionEnum_Card9 ActionEnum = "Card9"
	ActionEnum_CardT ActionEnum = "CardT"
	ActionEnum_CardJ ActionEnum = "CardJ"
	ActionEnum_CardQ ActionEnum = "CardQ"
	ActionEnum_CardK ActionEnum = "CardK"
	ActionEnum_CardA ActionEnum = "CardA"

	ActionEnum_Fold  ActionEnum = "Fold"
	ActionEnum_Check ActionEnum = "Check"
	ActionEnum_Call  ActionEnum = "Call"
	ActionEnum_Bet   ActionEnum = "Bet"
	ActionEnum_AllIn ActionEnum = "AllIn"
)

func (p ActionEnum) String() string {
	return string(p)
}

type Space interface {
	Contain(act ActionEnum) bool
	Acts() []ActionEnum
	Sample() ActionEnum
	Shuffle()
	String() string
	Clone() Space
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
