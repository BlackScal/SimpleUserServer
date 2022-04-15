package utils

import (
	"math/rand"
	"time"
)

const (
	Alnum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenRandomString(length int) string {

	bytes := []byte(Alnum)
	bytesLen := len(bytes)
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, bytes[random.Intn(bytesLen)])
	}
	return string(result)
}
