package common

import (
	"errors"
	"math/rand"
)

type Policy map[ActionEnum]float64

func (p *Policy) Sample() (act ActionEnum, err error) {
	var (
		sum    float64 = 0
		weight float64
	)

	for _, weight = range *p {
		sum += weight
	}
	if sum == 0 {
		err = errors.New("sum weight is 0")
		return
	}

	var rate = rand.Float64()
	var target = sum * rate
	sum = 0
	for act, weight = range *p {
		sum += weight
		if sum >= target {
			return
		}
	}
	return
}
