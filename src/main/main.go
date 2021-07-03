package main

import (
	"fmt"

	"github.com/google/btree"
)

// type item interface {
// 	Less(than Item) bool
// }

// type value btree.Item

type Help struct {
	value int
}

func (n Help) Less(than btree.Item) bool {
	if h, ok := than.(Help); ok {
		return n.value < h.value
	}
	// I've no idea what you'd do in a default case here...
	return n.value < than.(Help).value
}

func main() {
	tree := btree.New(5)
	var keyz btree.Item
	// x := 12
	for x := 0; x < 1000000; x++ {
		// fmt.Printf("%v\n", x)
		keyz = Help{value: 12}
		tree.ReplaceOrInsert(keyz)

	}

	// tree.ReplaceOrInsert(keyz)
	fmt.Printf("%+v", tree)
}
