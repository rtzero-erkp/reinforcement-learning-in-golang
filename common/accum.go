package common

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

func (p SearchMethod) String() string {
	var line string
	switch p {
	case SearchMethodEnum_Random:
		line = "Random"
	case SearchMethodEnum_MeanQ:
		line = "MeanQ"
	case SearchMethodEnum_EpsilonGreed:
		line = "EpsilonGreed"
	case SearchMethodEnum_SoftMax:
		line = "SoftMax"
	case SearchMethodEnum_UCB:
		line = "UCB"
	default:
		line = "unknown"
	}
	return line
}

type Accumulate interface {
	Reset()
	CountAct(act ActionEnum) float64
	Count() float64
	Mean() float64
	MeanAct(act ActionEnum) float64
	Add(act ActionEnum, reward float64)
	Sample(space Space, method SearchMethod, arg ...interface{}) ActionEnum
	String() string
}

var _ Accumulate = &accum{}

type accum struct {
	rewardCum map[ActionEnum]float64
	countCum  map[ActionEnum]float64
	reward    float64
	count     float64
}

func (p *accum) Reset() {
	p.rewardCum = map[ActionEnum]float64{}
	p.countCum = map[ActionEnum]float64{}
	p.count = 0
	p.reward = 0
}
func (p *accum) CountAct(act ActionEnum) float64 {
	return p.countCum[act]
}
func (p *accum) Count() float64 {
	return p.count
}
func (p *accum) Mean() float64 {
	return p.reward / p.count
}
func (p *accum) MeanAct(act ActionEnum) float64 {
	//log.Printf("count:%v, reward:%v", p.countCum[act], p.rewardCum[act])
	if p.countCum[act] == 0 {
		return 0
	} else {
		return p.rewardCum[act] / p.countCum[act]
	}
}
func (p *accum) Add(act ActionEnum, reward float64) {
	p.rewardCum[act] += reward
	p.countCum[act] += 1
	p.reward += reward
	p.count += 1
}
func (p *accum) Sample(space Space, method SearchMethod, arg ...interface{}) ActionEnum {
	for _, act := range space.Acts() {
		if p.countCum[act] == 0 {
			return act
		}
	}

	var act ActionEnum
	switch method {
	case SearchMethodEnum_Random:
		act = p.SampleRandom(space)
	case SearchMethodEnum_MeanQ:
		act = p.SampleMeanQ(space)
	case SearchMethodEnum_EpsilonGreed:
		act = p.SampleEpsilonGreed(space, arg[0].(float64))
	case SearchMethodEnum_SoftMax:
		act = p.SampleSoftMax(space, arg[0].(float64))
	case SearchMethodEnum_UCB:
		act = p.SampleUCB(space)
	default:
		act = space.Sample()
	}

	return act
}
func (p *accum) String() string {
	var line = "\n"
	var actMax []ActionEnum
	var meanMax float64
	var rewardSum float64
	var countSum float64
	for act, reward := range p.rewardCum {
		var count = p.countCum[act]
		var mean = p.MeanAct(act)
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
	var meanSum float64
	if countSum == 0 {
		meanSum = 0
	} else {
		meanSum = rewardSum / countSum
	}
	line += fmt.Sprintf("[accum] best, Act:%v mean:%10.7f\n", actMax, meanMax)
	line += fmt.Sprintf("[accum] total, Reward:%10.7f, count:%10.0f, mean:%10.7f\n", rewardSum, countSum, meanSum)
	return line
}

func NewAccum() Accumulate {
	var p = &accum{
		rewardCum: map[ActionEnum]float64{},
		countCum:  map[ActionEnum]float64{},
		count:     0,
		reward:    0,
	}
	return p
}

func (p *accum) SampleRandom(space Space) ActionEnum {
	return space.Sample()
}
func (p *accum) SampleMeanQ(space Space) ActionEnum {
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
func (p *accum) SampleEpsilonGreed(space Space, epsilon float64) ActionEnum {
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
func (p *accum) SampleSoftMax(space Space, tau float64) ActionEnum {
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
func (p *accum) SampleUCB(space Space) ActionEnum {
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
