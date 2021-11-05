package common

type Accumulate interface {
	CountAct(act ActionEnum) float64
	Count() float64
	Mean(act ActionEnum) float64
	Add(act ActionEnum, reward float64)
	String() string
}
