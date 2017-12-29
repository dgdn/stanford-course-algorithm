package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	calMedian()
}

func test() {
	scanner := bufio.NewScanner(os.Stdin)
	heap := new(Heap)

	for i := 0; i < 4 && scanner.Scan(); i++ {
		d, err := strconv.Atoi(scanner.Text())
		if err == nil {
			heap.Insert(d)
			fmt.Println(heap.Min())
		}
	}
	for len(heap.List) > 0 {
		fmt.Println(heap.ExtractMin())
	}

}

func printHeap(h []*Node) {
	fmt.Print("[")
	for _, d := range h {
		fmt.Print(d.Key, " ")
	}
	fmt.Print("]")

}

func calMedian() {
	f, err := os.Open("Median.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	results := median(f)
	var total int
	for _, r := range results {
		total += r
	}
	fmt.Println(total % 10000)
}

func median(reader io.Reader) []int {

	scanner := bufio.NewScanner(reader)
	maxHeap := new(MaxHeap)
	minHeap := new(Heap)

	var medians []int
	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("not a number", scanner.Text())
			continue
		}

		if len(maxHeap.List) == 0 {
			maxHeap.Insert(d)
			medians = append(medians, maxHeap.Max())
			continue
		}

		if len(minHeap.List) == 0 {
			if d < maxHeap.Max() {
				n := maxHeap.ExtractMax()
				minHeap.Insert(n)
				maxHeap.Insert(d)
			} else {
				minHeap.Insert(d)
			}
			medians = append(medians, maxHeap.Max())
			continue
		}

		if minHeap.Min() < d {
			if len(minHeap.List) < len(maxHeap.List) {
				minHeap.Insert(d)
			} else {
				n := minHeap.ExtractMin()
				maxHeap.Insert(n)
				minHeap.Insert(d)
			}
		} else if maxHeap.Max() >= d {
			if len(maxHeap.List) <= len(minHeap.List) {
				maxHeap.Insert(d)
			} else {
				n := maxHeap.ExtractMax()
				minHeap.Insert(n)
				maxHeap.Insert(d)
			}
		} else {
			if len(maxHeap.List) <= len(minHeap.List) {
				maxHeap.Insert(d)
			} else {
				minHeap.Insert(d)
			}
		}

		medians = append(medians, maxHeap.Max())

	}
	return medians

}

type Node struct {
	Key    int
	Parent *Node
	LChild *Node
	RChild *Node
}

type Heap struct {
	List []*Node
}

func (h *Heap) Min() int {
	return h.List[0].Key
}

func (h *Heap) ExtractMin() int {
	min := h.List[0].Key

	if len(h.List) == 1 {
		h.List = nil
		return min
	}

	//assign the root node key to be the key of rightmost buttom node
	rightmostNode := h.List[len(h.List)-1]
	h.List[0].Key = rightmostNode.Key
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
			tmp := child.Key
			child.Key = n.Key
			n.Key = tmp

			n = child
			child = nil
		} else {
			break
		}

	}

	return min

}

func (h *Heap) Insert(d int) {

	n := new(Node)
	n.Key = d
	h.List = append(h.List, n)

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

	//bubble up
	for n.Parent != nil && n.Key < n.Parent.Key {
		//swap
		tmp := n.Key
		n.Key = n.Parent.Key
		n.Parent.Key = tmp

		n = n.Parent
	}

}

type MaxHeap struct {
	List []*Node
}

func (h *MaxHeap) ExtractMax() int {
	min := h.List[0].Key

	if len(h.List) == 1 {
		h.List = nil
		return min
	}

	//assign the root node key to be the key of rightmost buttom node
	rightmostNode := h.List[len(h.List)-1]
	h.List[0].Key = rightmostNode.Key
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
			if n.Key < n.LChild.Key {
				child = n.LChild
			}
		}
		if n.LChild == nil && n.RChild != nil {
			if n.Key < n.RChild.Key {
				child = n.RChild
			}
		}
		if n.LChild != nil && n.RChild != nil {

			if n.Key < n.LChild.Key && n.Key >= n.RChild.Key {
				child = n.LChild
			}
			if n.Key < n.RChild.Key && n.Key >= n.LChild.Key {
				child = n.RChild
			}
			if n.Key < n.RChild.Key && n.Key < n.LChild.Key {
				if n.RChild.Key < n.LChild.Key {
					child = n.LChild
				} else {
					child = n.RChild
				}
			}
		}

		if child != nil {
			tmp := child.Key
			child.Key = n.Key
			n.Key = tmp

			n = child
			child = nil
		} else {
			break
		}

	}

	return min

}

func (h *MaxHeap) Max() int {
	return h.List[0].Key
}

func (h *MaxHeap) Insert(d int) {
	n := new(Node)
	n.Key = d
	h.List = append(h.List, n)

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

	//bubble up
	for n.Parent != nil && n.Key > n.Parent.Key {
		//swap
		tmp := n.Key
		n.Key = n.Parent.Key
		n.Parent.Key = tmp

		n = n.Parent
	}
}
