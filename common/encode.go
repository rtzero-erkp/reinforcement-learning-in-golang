package common

import "fmt"

type Encoder interface {
	Hash(info Info) string
}

var _ Encoder = &Encode{}

type Encode struct {
	keys []string
	vals []float64
}

func NewEncoder(keys []string, vals []float64) Encoder {
	var o = &Encode{
		keys: keys,
		vals: vals,
	}
	return o
}
func (p *Encode) Hash(info Info) string {
	var hash = ""
	if p.keys != nil {
		for i, key := range p.keys {
			var v0 = p.vals[i]
			var v1 = info.Get(key)
			if v1 != nil {
				hash += fmt.Sprint(int(v0*v1.(float64))) + " "
			} else {
				hash += "0 "
			}
		}
	}
	return hash
}
