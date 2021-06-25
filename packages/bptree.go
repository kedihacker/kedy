package bptree

type node struct {
	items []int
	child []*node
}

type BTree struct {
	degree int
	// length int
	root *node
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

//right items copyied to new node
func (n *node) split() (int, *node) {
	new := node{} //new
	split_index := len(n.items) / 2
	item := n.items[split_index]                            //set index
	new.items = append(new.items, n.items[split_index:]...) //copy items
	new_array := make([]int, 5)                             //new array
	copy(new_array, n.items[:split_index])                  //turnucatw array

	n.items = new_array //ASSİNG array

	if len(n.child) > 0 { //if isnt leaf
		new.child = append(new.child, n.child[split_index:]...) //copy items
		new_array2 := make([]*node, 5)                          //new array
		copy(new_array2, n.child[:split_index])                 //copy child
		n.child = new_array2                                    //assing array
	}
	return item, &new

}

// func (t *BTree) ReplaceOrInsert(item Item) Item {
// 	if item == nil {
// 		panic("nil item being added to BTree")
// 	}
// 	if t.root == nil {
// 		t.root = t.cow.newNode()
// 		t.root.items = append(t.root.items, item)
// 		t.length++
// 		return nil
// 	} else {
// 		t.root = t.root.mutableFor(t.cow)
// 		if len(t.root.items) >= t.maxItems() {
// 			item2, second := t.root.split(t.maxItems() / 2)
// 			oldroot := t.root
// 			t.root = t.cow.newNode()
// 			t.root.items = append(t.root.items, item2)
// 			t.root.children = append(t.root.children, oldroot, second)
// 		}
// 	}
// 	out := t.root.insert(item, t.maxItems())
// 	if out == nil {
// 		t.length++
// 	}
// 	return out
// }

func (n *node) find(value int) *node {

}

func (b *BTree) Insert(value int) {
	if b.root == nil {
		new_root := node{}
		new_root.insert()
	}
}
