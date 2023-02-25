package main

import (
	"fmt"

	"github.com/tachunwu/producekv"
)

func main() {

	key, value := producekv.GenerateRandomKeyValuePair()
	fmt.Println(key, value)

	prefix := []byte("/table")
	kvPairs := producekv.GeneratePrefixKeyValuePairs(prefix, 5)
	for _, kv := range kvPairs {
		fmt.Println(string(kv))
	}
}
