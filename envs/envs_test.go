package envs

import (
	"fmt"
	"gameServer/common"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestEnvs0(t *testing.T) {
	var (
		res  *common.Result
		envs = []common.Env{
			NewBanditsEnv(5),
			NewCartPoleEnv(2.4, 12),
			NewAKQEnv(3),
		}
	)
	for _, env := range envs {
		Convey(fmt.Sprintf("TestEnvs0:%v", env), t, func() {
			log.Printf("TestEnvs0:%v", env)
			env.Reset()
			for {
				var act = env.ActionSpace().Sample()
				res = env.Step(act)
				if res.Done {
					break
				}
			}
			log.Println(res)
		})
	}
}
