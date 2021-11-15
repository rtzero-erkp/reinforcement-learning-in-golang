package agent_model

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func Test_model_tree_0(t *testing.T) {
	var (
		env                 = envs.NewBanditsEnv(5)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelTree           = common.NewModelTree(common.NodeEnum_Q, update_AvgQ)
		agents              = []common.Agent{
			NewModelTree(modelTree, search_MC, update_AvgQ),
			NewModelTree(modelTree, search_EpsilonGreed, update_AvgQ),
			NewModelTree(modelTree, search_SoftMax, update_AvgQ),
			NewModelTree(modelTree, search_UCB, update_AvgQ),
		}
	)
	Convey(fmt.Sprintf("[Test_model_tree_0] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			_, info := env.Reset()
			//log.Printf("env:%v", info)

			modelTree.Reset()
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

func Test_model_tree_1(t *testing.T) {
	const rewardLimit = 3000.0
	var (
		env                 = envs.NewCartPoleEnv(2.4, 12)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelTree           = common.NewModelTree(common.NodeEnum_Q, update_AvgQ)
		agents              = []common.Agent{
			NewModelTree(modelTree, search_MC, update_AvgQ),
			NewModelTree(modelTree, search_EpsilonGreed, update_AvgQ),
			NewModelTree(modelTree, search_SoftMax, update_AvgQ),
			NewModelTree(modelTree, search_UCB, update_AvgQ),
		}
	)
	Convey(fmt.Sprintf("[Test_model_tree_1] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			env.Reset()
			modelTree.Reset()
			reward := 0.0
			for {
				agent.Train(env, 200)
				act := agent.Policy(env)
				res := env.Step(act)
				reward += res.Reward[0]
				if res.Done {
					break
				}
				if reward > rewardLimit {
					log.Printf("reward:%v > rewardLimit:%v", reward, rewardLimit)
					break
				}
			}
			log.Printf("reward:%v", reward)
		}
	})
}

func Test_model_tree_2(t *testing.T) {
	var (
		// 不推荐使用Tree方式, 导致内存不可控
		env                 = envs.NewMazeEnv(3, 3)
		update_AvgQ         = common.NewUpdateMethod(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchMethod(common.SearchEnum_MC)
		search_EpsilonGreed = common.NewSearchMethod(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchMethod(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchMethod(common.SearchEnum_UCB)
		modelTree           = common.NewModelTree(common.NodeEnum_Q, update_AvgQ)
		agents              = []common.Agent{
			NewModelTree(modelTree, search_MC, update_AvgQ),
			NewModelTree(modelTree, search_EpsilonGreed, update_AvgQ),
			NewModelTree(modelTree, search_SoftMax, update_AvgQ),
			NewModelTree(modelTree, search_UCB, update_AvgQ),
		}
	)
	Convey(fmt.Sprintf("[Test_model_tree_2] env:%v", env), t, func() {
		for _, agent := range agents {
			log.Printf("agent:%v", agent)
			env.Reset()
			modelTree.Reset()
			reward := 0.0
			step := 0.0
			for {
				agent.Train(env, 200)
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
