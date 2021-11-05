package common

type Reward interface {
	Mean(act ActionEnum) float64
	Add(act ActionEnum, reward float64)
	String() string
}
