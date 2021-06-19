package bptree

type node struct {
	items []int
	child []*node
}

func (n *node) splt(i int) (Item, *node) {
	item := n.items[i]
	next := n.cow.newNode()
	next.items = append(next.items, n.items[i+1:]...) //add items to the irght
	n.items.truncate(i)
	if len(n.children) > 0 {
		next.children = append(next.children, n.children[i+1:]...)
		n.children.truncate(i + 1)
	}
	return item, next
}

func (n *node) split() *node {
	new := node{}                                             //new
	split_index := len(n.items) / 2                           //set index
	new.items = append(new.items, n.items[split_index+1:]...) //copy items
	new_array := make([]int, 5)                               //new array
	copy(new_array, n.items[split_index+1:])                  //turnucatw array
	if len(n.child) > 0 {                                     //if isnt leaf

	}
	return &new

}
