package envs

import "log"

type GameEnum int
type ActionEnum int

const (
	GameEnum_CartPole_v0 GameEnum = 10001

	ActionEnum_Up    ActionEnum = 20100
	ActionEnum_Down  ActionEnum = 20101
	ActionEnum_Right ActionEnum = 20102
	ActionEnum_Left  ActionEnum = 20103
	ActionEnum_0     ActionEnum = 20200
	ActionEnum_1     ActionEnum = 20201
	ActionEnum_2     ActionEnum = 20202
	ActionEnum_3     ActionEnum = 20203
	ActionEnum_4     ActionEnum = 20204
	ActionEnum_5     ActionEnum = 20205
	ActionEnum_6     ActionEnum = 20206
	ActionEnum_7     ActionEnum = 20207
	ActionEnum_8     ActionEnum = 20208
	ActionEnum_9     ActionEnum = 20209
)

func Game2Str(game GameEnum) string {
	switch game {
	case GameEnum_CartPole_v0:
		return "CartPole_v0"
	default:
		log.Fatalf("unkown game:%v", game)
		return ""
	}
}
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
	case ActionEnum_0:
		return "0"
	case ActionEnum_1:
		return "1"
	case ActionEnum_2:
		return "2"
	case ActionEnum_3:
		return "3"
	case ActionEnum_4:
		return "4"
	case ActionEnum_5:
		return "5"
	case ActionEnum_6:
		return "6"
	case ActionEnum_7:
		return "7"
	case ActionEnum_8:
		return "8"
	case ActionEnum_9:
		return "9"
	default:
		log.Fatalf("unkown %v", act)
		return ""
	}
}
