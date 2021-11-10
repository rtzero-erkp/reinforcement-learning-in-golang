package common

type Agent interface {
	Reset()
	Policy(state Info, space Space) (act ActionEnum)
	Reward(state Info, act ActionEnum, reward float64)
	String() string
}
