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
		update_AvgQ         = common.NewUpdateParam(common.UpdateEnum_AvgQ)
		search_MC           = common.NewSearchParam(common.SearchEnum_MC)
		search_AvgQ         = common.NewSearchParam(common.SearchEnum_AvgQ)
		search_EpsilonGreed = common.NewSearchParam(common.SearchEnum_EpsilonGreed, 0.5)
		search_SoftMax      = common.NewSearchParam(common.SearchEnum_SoftMax, 0.5)
		search_UCB          = common.NewSearchParam(common.SearchEnum_UCB)
		modelTree           = common.NewModelTree(common.NodeEnum_Q, update_AvgQ)
		agents              = []common.Agent{
			NewModelTree(modelTree, search_MC),
			NewModelTree(modelTree, search_AvgQ),
			NewModelTree(modelTree, search_EpsilonGreed),
			NewModelTree(modelTree, search_SoftMax),
			NewModelTree(modelTree, search_UCB),
		}
	)
	for _, agent := range agents {
		Convey(fmt.Sprintf("[Test_model_tree_0] env:%v, agent:%v", env, agent), t, func() {
			log.Printf("[Test_model_tree_0] env:%v, agent:%v", env, agent)
			_, info := env.Reset()
			agent.Train(env, 200)
			act := agent.Policy(env)

			var val = info.Get(fmt.Sprintf("ex%v", act)).(float64)
			for _, actI := range common.NewSpaceVecByNum(5).Acts() {
				var valI = info.Get(fmt.Sprintf("ex%v", actI)).(float64)
				if (actI != act) && (valI > val) {
					t.Fatalf("act:%v is not best", act)
				}
			}
			log.Printf("act:%v is best", act)
		})
	}
}