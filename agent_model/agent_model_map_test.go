package agent_model

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	"gameServer/utils"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func Test_model_map_0(t *testing.T) {
	var (
		encoder             = common.NewEncoderFloat64([]string{}, []float64{})
		env                 = envs.NewBanditsEnv(5)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelMap            = common.NewModelMap(common.NodeEnum_Q)
		agents              = []common.Agent{
			NewModelMap(modelMap, search_MC, update_AvgQ, encoder),
			NewModelMap(modelMap, search_EpsilonGreed, update_AvgQ, encoder),
			NewModelMap(modelMap, search_SoftMax, update_AvgQ, encoder),
			NewModelMap(modelMap, search_UCB, update_AvgQ, encoder),
		}
	)
	Convey(fmt.Sprintf("[Test_model_map_0] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			_, info := env.Reset()
			//log.Printf("env:%v", info)

			modelMap.Reset()
			agent.Train(env, 200)
			act := agent.Policy(env)

			var val = info.Get(fmt.Sprintf("ex%v", act)).(float64)
			for _, actI := range common.NewActsVecByNum(5).All() {
				var valI = info.Get(fmt.Sprintf("ex%v", actI)).(float64)
				if (actI != act) && (valI > val) {
					t.Fatalf("act:%v is not best", act)
				}
			}
			log.Printf("act:%v is best", act)
		}
	})
}

func Test_model_map_1(t *testing.T) {
	const rewardLimit = 3000.0
	var (
		encoder = common.NewEncoderFloat64(
			[]string{"x", "xDot", "theta", "thetaDot"},
			[]float64{50 / 2.4, 20, 50 / 12, 20})
		env                 = envs.NewCartPoleEnv(2.4, 12)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelMap            = common.NewModelMap(common.NodeEnum_Q)
		agents              = []common.Agent{
			NewModelMap(modelMap, search_MC, update_AvgQ, encoder),
			NewModelMap(modelMap, search_EpsilonGreed, update_AvgQ, encoder),
			NewModelMap(modelMap, search_SoftMax, update_AvgQ, encoder),
			NewModelMap(modelMap, search_UCB, update_AvgQ, encoder),
		}
	)
	Convey(fmt.Sprintf("[Test_model_map_1] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			env.Reset()
			modelMap.Reset()
			reward := 0.0
			step := 0.0
			for {
				utils.bar(step, rewardLimit, "reward")
				//agent.Train(env, 200)
				agent.Train(env, 50)
				//agent.Train(env, 20)
				act := agent.Policy(env)
				//log.Printf("state:%v, act:%v", env.State(), act)
				res := env.Step(act)
				reward += res.Reward[0]
				step += 1
				if res.Done {
					break
				}
				if reward > rewardLimit {
					log.Printf("reward:%v > rewardLimit:%v", reward, rewardLimit)
					break
				}
			}
			log.Printf("reward:%v, step:%v", reward, step)
		}
	})
}

func Test_model_map_2(t *testing.T) {
	var (
		encoder             = common.NewEncoderString([]string{"pt"})
		env                 = envs.NewMazeEnv(3, 3)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelMap            = common.NewModelMap(common.NodeEnum_Q)
		agents              = []common.Agent{
			NewModelMap(modelMap, search_MC, update_AvgQ, encoder),
			NewModelMap(modelMap, search_EpsilonGreed, update_AvgQ, encoder),
			NewModelMap(modelMap, search_SoftMax, update_AvgQ, encoder),
			NewModelMap(modelMap, search_UCB, update_AvgQ, encoder),
		}
	)
	Convey(fmt.Sprintf("[Test_model_map_2] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			env.Reset()
			modelMap.Reset()
			reward := 0.0
			step := 0.0
			for {
				agent.Train(env, 200)
				//agent.Train(env, 10)
				//agent.Train(env, 2)
				break
				act := agent.Policy(env)
				//log.Printf("state:%v, act:%v", env.State(), act)
				res := env.Step(act)
				reward += res.Reward[0]
				step += 1
				if res.Done {
					break
				}
			}
			log.Printf("reward:%v, step:%v", reward, step)
		}
	})
}
