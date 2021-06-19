package main

import "github.com/google/btree"


type btree.Item interface {
	Less(than Item) bool
}

type value struct {
	value int
}

func (n value) Less(than value) bool {
	return n.value < than.value
}

func main() {

	tree := btree.New(3)

	var a value

	a.value = 12

	tree.ReplaceOrInsert(a)
}
