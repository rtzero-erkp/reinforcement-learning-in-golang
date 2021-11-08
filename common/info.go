package common

import "fmt"

type Info interface {
	Get(key string) float64
	Set(key string, val float64)
	Add(key string, val float64)
	String() string
	Keys() []string
	Clone() Info
}

var _ Info = &InfoMap{}

type InfoMap map[string]float64

func (p *InfoMap) Clone() Info {
	var cp = &InfoMap{}
	for k, v := range *p {
		(*cp)[k] = v
	}
	return cp
}

func NewInfoMap() Info {
	return &InfoMap{}
}
func (p *InfoMap) Get(key string) float64 {
	return (*p)[key]
}
func (p *InfoMap) Keys() []string {
	return (*p).Keys()
}
func (p *InfoMap) Set(key string, val float64) {
	(*p)[key] = val
}
func (p *InfoMap) Add(key string, val float64) {
	(*p)[key] += val
}
func (p *InfoMap) String() string {
	var line = "\n"
	for k, v := range *p {
		line += fmt.Sprintf("[infos] key:%v, val:%10.7f\n", k, v)
	}
	return line
}
