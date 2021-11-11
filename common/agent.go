package common

type Agent interface {
	Train(env Env, trainNum int) interface{}
	Policy(env Env) (act ActionEnum)
	String() string
}
