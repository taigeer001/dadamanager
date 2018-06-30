package admin

import (
	"net/http"
	"time"
	"../../../framework"
)

type RequestWallUnit struct {
	N int
	Datetime int64
}

var rwus = make(map[string]RequestWallUnit)

func ClearRuquest()  {
	for k, v := range rwus{
		if time.Now().Unix()-v.Datetime >= 300 {
			delete(rwus, k)
		}
	}
}

func RequestOK(req *http.Request) bool {
	return RequestDomianOK("", req)
}

func RequestDomianOK(name string, req *http.Request) bool {
	cip := xf.GetClientIP(req)+"#"+name
	if v, ok := rwus[cip]; ok {
		rwus[cip] = RequestWallUnit{N: v.N, Datetime:time.Now().Unix()}
		return v.N <3
	}
	return true
}

func RequestUpdate(req *http.Request) {
	RequestDomianUpdate("#", req)
}

func RequestDomianUpdate(name string, req *http.Request) {
	cip := xf.GetClientIP(req)+"#"+name
	if v, ok := rwus[cip]; ok {
		rwus[cip] = RequestWallUnit{N: v.N+1, Datetime:time.Now().Unix()}
	} else {
		rwus[cip] = RequestWallUnit{N: 1, Datetime:time.Now().Unix()}
	}
}
