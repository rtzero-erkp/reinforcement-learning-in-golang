package common

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Accumulate interface {
	Reset()
	CountAct(act ActionEnum) float64
	Count() float64
	Mean() float64
	MeanAct(act ActionEnum) float64
	Add(act ActionEnum, reward float64)
	Sample(space Space, search *SearchParam) ActionEnum
	String() string
}

var _ Accumulate = &accumulate{}

type accumulate struct {
	rewardCum map[ActionEnum]float64
	countCum  map[ActionEnum]float64
	reward    float64
	count     float64
}

func NewAccumulate() Accumulate {
	var p = &accumulate{
		rewardCum: map[ActionEnum]float64{},
		countCum:  map[ActionEnum]float64{},
		count:     0,
		reward:    0,
	}
	return p
}
func (p *accumulate) Reset() {
	p.rewardCum = map[ActionEnum]float64{}
	p.countCum = map[ActionEnum]float64{}
	p.count = 0
	p.reward = 0
}
func (p *accumulate) CountAct(act ActionEnum) float64 {
	return p.countCum[act]
}
func (p *accumulate) Count() float64 {
	return p.count
}
func (p *accumulate) Mean() float64 {
	return p.reward / p.count
}
func (p *accumulate) MeanAct(act ActionEnum) float64 {
	//log.Printf("count:%v, reward:%v", p.countCum[act], p.rewardCum[act])
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}
func (p *accumulate) Add(act ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
	p.reward += reward
	p.count += 1
}
func (p *accumulate) String() string {
	var line = "\n"
	var actMax []ActionEnum
	var meanMax float64
	var rewardSum float64
	var countSum float64
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		var mean = p.MeanAct(act)
		line += fmt.Sprintf("[accumulate] Act:%v, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", act, reward, count, mean)
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
func (p *accumulate) Sample(space Space, search *SearchParam) ActionEnum {
	for _, act := range space.Acts() {
		if p.countCum[act] == 0 {
			return act
		}
	}

	var act ActionEnum
	switch search.Model {
	case SearchEnum_MC:
		act = p.sampleMC(space)
	case SearchEnum_AvgQ:
		act = p.sampleMeanQ(space)
	case SearchEnum_EpsilonGreed:
		act = p.sampleEpsilonGreed(space, search.Args[0].(float64))
	case SearchEnum_SoftMax:
		act = p.sampleSoftMax(space, search.Args[0].(float64))
	case SearchEnum_UCB:
		act = p.sampleUCB(space)
	default:
		act = space.Sample()
	}

	return act
}
func (p *accumulate) sampleMC(space Space) ActionEnum {
	return space.Sample()
}
func (p *accumulate) sampleMeanQ(space Space) ActionEnum {
	var actsMax []ActionEnum
	var valMax float64
	for _, act := range space.Acts() {
		val := p.MeanAct(act)
		if (len(actsMax) == 0) || (val > valMax) {
			actsMax = []ActionEnum{act}
			valMax = val
		} else
		if val == valMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("space is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
func (p *accumulate) sampleEpsilonGreed(space Space, epsilon float64) ActionEnum {
	var rate = rand.Float64()
	if rate < epsilon {
		return space.Sample()
	} else {
		var qMax float64
		var actsMax []ActionEnum

		for _, act := range space.Acts() {
			var q = p.MeanAct(act)
			if (len(actsMax) == 0) || (q > qMax) {
				qMax = q
				actsMax = []ActionEnum{act}
			} else
			if q == qMax {
				actsMax = append(actsMax, act)
			}
		}
		if len(actsMax) == 0 {
			log.Fatalln("space is nil")
		}
		var idx = rand.Intn(len(actsMax))
		return actsMax[idx]
	}
}
func (p *accumulate) sampleSoftMax(space Space, tau float64) ActionEnum {
	var probSum float64 = 0
	for _, act := range space.Acts() {
		var q = p.MeanAct(act)
		probSum += math.Exp(q / tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	var probMax float64 = 0
	var actsMax []ActionEnum
	for _, act := range space.Acts() {
		var q = p.MeanAct(act)
		var prob = math.Exp(q/tau) / probSum
		probCum += prob
		if probCum > rate {
			return act
		}
		if prob > probMax {
			actsMax = []ActionEnum{act}
			probMax = prob
		} else
		if prob == probMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("space is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
func (p *accumulate) sampleUCB(space Space) ActionEnum {
	countSum := p.Count()
	var actsMax []ActionEnum
	var ucbMax float64
	for _, act := range space.Acts() {
		count := p.CountAct(act)
		if count == 0 {
			return act
		}
		upperBound := math.Sqrt((2 * math.Log(countSum)) / count)
		q := p.MeanAct(act)
		ucb := upperBound + q
		if (len(actsMax) == 0) || ucb > ucbMax {
			ucbMax = ucb
			actsMax = []ActionEnum{act}
		} else
		if ucb == ucbMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("space is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
