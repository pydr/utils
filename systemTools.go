package utils

import (
	"net"
	"net/http"
	"time"
)

// 获取本机ip
func GetLocalIp() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	return
}

// 请求重试工具
func Request(client *http.Client, req *http.Request, retry int) (resp *http.Response, err error) {
	reqTimes := 1
	for {
		resp, err = client.Do(req)
		if err != nil {
			if reqTimes > retry {
				return
			}
			reqTimes++
			time.Sleep(time.Duration(2*reqTimes) * time.Second)
			continue
		}
		return
	}
}
