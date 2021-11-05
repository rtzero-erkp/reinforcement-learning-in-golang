package common

type Policy interface {
	Sample() (act ActionEnum)
	Set(act ActionEnum, weight float64)
	Clean()
}
