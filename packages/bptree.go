package bptree

type node struct {
	items []int
	child []*node
}

// func (n *node) splt(i int) (Item, *node) {
// 	item := n.items[i]
// 	next := n.cow.newNode()
// 	next.items = append(next.items, n.items[i+1:]...) //add items to the irght
// 	n.items.truncate(i)
// 	if len(n.children) > 0 {
// 		next.children = append(next.children, n.children[i+1:]...)
// 		n.children.truncate(i + 1)
// 	}
// 	return item, next
// }

func (n *node) split() (int, *node) {
	new := node{} //new
	split_index := len(n.items) / 2
	item := n.items[split_index]                              //set index
	new.items = append(new.items, n.items[split_index+1:]...) //copy items
	new_array := make([]int, 5)                               //new array
	copy(new_array, n.items[:split_index])                    //turnucatw array
	n.items = new_array                                       //ASSİNG array

	if len(n.child) > 0 { //if isnt leaf
		new.child = append(new.child, n.child[split_index+1:]...) //copy items
		new_array2 := make([]*node, 5)                            //new array
		copy(new_array2, n.child[:split_index])                   //copy child
		n.child = new_array2                                      //assing array
	}
	return item, &new

}

func (n *node) Insert(value int) {

}
