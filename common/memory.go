package common

//type Memories interface {
//	Add(state State, act ActionEnum, reward float64)
//	Clear()
//	Get() []*Memory
//}

type Memory struct {
	State  State
	Act    ActionEnum
	Reward float64
}
