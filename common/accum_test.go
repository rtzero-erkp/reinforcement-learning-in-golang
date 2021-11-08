package common

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"math/rand"
	"testing"
)

func TestAccum0(t *testing.T) {
	Convey(fmt.Sprintf("TestAccum0:"), t, func() {
		t.Logf("TestAccum0:")
		var rate = rand.Float64()
		log.Printf("rate:%v", rate)
		var accum = NewAccum()
		for i := 0; i < 100; i++ {
			accum.Add(ActionEnum_Right, rate)
		}
		log.Println(accum)
	})
}
