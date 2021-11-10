package agent_offline

import (
	"fmt"
	"gameServer/common"
	"gameServer/envs"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestOfflineSearch(t *testing.T) {
	var (
		banditsNum = 5
		trainNum   = 1000

		mesh = common.NewEncoder([]string{}, []float64{})
		env  = envs.NewBanditsEnv(banditsNum)

		agents = []common.AgentOffline{
			NewSearch(env, mesh, common.SearchMethodEnum_Random),
			NewSearch(env, mesh, common.SearchMethodEnum_MeanQ),
			NewSearch(env, mesh, common.SearchMethodEnum_EpsilonGreed, 0.5),
			NewSearch(env, mesh, common.SearchMethodEnum_SoftMax, 0.5),
			NewSearch(env, mesh, common.SearchMethodEnum_UCB),
		}
	)
	for _, agent := range agents {
		Convey(fmt.Sprintf("TestOfflineSearch:%v", agent), t, func() {
			log.Printf("TestOfflineSearch:%v", agent)
			state, info := env.Reset()
			agent.Reset()
			agent.Train(trainNum)
			act := agent.Policy(state, env.Space())
			log.Printf("act:%v", act)
			log.Print(info)
		})
	}
}
