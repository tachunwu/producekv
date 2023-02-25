package producekv

import "fmt"

func GeneratePrefixKeyValuePairs(prefix []byte, count int) [][]byte {
	kvPairs := make([][]byte, count)

	for i := 0; i < count; i++ {
		key := make([]byte, len(prefix)+4)
		copy(key, prefix)
		copy(key[len(prefix):], []byte(fmt.Sprintf("%04d", i)))
		value := GenerateRandomBytes(20)
		kvPairs[i] = append(key, value...)
	}

	return kvPairs
}
