package common

import "fmt"

type Stater interface {
	Hash(mesh *Mesh) string
	Hash2(mesh *Mesh) []int
	Clone() State
	String() string
}

var _ Stater = &State{}

type State []float64
type Mesh []float64

func (p *State) Hash(mesh *Mesh) string {
	var hash = ""
	for i, v := range *p {
		hash += fmt.Sprint(int(v * (*mesh)[i]))
	}
	return hash
}
func (p *State) Hash2(mesh *Mesh) []int {
	var hash []int
	for i, v := range *p {
		hash = append(hash, int(v*(*mesh)[i]))
	}
	return hash
}
func (p *State) Clone() State {
	var cp = State{}
	for _, v := range *p {
		cp = append(cp, v)
	}
	return cp
}
func (p *State) String() string {
	var line = "[stateVec]"
	for _, v := range *p {
		line += fmt.Sprintf(" %10.7f", v)
	}
	return line
}

func NewState(vec []float64) Stater {
	var o = &State{}
	for _, v := range vec {
		*o = append(*o, v)
	}
	return o
}
func NewMesh(vec ...float64) *Mesh {
	var o = &Mesh{}
	for _, v := range vec {
		*o = append(*o, v)
	}
	return o
}
