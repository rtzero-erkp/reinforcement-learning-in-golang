package common

import (
	"fmt"
	"math/rand"
)

type Policy interface {
	Sample() (act ActionEnum)
	Set(act ActionEnum, weight float64)
	Clean()
	String() string
}

var _ Policy = &PolicyPlus{}

type PolicyPlus map[ActionEnum]float64

func NewPolicyPlus() Policy {
	return &PolicyPlus{}
}

func (p *PolicyPlus) Sample() (act ActionEnum) {
	var (
		sum    float64 = 0
		weight float64
	)

	for _, weight = range *p {
		sum += weight
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
func (p *PolicyPlus) Set(act ActionEnum, weight float64) {
	if weight < 0 {
		weight = 0
	}
	(*p)[act] = weight
}
func (p *PolicyPlus) Clean() {
	for key := range *p {
		delete(*p, key)
	}
}
func (p *PolicyPlus) String() string {
	var line = ""
	for k, v := range *p {
		line += fmt.Sprintf("[policy] %v:%f\n", k, v)
	}
	return line
}
