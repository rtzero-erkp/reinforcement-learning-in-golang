package common

import "log"

type UpdateMethod struct {
	Model UpdateEnum
	Args  []interface{}
}

func NewUpdateMethod(mode UpdateEnum, args ...interface{}) *UpdateMethod {
	return &UpdateMethod{
		Model: mode,
		Args:  args,
	}
}

func (p *UpdateMethod) VMap(mem *MemoryCode, model *ModelMap) {
	codeFrom := mem.From
	codeTo := mem.To
	valueFrom := model.getV(codeFrom)
	valueTo := model.getV(codeTo)
	switch p.Model {
	case UpdateEnum_DT:
		// v(s) = v(s) + alpha * (r + lambda * (v(s') - v(s)))
		alpha := p.Args[0].(float64)
		lambda := p.Args[1].(float64)
		//log.Printf("valueFrom:%v", valueFrom)
		//log.Printf("valueTo:%v", valueTo)
		//log.Printf("alpha:%v", alpha)
		//log.Printf("lambda:%v", lambda)
		valueFrom = valueFrom + alpha*(mem.Reward+lambda*(valueTo-valueFrom))
		//log.Printf("valueFrom:%v", valueFrom)
		model.nodes[codeFrom] = valueFrom
	case UpdateEnum_SARSA:
		log.Fatal("not impl")
	}
}
func (p *UpdateMethod) QMap(mem *MemoryCode, model *ModelMap) {
	code := mem.From
	accum := model.getQ(code)
	accum.Add(mem.Act, mem.Reward)
}
