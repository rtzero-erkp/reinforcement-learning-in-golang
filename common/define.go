package common

import "fmt"

type ActEnum string

const (
	ActionEnum_Up    ActEnum = "Up"
	ActionEnum_Down  ActEnum = "Down"
	ActionEnum_Right ActEnum = "Right"
	ActionEnum_Left  ActEnum = "Left"

	ActionEnum_Card2 ActEnum = "Card2"
	ActionEnum_Card3 ActEnum = "Card3"
	ActionEnum_Card4 ActEnum = "Card4"
	ActionEnum_Card5 ActEnum = "Card5"
	ActionEnum_Card6 ActEnum = "Card6"
	ActionEnum_Card7 ActEnum = "Card7"
	ActionEnum_Card8 ActEnum = "Card8"
	ActionEnum_Card9 ActEnum = "Card9"
	ActionEnum_CardT ActEnum = "CardT"
	ActionEnum_CardJ ActEnum = "CardJ"
	ActionEnum_CardQ ActEnum = "CardQ"
	ActionEnum_CardK ActEnum = "CardK"
	ActionEnum_CardA ActEnum = "CardA"

	ActionEnum_Fold  ActEnum = "Fold"
	ActionEnum_Check ActEnum = "Check"
	ActionEnum_Call  ActEnum = "Call"
	ActionEnum_Bet   ActEnum = "Bet"
	ActionEnum_AllIn ActEnum = "AllIn"
)

func (p ActEnum) String() string {
	return string(p)
}
func (p ActEnum) Same(a ActEnum) bool {
	return string(p) == string(a)
}

const (
	PlayerEnum_unknown = "unknown"
	PlayerEnum_1       = "P1"
	PlayerEnum_2       = "P2"
	PlayerEnum_3       = "P3"
	PlayerEnum_4       = "P4"
	PlayerEnum_5       = "P5"
	PlayerEnum_6       = "P6"
	PlayerEnum_7       = "P7"
	PlayerEnum_8       = "P8"
)

type CardEnum int

const (
	CardEnum_Card2 CardEnum = 1
	CardEnum_Card3 CardEnum = 2
	CardEnum_Card4 CardEnum = 3
	CardEnum_Card5 CardEnum = 4
	CardEnum_Card6 CardEnum = 5
	CardEnum_Card7 CardEnum = 6
	CardEnum_Card8 CardEnum = 7
	CardEnum_Card9 CardEnum = 8
	CardEnum_CardT CardEnum = 9
	CardEnum_CardJ CardEnum = 10
	CardEnum_CardQ CardEnum = 11
	CardEnum_CardK CardEnum = 12
	CardEnum_CardA CardEnum = 13
)

func (p CardEnum) String() string {
	return fmt.Sprint(int(p))
}
func (p CardEnum) Big(a CardEnum) bool {
	return p > a
}
func (p CardEnum) Same(a CardEnum) bool {
	return p == a
}
func (p CardEnum) Small(a CardEnum) bool {
	return p < a
}

type SearchEnum string

const (
	SearchEnum_MC           SearchEnum = "MC"
	SearchEnum_AvgQ         SearchEnum = "AvgQ"
	SearchEnum_EpsilonGreed SearchEnum = "EpsilonGreed"
	SearchEnum_SoftMax      SearchEnum = "SoftMax"
	SearchEnum_UCB          SearchEnum = "UCB"
)

func (p SearchEnum) String() string {
	return string(p)
}

type UpdateEnum string

const (
	UpdateEnum_AvgQ  UpdateEnum = "AvgQ"
	UpdateEnum_DT    UpdateEnum = "DT"
	UpdateEnum_SARSA UpdateEnum = "SARSA"
)

func (p UpdateEnum) String() string {
	return string(p)
}

type ModelEnum string

const (
	NodeEnum_Value  ModelEnum = "Value"
	NodeEnum_Policy ModelEnum = "Policy"
	NodeEnum_Q      ModelEnum = "Q"
	NodeEnum_Cfr    ModelEnum = "Cfr"
)

func (p ModelEnum) String() string {
	return string(p)
}

type UpdateParam struct {
	Model UpdateEnum
	Args  []interface{}
}
type SearchParam struct {
	Model SearchEnum
	Args  []interface{}
}

func NewUpdateParam(mode UpdateEnum, args ...interface{}) *UpdateParam {
	return &UpdateParam{
		Model: mode,
		Args:  args,
	}
}
func NewSearchParam(mode SearchEnum, args ...interface{}) *SearchParam {
	return &SearchParam{
		Model: mode,
		Args:  args,
	}
}
