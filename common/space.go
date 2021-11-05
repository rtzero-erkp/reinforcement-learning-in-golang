package common

type Space interface {
	Contain(act ActionEnum) bool
	Acts() []ActionEnum
	Sample() ActionEnum
}
