package envs

import "math/rand"

type AuctionSpacer interface {
	Contain(act ActionEnum) bool
	Sample() ActionEnum
}

var _ AuctionSpacer = &ActionSpace{}

type ActionSpace struct {
	acts []ActionEnum
}

func (p *ActionSpace) Sample() ActionEnum {
	var idx = rand.Intn(len(p.acts))
	return p.acts[idx]
}

func NewActions(acts ...ActionEnum) AuctionSpacer {
	var p = &ActionSpace{
		acts: []ActionEnum{},
	}
	for _, act := range acts {
		p.acts = append(p.acts, act)
	}
	return p
}

func (p *ActionSpace) Contain(act ActionEnum) bool {
	for _, actI := range p.acts {
		if actI == act {
			return true
		}
	}
	return false
}
