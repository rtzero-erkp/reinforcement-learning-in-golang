package common

type Result struct {
	State  Info
	Reward []float64
	Done   bool
	Info   Info
}

type Env interface {
	String() string                    // 打印
	ActionSpace() Space                // 行动空间
	Step(act ActionEnum) (res *Result) // 执行一步
	Reset() Info                       // 重置游戏
	Clone() Env                        // 复制游戏
	Set(state Info)                    // 设置为目标状态
}
