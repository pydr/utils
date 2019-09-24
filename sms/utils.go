package sms

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// 生成随机数
func genNonce(limit int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return strconv.Itoa(r.Intn(limit))
}

// 生成签名
func makeSign(secretKey, nonce, timeStamp, mobile string) string {
	str := "appkey=" + secretKey + "&random=" + nonce + "&time=" + timeStamp + "&mobile=" + mobile
	h := sha256.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}
