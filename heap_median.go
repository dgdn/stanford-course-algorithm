package main

func main() {

}

type Node struct {
	Key    int
	Parent *Node
	LChild *Node
	RChild *Node
}

type Heap struct {
	Root             *Node
	RightButtomChild *Node
}

func (h *Heap) ExtractMin() int {

}

func (h *Heap) Insert(d int) {

	n := new(Node)
	n.Key = d
	h.RightButtomChild = n
	if h.Root == nil {
		h.Root = n
		return
	}

	//

	//bubble up
	for n.Key < n.Parent.Key {
		//swap
		tmp := n.Key
		n.Key = n.Parent.Key
		n.Parent.Key = tmp

		n = n.Parent
		if n.Parent == nil {
			break
		}
	}

}
