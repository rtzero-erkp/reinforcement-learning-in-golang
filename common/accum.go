package common

import "fmt"

var _ Accumulate = &Accum{}

type Accum struct {
	rewardCum map[ActionEnum]float64
	countCum  map[ActionEnum]float64
	count     float64
}

func (p *Accum) CountAct(act ActionEnum) float64 {
	return p.countCum[act]
}
func (p *Accum) Count() float64 {
	return p.count
}
func (p *Accum) String() string {
	var line = "\n"
	var actMax []ActionEnum
	var meanMax float64
	var rewardSum float64
	var countSum float64
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		var mean = p.Mean(act)
		line += fmt.Sprintf("[accum] Act:%v, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", act, reward, count, mean)
		if (len(actMax) == 0) || (mean > meanMax) {
			meanMax = mean
			actMax = []ActionEnum{act}
		} else
		if mean == meanMax {
			actMax = append(actMax, act)
		}
		rewardSum += reward
		countSum += count
	}
	line += fmt.Sprintf("[accum] best, Act:%v mean:%10.7f\n", actMax, meanMax)
	line += fmt.Sprintf("[accum] total, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", rewardSum, countSum, rewardSum/countSum)
	return line
}
func (p *Accum) Mean(act ActionEnum) float64 {
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}
func (p *Accum) Add(act ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
	p.count += 1
}

func NewAccum() Accumulate {
	var p = &Accum{
		rewardCum: map[ActionEnum]float64{},
		countCum:  map[ActionEnum]float64{},
		count:     0,
	}
	return p
}
