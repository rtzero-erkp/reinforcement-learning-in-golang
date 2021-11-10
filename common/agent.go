package common

type AgentOnline interface {
	Reset()
	Policy(state Info, space Space) (act ActionEnum)
	Reward(state Info, act ActionEnum, reward float64)
	String() string
}

type AgentOffline interface {
	Reset()
	Train(trainNum int)
	Policy(state Info, space Space) (act ActionEnum)
	String() string
}
