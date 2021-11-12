package envs

import (
	"fmt"
	"gameServer/common"
	"log"
)

var _ common.Env = &MazeEnv{}

type MazeEnv struct {
	state common.Info
	info  common.Info
	acts  common.Acts
}

func NewMazeEnv(deep int, expand int) common.Env {
	o := &MazeEnv{
		state: common.NewInfoMap(),
		info:  common.NewInfoMap(),
		acts:  common.NewActsVec(),
	}
	o.info.Set("deep", deep)
	o.info.Set("expand", expand)
	return o
}

func (p *MazeEnv) String() string {
	return fmt.Sprintf("Maze")
}
func (p *MazeEnv) Clone() common.Env {
	var cp = &MazeEnv{
		state: p.state.Clone(),
		info:  p.info.Clone(),
		acts:  p.acts.Clone(),
	}
	return cp
}
func (p *MazeEnv) Acts() common.Acts {
	return p.acts
}
func (p *MazeEnv) State() common.Info {
	return p.state.Clone()
}
func (p *MazeEnv) Reset() (common.Info, common.Info) {
	var start = common.ActEnum("0")
	var end = start
	deep := p.info.Get("deep").(int)
	for i := 0; i < deep; i++ {
		end += ".0"
	}

	p.info.Set("step", 0.0)
	p.info.Set("start", start)
	p.info.Set("end", end)

	p.state.Clear()
	p.state.Set("pt", start)

	p.gen(start)

	return p.state.Clone(), p.info.Clone()
}
func (p *MazeEnv) Step(act common.ActEnum) (res *common.Result) {
	if !p.acts.Contain(act) {
		log.Fatal(fmt.Sprintf("acts:%v not contain act:%v", p.acts, act))
	}

	p.state.Clear()
	p.state.Set("pt", act)
	end := p.info.Get("end").(common.ActEnum)
	step := p.info.Get("step").(float64)

	res = &common.Result{
		Reward: []float64{0},
		Done:   end == act,
	}
	if res.Done {
		res.Reward[0] = 1.0 / step
		p.acts.Clear()
	} else {
		p.info.Set("step", step+1)
		p.gen(act)
	}
	res.State = p.state.Clone()
	res.Info = p.info.Clone()

	return res
}
func (p *MazeEnv) gen(act common.ActEnum) {
	expand := p.info.Get("expand").(int)
	deep := p.info.Get("deep").(int)

	var actStrLen = len(act)
	var last = -1
	var count = -1
	for i := 0; i < actStrLen; i++ {
		if act[actStrLen-1-i] == '.' {
			if last == -1 {
				last = i
			}
			count += 1
		}
	}

	p.acts.Clear()
	if last != -1 {
		p.acts.AddEnum(act[:actStrLen-1-last])
	}
	if count+1 < deep {
		for actI := 0; actI < expand; actI++ {
			p.acts.AddEnum(common.ActEnum(fmt.Sprintf("%v.%v", act, actI)))
		}
	}
}
