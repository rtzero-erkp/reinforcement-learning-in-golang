package common

type AgentModel interface {
	Train(env Env, trainNum int) interface{}
	Policy(env Env) (act ActionEnum)
	String() string
}
