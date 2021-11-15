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
	//log.Printf("env:%v", env)
	//log.Printf("state:%v", env.State())
	//log.Printf("encoder:%v", encoder)
	//log.Printf("model:%v", model)
	code := encoder.Hash(env.State())
	//log.Printf("code:%v", code)
	accum := model.getQ(code)
	//log.Printf("accum:%v", accum)
	acts := env.Acts()
	//log.Printf("acts:%v", acts)
	act = p.Accum(accum, acts)
	//log.Printf("act:%v", act)
	return
}
func (p *SearchMethod) VMap(env Env, encoder Encoder, model *ModelMap) (act ActEnum) {
	accum := NewAccumulate()
	acts := env.Acts()
	for _, actI := range acts.All() {
		envCrt := env.Clone()
		res := envCrt.Step(actI)
		code := encoder.Hash(res.State)
		reward := model.getV(code)
		accum.Add(actI, reward)
	}
	p.Accum(accum, acts)
	return
}
func (p *SearchMethod) Accum(accum Accumulate, acts Acts) (act ActEnum) {
	switch p.Model {
	case SearchEnum_MC:
		act = p.mc(acts)
	case SearchEnum_AvgQ:
		act = p.avgQ(accum, acts)
	case SearchEnum_EpsilonGreed:
		act = p.epsilonGreed(accum, acts)
	case SearchEnum_SoftMax:
		act = p.softmax(accum, acts)
	case SearchEnum_UCB:
		act = p.ucb(accum, acts)
	default:
		act = acts.Sample()
	}
	return
}

func (p *SearchMethod) mc(acts Acts) (act ActEnum) {
	act = acts.Sample()
	return act
}
func (p *SearchMethod) avgQ(accum Accumulate, acts Acts) (act ActEnum) {
	actsMax := NewActsMax()
	for _, actI := range acts.All() {
		val := accum.MeanAct(actI)
		actsMax.Add(actI, val)
	}
	act = actsMax.Sample()
	return
}
func (p *SearchMethod) epsilonGreed(accum Accumulate, acts Acts) (act ActEnum) {
	epsilon := p.Args[0].(float64)
	rate := rand.Float64()
	if rate < epsilon {
		return acts.Sample()
	} else {
		actsMax := NewActsMax()
		for _, actI := range acts.All() {
			q := accum.MeanAct(actI)
			actsMax.Add(actI, q)
		}
		act = actsMax.Sample()
		return
	}
}
func (p *SearchMethod) softmax(accum Accumulate, acts Acts) (act ActEnum) {
	tau := p.Args[0].(float64)
	var probSum float64 = 0
	for _, actI := range acts.All() {
		var q = accum.MeanAct(actI)
		probSum += math.Exp(q / tau)
	}
	var rate = rand.Float64()
	var probCum float64 = 0
	actsMax := NewActsMax()
	for _, actI := range acts.All() {
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
func (p *SearchMethod) ucb(accum Accumulate, acts Acts) (act ActEnum) {
	countSum := accum.Count()
	actsMax := NewActsMax()
	var upperBound float64
	for _, actI := range acts.All() {
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
