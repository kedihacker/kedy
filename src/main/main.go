package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randguid [16]byte
	start := time.Now()
	rand := rand.NewSource(0)
	for x := 0; x < 1000; x++ {
		rand.Read(randguid[:])

		fmt.Printf("%x\n", randguid)
	}
	fmt.Print(time.Since(start))
}
