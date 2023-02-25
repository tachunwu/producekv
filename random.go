package producekv

import (
	cryptorand "crypto/rand"
	"math/rand"
	"time"
)

func GenerateRandomBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	cryptorand.Read(result)
	return result
}

func GenerateRandomKeyValuePair() (key []byte, value []byte) {
	key = GenerateRandomBytes(10)
	value = GenerateRandomBytes(20)
	return key, value
}
