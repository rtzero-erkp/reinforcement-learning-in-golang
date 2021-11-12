package envs

import (
	"fmt"
	"gameServer/common"
	"log"
	"math/rand"
)

/* DESC:
 *  3 card deck: Ace > King > Queen
 *  2 Players, each player gets one card face+down
 *  Higher card wins
 *  Betting: (half+street game)
 *    Ante: 1 chip
 *    Player 1: bet P, or check
 *    Player 2: call or fold
 *  Stakes: scheduling signup order by chip count
 *  -------------------------------------------------------
 *  E(P0) |  P1  |   A  |   A  |   K  |   K  |   Q  |   Q
 *  ------|------|------|------|------|------|------|------
 *    P0  |      | Bet  | Fold | Bet  | Fold | Bet  | Fold
 *  ------|------|------|------|------|------|------|------
 *    A   | Call |      |      | -1-P |  -1  | -1-P |  -1
 *    A   | Fold |      |      |  +1  |  +1  |  +1  |  +1
 *    K   | Call | +1+P |  -1  |      |      | -1-P |  -1
 *    K   | Fold |  +1  |  +1  |      |      |  +1  |  +1
 *    Q   | Call | +1+P |  -1  | +1+P |  -1  |      |
 *    Q   | Fold |  +1  |  +1  |  +1  |  +1  |      |
 *  -------------------------------------------------------
 */

type AKQEnv struct {
	state common.Info // 玩家状态
	info  common.Info // 游戏信息
	acts  common.Acts // 行动空间
}

func NewAKQEnv(P float64) common.Env {
	o := &AKQEnv{
		state: common.NewInfoMap(),
		info:  common.NewInfoMap(),
		acts:  common.NewActsVecByEnum(),
	}

	o.info.Set("P", P)
	return o
}

func (p *AKQEnv) String() string {
	return "AKQ"
}
func (p *AKQEnv) Clone() common.Env {
	var cp = &AKQEnv{
		state: p.state.Clone(),
		info:  p.info.Clone(),
		acts:  p.acts.Clone(),
	}

	return cp
}
func (p *AKQEnv) Acts() common.Acts {
	return p.acts
}
func (p *AKQEnv) State() common.Info {
	return p.state
}
func (p *AKQEnv) Reset() (common.Info, common.Info) {
	var dealt = []common.CardEnum{common.CardEnum_CardA, common.CardEnum_CardK, common.CardEnum_CardQ}
	for i := 0; i < len(dealt); i++ {
		var j = rand.Intn(len(dealt))
		var tmp = dealt[i]
		dealt[i] = dealt[j]
		dealt[j] = tmp
	}
	p.state.Set(common.PlayerEnum_1, dealt[0])
	p.state.Set(common.PlayerEnum_2, dealt[1])
	p.state.Set("crt", common.PlayerEnum_1)
	p.acts.Clear()
	p.acts.AddEnum(common.ActionEnum_Bet, common.ActionEnum_Fold)
	return p.state, p.info
}
func (p *AKQEnv) Step(act common.ActEnum) (res *common.Result) {
	if !p.acts.Contain(act) {
		log.Fatal(fmt.Sprintf("actions acts not contain act:%v", act))
	}

	var player = p.state.Get("crt").(string)
	var P = p.info.Get("P").(float64)
	var anti float64 = 1
	p.state.Set(player+":act", act)

	res = &common.Result{State: p.state, Info: p.info}
	if player == common.PlayerEnum_1 {
		if act == common.ActionEnum_Fold {
			res.Done = true
			res.Reward = []float64{-anti, +anti}
			p.state.Set("crt", common.PlayerEnum_unknown)
			return
		} else
		if act == common.ActionEnum_Bet {
			p.acts.Clear()
			p.acts.AddEnum(common.ActionEnum_Fold, common.ActionEnum_Call)
			p.state.Set("crt", common.PlayerEnum_2)
			res.Done = false
			res.Reward = []float64{}
			return
		}
	}
	if player == common.PlayerEnum_2 {
		if act == common.ActionEnum_Fold {
			res.Done = true
			res.Reward = []float64{+anti, -anti}
			p.state.Set("crt", common.PlayerEnum_unknown)
			return
		} else
		if act == common.ActionEnum_Call {
			var dealt1 = p.state.Get(common.PlayerEnum_1).(common.CardEnum)
			var dealt2 = p.state.Get(common.PlayerEnum_2).(common.CardEnum)
			if dealt1.Big(dealt2) {
				res.Reward = []float64{+anti + P, -anti - P}
			} else
			if dealt1.Same(dealt2) {
				res.Reward = []float64{0, 0}
			} else {
				res.Reward = []float64{-anti - P, +anti + P}
			}
			res.Done = true
			p.state.Set("crt", common.PlayerEnum_unknown)
			return
		}
	}
	log.Fatal("unknown error")
	return
}
