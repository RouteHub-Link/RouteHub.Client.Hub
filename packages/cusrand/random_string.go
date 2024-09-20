package cusrand

import (
	"crypto/rand"
	"math/big"
)

func UniqueRandomString(count int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, count)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err) // handle error appropriately in real applications
		}
		b[i] = charset[num.Int64()]
	}
	return string(b)
}
