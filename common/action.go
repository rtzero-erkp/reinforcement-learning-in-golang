package common

import (
	"fmt"
)

type ActionEnum int

const (
	ActionEnum_Unknown ActionEnum = -1

	ActionEnum_Up    ActionEnum = 20100
	ActionEnum_Down  ActionEnum = 20101
	ActionEnum_Right ActionEnum = 20102
	ActionEnum_Left  ActionEnum = 20103
)

func (p ActionEnum) String() string {
	switch p {
	case ActionEnum_Up:
		return "U"
	case ActionEnum_Down:
		return "D"
	case ActionEnum_Right:
		return "R"
	case ActionEnum_Left:
		return "L"
	default:
		return fmt.Sprintf("%v", int(p))
	}
}
