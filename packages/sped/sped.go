package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	max = 1000
)

func main() {
	start := time.Now()
	w := new(sync.WaitGroup)
	w.Add(1)
	for x := 0; x < max; x++ {
		go func(w *sync.WaitGroup) {
			w.Wait()
		}(w)
	}

	fmt.Println(time.Since(start) / max)
	fmt.Printf("%v\n", runtime.NumGoroutine())
	w.Done()
	fmt.Printf("%v\n", runtime.NumGoroutine())

	start = time.Now()
	w = new(sync.WaitGroup)
	w.Add(1)
	for x := 0; x < max; x++ {
		go func(w *sync.WaitGroup) {
			w.Wait()
		}(w)
	}

	fmt.Println(time.Since(start) / max)
	fmt.Printf("%v\n", runtime.NumGoroutine())
	w.Done()
	fmt.Printf("%v\n", runtime.NumGoroutine())
}
