package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("huffman.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	heap := NewHeap()
	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		data := scanner.Text()
		d, err := strconv.Atoi(data)
		if err != nil {
			continue
		}
		heap.Insert(d, &NodeData{Label: fmt.Sprint(i)})

	}
	tree, _ := haffman(heap)

	var minLevel = 100000
	var maxLevel = 0
	iterateTreeLevel(tree, 0, &minLevel, &maxLevel)

	//minus 1 represent only count intermedia edge
	//TODO to figure out why maxLevel should minus 2
	fmt.Println("min level:", minLevel-1, "\nmax level:", maxLevel-2)
}

func iterateTreeLevel(tree *Tree, level int, minLevel, maxLevel *int) {
	level++
	if tree.Left != nil {
		iterateTreeLevel(tree.Left, level, minLevel, maxLevel)
	} else {
		if level < *minLevel {
			*minLevel = level
		}
		if level > *maxLevel {
			*maxLevel = level
		}
	}
	if tree.Right != nil {
		iterateTreeLevel(tree.Right, level, minLevel, maxLevel)
	} else {
		if level < *minLevel {
			*minLevel = level
		}
		if level > *maxLevel {
			*maxLevel = level
		}
	}
}

func haffman(heap *Heap) (*Tree, map[string]*Tree) {
	if len(heap.List) == 2 {
		treeMap := make(map[string]*Tree)
		tree := new(Tree)
		tree.Label = heap.List[0].NodeData.Label + heap.List[1].NodeData.Label
		tree.Left = new(Tree)
		tree.Left.Label = heap.List[0].NodeData.Label
		treeMap[tree.Left.Label] = tree.Left
		tree.Right = new(Tree)
		tree.Right.Label = heap.List[1].NodeData.Label
		treeMap[tree.Right.Label] = tree.Right
		return tree, treeMap
	}

	hnb := heap.ExtractMin()
	hns := heap.ExtractMin()
	label := fmt.Sprintf("%d%d", hnb.Key, hns.Key)
	heap.Insert(hnb.Key+hns.Key, &NodeData{Label: label})

	ntree, treeMap := haffman(heap)
	ptree := treeMap[label]
	ptree.Left = new(Tree)
	ptree.Left.Label = hns.NodeData.Label
	treeMap[ptree.Left.Label] = ptree.Left

	ptree.Right = new(Tree)
	ptree.Right.Label = hnb.NodeData.Label
	treeMap[ptree.Right.Label] = ptree.Right

	return ntree, treeMap
}

type Tree struct {
	Label string
	Left  *Tree
	Right *Tree
}

type NodeData struct {
	Label string
}

type Node struct {
	Key      int
	NodeData *NodeData
	Parent   *Node
	LChild   *Node
	RChild   *Node
}

type Heap struct {
	List []*Node
	Map  map[string]*Node
}

func NewHeap() *Heap {
	h := new(Heap)
	h.Map = make(map[string]*Node)
	return h
}

func (h *Heap) ShouldBubbleUp(n *Node) {
	//bubble up
	for n.Parent != nil && n.Key < n.Parent.Key {
		h.Map[n.Parent.NodeData.Label] = n
		h.Map[n.NodeData.Label] = n.Parent

		//swap
		tmp := n.Key
		tmpV := n.NodeData
		n.Key = n.Parent.Key
		n.NodeData = n.Parent.NodeData
		n.Parent.Key = tmp
		n.Parent.NodeData = tmpV

		n = n.Parent
	}
}

func (h *Heap) Min() *Node {
	return h.List[0]
}

func (h *Heap) ExtractMin() *Node {
	min := &Node{
		Key:      h.List[0].Key,
		NodeData: h.List[0].NodeData,
	}

	if len(h.List) == 1 {
		h.List = nil
		return min
	}

	//assign the root node key to be the key of rightmost buttom node
	rightmostNode := h.List[len(h.List)-1]
	h.List[0].Key = rightmostNode.Key
	h.List[0].NodeData = rightmostNode.NodeData
	//delete the rightmost buttom node
	h.List = h.List[:len(h.List)-1]
	if rightmostNode.Parent.LChild == rightmostNode {
		rightmostNode.Parent.LChild = nil
	} else {
		rightmostNode.Parent.RChild = nil
	}

	//bubble down
	n := h.List[0]
	var child *Node
	for {

		//select the right node to bubble down
		if n.LChild != nil && n.RChild == nil {
			if n.Key > n.LChild.Key {
				child = n.LChild
			}
		}
		if n.LChild == nil && n.RChild != nil {
			if n.Key > n.RChild.Key {
				child = n.RChild
			}
		}
		if n.LChild != nil && n.RChild != nil {

			if n.Key > n.LChild.Key && n.Key <= n.RChild.Key {
				child = n.LChild
			}
			if n.Key > n.RChild.Key && n.Key <= n.LChild.Key {
				child = n.RChild
			}
			if n.Key > n.RChild.Key && n.Key > n.LChild.Key {
				if n.RChild.Key > n.LChild.Key {
					child = n.LChild
				} else {
					child = n.RChild
				}
			}
		}

		if child != nil {

			h.Map[n.NodeData.Label] = child
			h.Map[child.NodeData.Label] = n

			tmp := child.Key
			tmpV := child.NodeData
			child.Key = n.Key
			child.NodeData = n.NodeData
			n.Key = tmp
			n.NodeData = tmpV

			n = child
			child = nil

		} else {
			break
		}

	}

	return min

}

func (h *Heap) Insert(d int, v *NodeData) {

	n := new(Node)
	n.Key = d
	n.NodeData = v

	h.List = append(h.List, n)
	h.Map[v.Label] = n

	if len(h.List) == 1 {
		return
	}

	//choose parent node to insert as child
	idx := len(h.List)/2 - 1
	parent := h.List[idx]
	if parent.LChild == nil {
		parent.LChild = n
		n.Parent = parent
	} else if parent.RChild == nil {
		parent.RChild = n
		n.Parent = parent
	}

	h.ShouldBubbleUp(n)
}
