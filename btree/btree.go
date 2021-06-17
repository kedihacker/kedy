package main

import "sync"


type node struct {
	keys []int
	values []*node
}
type Btree struct {
	root []*node
	lock *sync.Locker
}

func (n *node) split() (old *node, new *node) {
	new_node := new(node)
	new_node.keys , n.keys = n.keys[]

	return
}
