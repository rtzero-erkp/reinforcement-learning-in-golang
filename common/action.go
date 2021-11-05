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
		return "[Act] Up"
	case ActionEnum_Down:
		return "[Act] Down"
	case ActionEnum_Right:
		return "[Act] Right"
	case ActionEnum_Left:
		return "[Act] Left"
	default:
		return fmt.Sprintf("[Act] %v", int(p))
	}
}
