package common

import (
	"fmt"
)

type Agent interface {
	Train(env Env, trainNum int) interface{}
	Policy(env Env) (act ActEnum)
	String() string
}

type Code struct {
	code interface{}
}
type Encoder interface {
	Hash(info Info) Code
}
type EncodeFloat64 struct {
	keys []string
	vals []float64
}
type EncodeString struct {
	keys []string
}

var _ Encoder = &EncodeFloat64{}

func (p Code) String() string {
	return fmt.Sprintf("Code:%v", p.code)
}

func NewEncoderFloat64(keys []string, vals []float64) Encoder {
	var o = &EncodeFloat64{
		keys: keys,
		vals: vals,
	}
	return o
}
func (p *EncodeFloat64) Hash(info Info) Code {
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
	return Code{code: hash}
}

func NewEncoderString(keys []string) Encoder {
	var o = &EncodeString{
		keys: keys,
	}
	return o
}
func (p *EncodeString) Hash(info Info) Code {
	//log.Printf("info:%v", info)
	var hash = ""
	if p.keys != nil {
		for _, key := range p.keys {
			//log.Printf("key:%v", key)
			var val = info.Get(key)
			//log.Printf("val:%v, %T", val, val)
			if val != nil {
				hash += fmt.Sprintf("%v ", val)
			} else {
				hash += "0 "
			}
		}
	}
	//log.Printf("hash:%v", hash)
	return Code{code: hash}
}
