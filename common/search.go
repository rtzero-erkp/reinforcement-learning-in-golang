package common

import (
	"math"
	"math/rand"
)

var SearchArgQ = NewSearchMethod(SearchEnum_AvgQ)

type SearchMethod struct {
	Model SearchEnum
	Args  []interface{}
}

func NewSearchMethod(mode SearchEnum, args ...interface{}) *SearchMethod {
	return &SearchMethod{
		Model: mode,
		Args:  args,
	}
}

func (p *SearchMethod) QMap(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	switch p.Model {
	case SearchEnum_MC:
		act = p.qMap_MC(env)
	case SearchEnum_AvgQ:
		act = p.qMap_AvgQ(env, encoder, model)
	case SearchEnum_EpsilonGreed:
		act = p.qMap_EpsilonGreed(env, encoder, model)
	case SearchEnum_SoftMax:
		act = p.qMap_SoftMax(env, encoder, model)
	case SearchEnum_UCB:
		act = p.qMap_UCB(env, encoder, model)
	default:
		act = env.Acts().Sample()
	}
	return
}
func (p *SearchMethod) qMap_MC(env Env) (act ActEnum) {
	act = env.Acts().Sample()
	return act
}
func (p *SearchMethod) qMap_AvgQ(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	code := encoder.Hash(env.State())
	accum := model.getQ(code)
	actsMax := NewActsMax()
	for _, actI := range env.Acts().All() {
		val := accum.MeanAct(actI)
		actsMax.Add(actI, val)
	}
	act = actsMax.Sample()
	return
}
func (p *SearchMethod) qMap_EpsilonGreed(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	epsilon := p.Args[0].(float64)
	rate := rand.Float64()
	if rate < epsilon {
		return env.Acts().Sample()
	} else {
		code := encoder.Hash(env.State())
		accum := model.getQ(code)
		actsMax := NewActsMax()
		for _, actI := range env.Acts().All() {
			q := accum.MeanAct(actI)
			actsMax.Add(actI, q)
		}
		act = actsMax.Sample()
		return
	}
}
func (p *SearchMethod) qMap_SoftMax(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	tau := p.Args[0].(float64)
	code := encoder.Hash(env.State())
	accum := model.getQ(code)
	var probSum float64 = 0
	for _, actI := range env.Acts().All() {
		var q = accum.MeanAct(actI)
		probSum += math.Exp(q / tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	actsMax := NewActsMax()
	for _, actI := range env.Acts().All() {
		var q = accum.MeanAct(actI)
		var prob = math.Exp(q/tau) / probSum
		probCum += prob
		if probCum > rate {
			return actI
		}
		actsMax.Add(actI, prob)
	}
	act = actsMax.Sample()
	return
}
func (p *SearchMethod) qMap_UCB(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	code := encoder.Hash(env.State())
	accum := model.getQ(code)
	countSum := accum.Count()
	actsMax := NewActsMax()
	var upperBound float64
	for _, actI := range env.Acts().All() {
		count := accum.CountAct(actI)
		if count == 0 {
			upperBound = math.Inf(1)
		} else {
			upperBound = math.Sqrt((2 * math.Log(countSum)) / count)
		}
		q := accum.MeanAct(actI)
		ucb := upperBound + q
		actsMax.Add(actI, ucb)
	}
	act = actsMax.Sample()
	return
}

func (p *SearchMethod) VMap(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	actsMax := NewActsMax()
	for _, actI := range env.Acts().All() {
		envCrt := env.Clone()
		res := envCrt.Step(actI)
		code := encoder.Hash(res.State)
		reward := model.getV(code)
		actsMax.Add(actI, reward)
	}
	act = actsMax.Sample()
	return
}

func (p *SearchMethod) Accum(accum Accumulate, acts Acts) (act ActEnum) {
	actsMax := NewActsMax()
	for _, actI := range acts.All() {
		mean := accum.MeanAct(actI)
		actsMax.Add(actI, mean)
	}
	act = actsMax.Sample()
	return
}
