package common

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Accumulate interface {
	Reset()
	CountAct(act ActEnum) float64
	Count() float64
	Mean() float64
	MeanAct(act ActEnum) float64
	Add(act ActEnum, reward float64)
	Sample(acts Acts, search *SearchParam) ActEnum
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
func (p *accumulate) Sample(acts Acts, search *SearchParam) ActEnum {
	var act ActEnum
	switch search.Model {
	case SearchEnum_MC:
		act = p.sampleMC(acts)
	case SearchEnum_AvgQ:
		act = p.sampleMeanQ(acts)
	case SearchEnum_EpsilonGreed:
		act = p.sampleEpsilonGreed(acts, search.Args[0].(float64))
	case SearchEnum_SoftMax:
		act = p.sampleSoftMax(acts, search.Args[0].(float64))
	case SearchEnum_UCB:
		act = p.sampleUCB(acts)
	default:
		act = acts.Sample()
	}

	return act
}
func (p *accumulate) sampleMC(acts Acts) ActEnum {
	return acts.Sample()
}
func (p *accumulate) sampleMeanQ(acts Acts) ActEnum {
	var actsMax []ActEnum
	var valMax float64
	for _, act := range acts.All() {
		val := p.MeanAct(act)
		if (len(actsMax) == 0) || (val > valMax) {
			actsMax = []ActEnum{act}
			valMax = val
		} else
		if val == valMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("acts is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
func (p *accumulate) sampleEpsilonGreed(acts Acts, epsilon float64) ActEnum {
	var rate = rand.Float64()
	if rate < epsilon {
		return acts.Sample()
	} else {
		var qMax float64
		var actsMax []ActEnum

		for _, act := range acts.All() {
			var q = p.MeanAct(act)
			if (len(actsMax) == 0) || (q > qMax) {
				qMax = q
				actsMax = []ActEnum{act}
			} else
			if q == qMax {
				actsMax = append(actsMax, act)
			}
		}
		if len(actsMax) == 0 {
			log.Fatalln("acts is nil")
		}
		var idx = rand.Intn(len(actsMax))
		return actsMax[idx]
	}
}
func (p *accumulate) sampleSoftMax(acts Acts, tau float64) ActEnum {
	var probSum float64 = 0
	for _, act := range acts.All() {
		var q = p.MeanAct(act)
		probSum += math.Exp(q / tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	var probMax float64 = 0
	var actsMax []ActEnum
	for _, act := range acts.All() {
		var q = p.MeanAct(act)
		var prob = math.Exp(q/tau) / probSum
		probCum += prob
		if probCum > rate {
			return act
		}
		if prob > probMax {
			actsMax = []ActEnum{act}
			probMax = prob
		} else
		if prob == probMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("acts is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
func (p *accumulate) sampleUCB(acts Acts) ActEnum {
	//log.Printf("[accum] acts:%v", acts)
	countSum := p.Count()
	var actsMax []ActEnum
	var ucbMax float64
	var upperBound float64
	for _, act := range acts.All() {
		count := p.CountAct(act)
		if count == 0 {
			upperBound = math.Inf(1)
		} else {
			upperBound = math.Sqrt((2 * math.Log(countSum)) / count)
		}
		q := p.MeanAct(act)
		ucb := upperBound + q
		if (len(actsMax) == 0) || ucb > ucbMax {
			ucbMax = ucb
			actsMax = []ActEnum{act}
		} else
		if ucb == ucbMax {
			actsMax = append(actsMax, act)
		}
	}
	if len(actsMax) == 0 {
		log.Fatalln("acts is nil")
	}
	var idx = rand.Intn(len(actsMax))
	return actsMax[idx]
}
