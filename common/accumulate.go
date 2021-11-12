package common

import (
	"fmt"
)

type Accumulate interface {
	Reset()
	CountAct(act ActEnum) float64
	Count() float64
	Mean() float64
	Exist(act ActEnum) (ok bool)
	MeanAct(act ActEnum) float64
	Add(act ActEnum, reward float64)
	String() string
}

var _ Accumulate = &accumulate{}

type accumulate struct {
	rewardCum map[ActEnum]float64
	countCum  map[ActEnum]float64
	reward    float64
	count     float64
}

func NewAccumulate() Accumulate {
	var p = &accumulate{
		rewardCum: map[ActEnum]float64{},
		countCum:  map[ActEnum]float64{},
		count:     0,
		reward:    0,
	}
	return p
}
func (p *accumulate) Reset() {
	p.rewardCum = map[ActEnum]float64{}
	p.countCum = map[ActEnum]float64{}
	p.count = 0
	p.reward = 0
}
func (p *accumulate) CountAct(act ActEnum) float64 {
	return p.countCum[act]
}
func (p *accumulate) Count() float64 {
	return p.count
}
func (p *accumulate) Mean() float64 {
	return p.reward / p.count
}
func (p *accumulate) Exist(act ActEnum) (ok bool) {
	_, ok = p.countCum[act]
	return ok
}
func (p *accumulate) MeanAct(act ActEnum) float64 {
	//log.Printf("count:%v, reward:%v", p.countCum[act], p.rewardCum[act])
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}
func (p *accumulate) Add(act ActEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
	p.reward += reward
	p.count += 1
}
func (p *accumulate) String() string {
	var line = "\n"
	var actMax []ActEnum
	var meanMax float64
	var rewardSum float64
	var countSum float64
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		var mean = p.MeanAct(act)
		line += fmt.Sprintf("[accumulate] Act:%v, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", act, reward, count, mean)
		if (len(actMax) == 0) || (mean > meanMax) {
			meanMax = mean
			actMax = []ActEnum{act}
		} else
		if mean == meanMax {
			actMax = append(actMax, act)
		}
		rewardSum += reward
		countSum += count
	}
	var meanSum float64
	if countSum == 0 {
		meanSum = 0
	} else {
		meanSum = rewardSum / countSum
	}
	line += fmt.Sprintf("[accumulate] best, Act:%v mean:%10.7f\n", actMax, meanMax)
	line += fmt.Sprintf("[accumulate] total, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", rewardSum, countSum, meanSum)
	return line
}
