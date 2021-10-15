package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	hash1, err := blake2b.New512([]byte("help"))
	if err != nil {
		panic(err)
	}
	hash1.Write([]byte("help2"))
	fmt.Printf("%x\n", hash1.Sum([]byte("")))
	fmt.Printf("%x\n", len([]byte("")))
}
