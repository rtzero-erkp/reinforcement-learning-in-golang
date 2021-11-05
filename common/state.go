package common

import "fmt"

type State []float64

func (p State) Encode(mesh []float64) []int {
	var code []int
	for i, v := range p {
		code = append(code, int(v*mesh[i]))
	}
	return code
}
func (p State) String() string {
	var line = "[stateVec]"
	for _, v := range p {
		line += fmt.Sprintf(" %10.7f", v)
	}
	return line
}
