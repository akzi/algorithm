package pairingheap

/*
https://en.wikipedia.org/wiki/Pairing_heap
*/

type (
	Iterator func(Item) bool

	Item interface {
		Less(Item) bool
	}
	Node struct {
		item     Item
		parent   *Node
		children []*Node
	}
	Heap struct {
		root *Node
	}
)

func (node *Node) detach() []*Node {
	for i, child := range node.parent.children {
		if child == node {
			copy(node.parent.children[i:], node.parent.children[i+1:])
			node.parent.children[len(node.parent.children)-1] = nil
			node.parent.children = node.parent.children[:len(node.parent.children)-1]
		}
	}
	return node.children
}

func (node *Node) nodeRange(iterator Iterator) bool {
	if node == nil {
		return true
	}
	if iterator(node.item) {
		for _, child := range node.children {
			if child.nodeRange(iterator) == false {
				return false
			}
		}
	}
	return true
}

func (node *Node) find(item Item) *Node {
	if !node.item.Less(item) &&
		!item.Less(node.item) {
		return node
	}
	for _, node := range node.children {
		if node := node.find(item); node == nil {
			continue
		}
		return node
	}
	return nil
}

func (heap *Heap) merge(first, second *Node) *Node {
	if first == nil {
		return second
	} else if second == nil {
		return first
	} else if first.item.Less(second.item) {
		first.children = append([]*Node{second}, first.children...)
		second.parent = first
		return first
	} else {
		second.children = append([]*Node{first}, second.children...)
		first.parent = second
		return second
	}
}

func (heap *Heap) mergePairs(nodes []*Node) *Node {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return nodes[0]
	default:
		return heap.merge(heap.merge(nodes[0], nodes[1]), heap.mergePairs(nodes[2:]))
	}
}

func New() *Heap {
	return &Heap{}
}

func (heap *Heap) FindMin() Item {
	if heap.root == nil {
		return nil
	}
	return heap.root.item
}

func (heap *Heap) Insert(item Item) {
	heap.root = heap.merge(heap.root, &Node{item: item})
}

func (heap *Heap) DeleteMin() {
	if heap.root == nil {
		return
	}
	heap.root = heap.mergePairs(heap.root.children)
}

func (heap *Heap) Delete(item Item) {
	if heap.root == nil {
		return
	}
	node := heap.root.find(item)
	if node == nil {
		return
	}
	//delete min
	if node == heap.root {
		heap.root = heap.mergePairs(heap.root.children)
		return
	}
	for _, node := range node.detach() {
		heap.root = heap.merge(heap.root, node)
	}
}

func (heap *Heap) Range(iterator Iterator) {
	heap.root.nodeRange(iterator)
}
