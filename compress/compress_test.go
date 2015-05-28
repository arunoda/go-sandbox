package main

import (
	"github.com/golang/snappy/snappy"
	"math/rand"
	"testing"
)

func BenchmarkCompressSnappy(b *testing.B) {
	originalData := make([]byte, 3600*24*200)
	for lc, _ := range originalData {
		randValue := rand.Intn(100)
		originalData[lc] = byte(randValue)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snappy.Encode(nil, originalData)
	}
}
