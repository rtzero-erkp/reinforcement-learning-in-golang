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

func Action2Str(act ActionEnum) string {
	switch act {
	case ActionEnum_Up:
		return "Up"
	case ActionEnum_Down:
		return "Down"
	case ActionEnum_Right:
		return "Right"
	case ActionEnum_Left:
		return "Left"
	default:
		return fmt.Sprintf("%v", int(act))
	}
}
