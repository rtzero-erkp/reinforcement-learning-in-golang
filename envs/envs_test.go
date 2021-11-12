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
			NewMazeEnv(3, 3),
			NewAKQEnv(3),
		}
	)
	for _, env := range envs {
		Convey(fmt.Sprintf("[TestEnvs0] env:%v", env), t, func() {
			log.Printf("[TestEnvs0] env:%v", env)
			env.Reset()
			step := 0
			for {
				step += 1
				var act = env.Acts().Sample()
				res = env.Step(act)
				if res.Done {
					break
				}
			}
			log.Print(res)
			log.Printf("step:%v", step)
		})
	}
}
