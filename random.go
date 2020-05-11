package utils

import (
	"math/rand"
	"time"
)

const (
	randomSeed    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomSeedLen = len(randomSeed)
)

func RandomBytes(l int) []byte {
	result := make([]byte, 0, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, randomSeed[r.Intn(randomSeedLen)])
	}
	return result
}

// 生成区间随机数的随机数
func GenNonce(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
