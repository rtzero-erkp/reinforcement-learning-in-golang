package common

import "fmt"

var _ Accumulate = &Accum1D{}

type Accum1D struct {
	rewardCum map[ActionEnum]float64
	countCum  map[ActionEnum]float64
	count     float64
}

func (p *Accum1D) CountAct(act ActionEnum) float64 {
	return p.countCum[act]
}

func (p *Accum1D) Count() float64 {
	return p.count
}

func (p *Accum1D) String() string {
	var line = "\n"
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		line += fmt.Sprintf("[accum1D] act:%v, reward:%10.7f, count:%10.0f, mean:%10.7f\n",
			act, reward, count, p.Mean(act))
	}
	return line
}

func (p *Accum1D) Mean(act ActionEnum) float64 {
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}

func (p *Accum1D) Add(act ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
	p.count += 1
}

func NewReward1D(space Space) Accumulate {
	var p = &Accum1D{
		rewardCum: map[ActionEnum]float64{},
		countCum:  map[ActionEnum]float64{},
		count:     0,
	}
	for _, act := range space.Acts() {
		p.rewardCum[act] = 0
		p.countCum[act] = 0
	}
	return p
}
