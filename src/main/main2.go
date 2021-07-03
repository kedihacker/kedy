package main

import (
	"fmt"
	"regexp"
)

func main() {
	help := regexp.MustCompile("12")
	fmt.Printf("%+v\n", help)
}
