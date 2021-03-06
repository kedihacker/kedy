package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type node struct {
	items   []int
	child   []*node
	numkeys int
}

const degree = 5

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
// which item
func (n *node) split() (int, *node, bool) {
	new := node{}                   //new
	split_index := len(n.items) / 2 // index to cut
	item := n.items[split_index]    // does item retrivition
	var leaf bool
	if len(n.child) > 0 { //if isnt leaf

		leaf = false

		new.items = append(new.items, n.items[split_index+1:]...) //copy items
		new_array := make([]int, 0)                               //new array
		n.items = new_array                                       //ASSİNG array
		copy(new_array, n.items[:split_index])                    //turnucatw array

		new.child = append(new.child, n.child[split_index+1:]...) //copy items
		new_array2 := make([]*node, 0)                            //new array
		n.child = new_array2                                      //assing array
		copy(new_array2, n.child[:split_index])                   //copy child

	}
	if len(n.child) == 0 {
		leaf = true

		new.items = append(new.items, n.items[split_index:]...) //copy items
		new_array := make([]int, len(n.items)-2)                //new array 2
		copy(new_array, n.items[:split_index])                  //turnucatw array
		n.items = new_array
	}
	return item, &new, leaf

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

//hangi node insert edilecek, bulunup bulunmadığı
//ilk value nil ise hata
func (n *node) find(value int) (*node, bool) {
	low := 0
	find := false
	if len(n.child) != 0 {
		for x := 0; x < len(n.items); x++ {

			if value < n.items[x] { // sıralı listede kendinden büyük en soldaki sayıyı bul

				low = x
				find = true
				break
			}
		}
	} else {
		for x := 0; x < len(n.items); x++ { // son node ise itemi bul ve gönder
			if value == n.items[x] {
				return n, true

			}
		}
		return n, false
	}
	if !find {
		if n.items[n.numkeys] < value {
			return n.child[n.numkeys+1].find(value)
		}

	} else {
		return n.child[low].find(value)
	}
	return nil, false
}

// indexi bulur insetedilmesi gerken yere
// eşitse doğru gönderiri
// eşitse o zaman hangi indexte bulundupunu gönderirir
func (n *node) insertion_index(value int) (int, bool) {
	i := sort.Search(len(n.items), func(i int) bool { //valuedan büyük en soldaki ndexi veriri
		if value < n.items[i] {
			return true

		} else {
			return false
		}
	})
	if i > 0 && !(n.items[i-1] < value) { //listede olup olmadığını bulır
		return i - 1, true
	}
	return i, false
}

// bulundumu ve çağıran çağırdığı ı split etmelimi
func (n *node) insert(value int) (bool, bool) {
	index, found := n.insertion_index(value)
	if found {
		return false, false
	}
	if len(n.child) == 0 { // buraya insert sonra split et geekitse
		target_array := make([]int, 1)

		target_array = append(target_array, n.items[:index]...)
		target_array = append(target_array, value)
		target_array = append(target_array, n.items[index:]...)

		n.items = target_array
		if !(len(n.items) < degree) {
			return false, true
		} else {
			return false, false
		}
	} else { // çocuğu varsa
		exists, full := n.child[index].insert(value)
		if exists {
			return true, false
		}
		if full {
			item, newnode, _ := n.child[index].split()
			n.child = append(n.child[:index+1], n.child[index:]...)
			n.child[index+1] = newnode
			n.items = append(n.items[:index+1], n.items[index:]...)
			n.items[index+1] = item
			return false, !(len(n.items) < degree)
		} else {
			return false, false
		}
	}

}

// func (n *node) insert(value int) {

// }

func (b *BTree) Insert(value int) {
	if b.root == nil {
		fmt.Print("tree is nil added")
		new_root := &node{}
		new_root.insert(value)
		b.root = new_root

	} else {
		_, split := b.root.insert(value)
		if split {
			item, newnode, _ := b.root.split()
			newroot := &node{}
			newroot.items = append(newroot.items, item)
			newroot.child = append(newroot.child, b.root)
			newroot.child = append(newroot.child, newnode)
			b.root = newroot
		}
	}
}

func main() {
	tree := &BTree{}
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	tree.Insert(int(r.Int63n(1213213)))
	fmt.Println(tree.root)

	tree.Insert(int(r.Int63n(1213213)))
	fmt.Println(tree.root)
	tree.Insert(int(r.Int63n(1213213)))
	tree.Insert(int(r.Int63n(1213213)))
	fmt.Println(tree.root)
	tree.Insert(int(r.Int63n(1213213)))
	fmt.Println(tree.root)
	tree.Insert(int(r.Int63n(1213213)))
	tree.Insert(int(r.Int63n(1213213)))
	tree.Insert(int(r.Int63n(1213213)))

	fmt.Printf("%#v\n", tree.root)
}
