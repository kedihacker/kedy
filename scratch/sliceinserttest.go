package main

import "fmt"

func main() {
	child := []string{"a", "b", "c"}
	index := 1
	newchar := "asd"
	child = append(child[:index+1], child[index:]...)
	child[index] = newchar
	fmt.Print(child)
}
