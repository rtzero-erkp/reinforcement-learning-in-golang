package agent_model

//var _ common.AgentModel = &MCTS{}
//
//type MCTS struct {
//	env    common.Env
//	mcNum  int
//	model  *common.AgentModelTree
//	method common.MethodEnum
//	arg    []interface{}
//}
//
//func (p *MCTS) Reset() {
//	p.model.Clear()
//}
//func (p *MCTS) String() string {
//	return "MCTS:" + p.method.String()
//}
//func (p *MCTS) Policy(state common.Info, space common.Space) common.ActionEnum {
//	var res *common.Result
//	var mem *common.Memory2
//	var memPath = common.NewMemPath()
//
//	for i0 := 0; i0 < p.mcNum; i0++ {
//		// reset
//		memPath.Clear()
//		var envCrt = p.env.Clone()
//		var path []common.ActionEnum
//		// find leaf point
//		for {
//			var node = p.model.Find(path...).(*common.NodeQ)
//			var spaceCrt = p.env.Space()
//			var act = node.Accum.Sample(spaceCrt, p.method, p.arg...)
//			res = envCrt.Step(act)
//			memPath.Add(path, act, res.Reward[0])
//			path = append(path, act)
//			if !res.Done {
//				break
//			}
//			if node.Accum.CountAct(act) == 0 {
//				break
//			}
//		}
//		// rand act to end
//		var reward float64 = 0
//		for !res.Done {
//			var spaceCrt = p.env.Space()
//			var act = spaceCrt.Sample()
//			res = envCrt.Step(act)
//			reward += res.Reward[0]
//		}
//		// update
//		var memories = memPath.Get()
//		var memoriesLen = len(memories)
//		for i1 := 0; i1 < memoriesLen; i1++ {
//			mem = memories[memoriesLen-1-i1]
//			var node = p.model.Find(mem.Path...).(*common.NodeQ)
//			node.Accum.Add(mem.Act, reward)
//			reward = node.Accum.Mean()
//		}
//	}
//
//	var node = p.model.Find().(*common.NodeQ)
//	var act = node.Accum.Sample(space, common.MethodEnum_MeanQ)
//	p.model.Move(act)
//	return act
//}
//func (p *MCTS) Reward(state common.Info, act common.ActionEnum, reward float64) {}
//
//func NewMCTS(env common.Env, mcNum int, method common.MethodEnum, arg ...interface{}) common.AgentModel {
//	var p = &MCTS{
//		env:    env,
//		mcNum:  mcNum,
//		method: method,
//		arg:    arg,
//		model:  common.NewModelTree(common.ModelEnum_Q),
//	}
//	return p
//}
