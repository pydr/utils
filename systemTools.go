package utils

import (
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
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

// 定时器
func Timer(option string, f func(string)) {
	go func() {
		for {
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			logs.Warn("下次执行时间", next.Sub(now))

			<-t.C
			f(option)
		}
	}()
}

// 生成区间随机数的随机数
func GenNonce(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(max-min) + min
}
