package xf

import (
	"net/http"
	"net"
)

func GetClientIP(req *http.Request) string  {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

