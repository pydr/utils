package sms

import "net/http"

const APIBASE = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?"

type Client struct {
	appid      string
	secretKey  string
	httpClient *http.Client
}

func NewSMSClient(appid, secretKey string) *Client {
	return &Client{
		appid:      appid,
		secretKey:  secretKey,
		httpClient: &http.Client{},
	}
}
