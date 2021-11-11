package agent_model

//var _ common.AgentModel = &Cfr{}
//
//type Cfr struct {
//	model *common.ModelTree // 模型
//}
//
//func (p *Cfr) Reset() {
//	p.model.Clear()
//}
//func (p *Cfr) String() string {
//	return "Cfr"
//}
//func (p *Cfr) Policy(state common.Info, space common.Space) common.ActionEnum {
//	return space.Sample()
//}
//func (p *Cfr) Reward(state common.Info, act common.ActionEnum, reward float64) {
//}
//
//func NewCFR() common.AgentModel {
//	var p = &Cfr{
//		model: common.NewModelTree(common.ModelEnum_Cfr),
//	}
//	return p
//}
