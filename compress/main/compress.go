package main

import (
	"fmt"
	"github.com/golang/snappy/snappy"
	// "io/ioutil"
	"math/rand"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	originalData := make([]byte, 3600*24*200)
	for lc, _ := range originalData {
		randValue := rand.Intn(100)
		// randValue =
		originalData[lc] = byte(randValue)
	}

	compressed, err := snappy.Encode(nil, originalData)
	check(err)

	fmt.Println("Ori:", len(originalData), " - Dest:", len(compressed))
	fmt.Println("Ratio:", float64(len(originalData))/float64(len(compressed)))
}
