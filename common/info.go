package common

import "fmt"

type Info interface {
	Hash(mesh *Mesh) string
	Get(key string) float64
	Set(key string, val float64)
	Add(key string, val float64)
	String() string
	Keys() []string
	Clone() Info
}

var _ Info = &infoMap{}

type infoMap map[string]float64

func (p *infoMap) Clone() Info {
	var cp = &infoMap{}
	for k, v := range *p {
		(*cp)[k] = v
	}
	return cp
}

func NewInfoMap() Info {
	return &infoMap{}
}
func (p *infoMap) Get(key string) float64 {
	return (*p)[key]
}
func (p *infoMap) Keys() []string {
	return (*p).Keys()
}
func (p *infoMap) Set(key string, val float64) {
	(*p)[key] = val
}
func (p *infoMap) Add(key string, val float64) {
	(*p)[key] += val
}
func (p *infoMap) String() string {
	var line = "\n"
	for k, v := range *p {
		line += fmt.Sprintf("[infos] key:%v, val:%10.7f\n", k, v)
	}
	return line
}
func (p *infoMap) Hash(mesh *Mesh) string {
	// switch type
	var hash = ""
	for i, key := range mesh.keys {
		var v0 = mesh.vals[i]
		var v1 = (*p)[key]
		hash += fmt.Sprint(int(v0*v1)) + " "
	}
	return hash
}

type Mesh struct {
	keys []string
	vals []float64
}

func NewMesh(keys []string, vals []float64) *Mesh {
	var o = &Mesh{
		keys: keys,
		vals: vals,
	}
	return o
}
