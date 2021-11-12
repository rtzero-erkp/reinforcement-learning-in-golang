package common

import "fmt"

type Info interface {
	Get(key string) interface{}
	Set(key string, val interface{})
	String() string
	Clone() Info
	Clear()
}

var _ Info = &infoMap{}

type infoMap struct {
	m map[string]interface{}
}

func NewInfoMap() Info {
	return &infoMap{
		m: map[string]interface{}{},
	}
}
func (p *infoMap) Get(key string) interface{} {
	return p.m[key]
}
func (p *infoMap) Set(key string, val interface{}) {
	p.m[key] = val
}
func (p *infoMap) String() string {
	var line = "\n"
	for k, v := range p.m {
		line += fmt.Sprintf("[infos] key:%v, val:%v\n", k, v)
	}
	return line
}
func (p *infoMap) Clone() Info {
	var cp = &infoMap{
		m: map[string]interface{}{},
	}
	for k, v := range p.m {
		cp.m[k] = v
	}
	return cp
}
func (p *infoMap) Clear() {
	p.m = map[string]interface{}{}
}
