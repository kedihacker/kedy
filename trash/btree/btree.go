package main

import (
	"sync"
)

type node struct {
	keys   []int
	values []*node
}
type Btree struct {
	root []*node
	lock *sync.Locker
}

func (n *node) split(yank bool) (old *node, new *node) {
	new_node := node{
		keys:   []int{},
		values: []*node{},
	}
	item := n.keys[len(n.keys)/2]

	return n, *new_node
}
