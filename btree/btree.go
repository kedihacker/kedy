package main

const key_size = 5

type node struct {
	parent *node
	values []int
	keys   []int
}

func (me *node) split() {
	if len(me.parent.keys) < key_size {
		new_leaf := node{parent: me.parent, values: nil, keys: nil}
		if me.keys == nil {
			new_leaf.values = me.values[3:]
			me.values = me.values[:3]

		}
		if me.keys != nil {

		}

	}
}

func main() {
	//https://github.com/torvalds/linux/blob/5bfc75d92efd494db37f5c4c173d3639d4772966/fs/hfsplus/btree.c

	// root := node{parent: nil, values: nil, keys: nil}

}
