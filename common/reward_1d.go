package common

import "fmt"

var _ Reward = &Reward1D{}

type Reward1D struct {
	rewardCum map[ActionEnum]float64
	countCum  map[ActionEnum]float64
}

func (p *Reward1D) String() string {
	var line = "\n"
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		line += fmt.Sprintf("[reward1D] act:%v, reward:%10.7f, count:%10.0f, mean:%10.7f\n",
			act, reward, count, p.Mean(act))
	}
	return line
}

func (p *Reward1D) Mean(act ActionEnum) float64 {
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}

func (p *Reward1D) Add(act ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
}

func NewReward1D(space Space) Reward {
	var p = &Reward1D{
		rewardCum: map[ActionEnum]float64{},
		countCum:  map[ActionEnum]float64{},
	}
	for _, act := range space.Acts() {
		p.rewardCum[act] = 0
		p.countCum[act] = 0
	}
	return p
}
