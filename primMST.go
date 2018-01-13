package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("mst.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g := newGraphReader(f)
	_, T := findMST(g)

	var totalCost int
	for _, e := range T {
		totalCost += e.Cost
	}
	fmt.Println(totalCost)
}

func (g *Graph) saveAndFetchVertex(label string) *Vertex {
	if v, ok := g.Vertices[label]; ok {
		return v
	} else {
		v := new(Vertex)
		v.Label = label
		g.Vertices[label] = v
		return v
	}
}

func (g *Graph) saveAndFetchEdge(v1, v2 *Vertex, length int) Edge {

	if e, exists := g.getEdge(v1, v2); exists {
		return e
	}

	var newEdge Edge
	newEdge.Cost = length
	newEdge.V1 = v1
	newEdge.V2 = v2

	g.Edges = append(g.Edges, newEdge)

	return newEdge

}

func (g *Graph) getEdge(v1, v2 *Vertex) (e Edge, exists bool) {

	for _, edge := range g.Edges {
		if edge.V1.Label == v1.Label && edge.V2.Label == v2.Label {
			return edge, true
		}
		if edge.V1.Label == v2.Label && edge.V2.Label == v1.Label {
			return edge, true
		}
	}
	return
}

func newGraphReader(reader io.Reader) *Graph {

	scanner := bufio.NewScanner(reader)
	g := new(Graph)
	g.Vertices = make(map[string]*Vertex)
	for i := 0; scanner.Scan(); i++ {
		//skip the first line
		if i == 0 {
			continue
		}

		rowItems := strings.Fields(scanner.Text())
		v1 := g.saveAndFetchVertex(rowItems[0])
		v2 := g.saveAndFetchVertex(rowItems[1])

		edgeLength, err := strconv.Atoi(rowItems[2])
		if err != nil {
			panic(err)
		}
		g.saveAndFetchEdge(v1, v2, edgeLength)

		v1.Edges = append(v1.Edges, Edge{Cost: edgeLength, V1: v1, V2: v2})
		v2.Edges = append(v2.Edges, Edge{Cost: edgeLength, V1: v2, V2: v1})

	}

	return g
}

func printGraph(g *Graph) {
	for k, v := range g.Vertices {
		fmt.Printf("node: key %v, label %v edges %v \n", k, v.Label, len(v.Edges))
	}
	for _, e := range g.Edges {
		fmt.Printf("edge:  (%v, %v) len %v \n", e.V1.Label, e.V2.Label, e.Cost)
	}
}

func findMST(g *Graph) (map[string]*Vertex, []Edge) {

	//vertice processed so far
	X := make(map[string]*Vertex)
	var T []Edge //chosen spanning tree edges

	heap := NewHeap()
	INFINIT_COST := 10000000

	//init the heap
	var initVertex *Vertex
	initVertexAjacentEdge := make(map[string]Edge)
	for _, vertex := range g.Vertices {
		if initVertex == nil {
			//arbitarily choose one vertex say the first
			initVertex = vertex
			X[initVertex.Label] = initVertex
			for _, edge := range initVertex.Edges {
				initVertexAjacentEdge[edge.V2.Label] = edge
			}
		} else {
			if edge, ok := initVertexAjacentEdge[vertex.Label]; ok {
				vertex.CurrentLeastCostEdge = edge
				heap.Insert(edge.Cost, vertex)
			} else {
				heap.Insert(INFINIT_COST, vertex)
			}
		}

	}

	//greedy iteration
	for len(X) != len(g.Vertices) {
		minNode := heap.ExtractMin()
		X[minNode.Vertex.Label] = minNode.Vertex
		T = append(T, minNode.Vertex.CurrentLeastCostEdge)

		for _, edge := range minNode.Vertex.Edges {
			// endponit in V-X
			if _, ok := X[edge.V2.Label]; !ok {

				//featch node
				wnode := heap.Map[edge.V2.Label]

				//compute new key
				if wnode.Key > edge.Cost {
					wnode.Vertex.CurrentLeastCostEdge = edge
					wnode.Key = edge.Cost
					heap.ShouldBubbleUp(wnode)
				}

			}
		}
	}

	return X, T
}

func debugInfo(heap *Heap, X map[string]*Vertex) {
	var xx []string
	for k := range X {
		xx = append(xx, k)
	}
	type d struct {
		key   int
		label string
	}
	var keys []d
	for _, n := range heap.List {
		keys = append(keys, d{n.Key, n.Vertex.Label})
	}
	fmt.Println(xx, keys)
}

type Graph struct {
	Vertices map[string]*Vertex
	Edges    []Edge
}

type Vertex struct {
	Label                string
	Edges                []Edge
	CurrentLeastCostEdge Edge
}

type Edge struct {
	Cost int
	V1   *Vertex
	V2   *Vertex
}

type Node struct {
	Key    int
	Vertex *Vertex
	Parent *Node
	LChild *Node
	RChild *Node
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
		h.Map[n.Parent.Vertex.Label] = n
		h.Map[n.Vertex.Label] = n.Parent

		//swap
		tmp := n.Key
		tmpV := n.Vertex
		n.Key = n.Parent.Key
		n.Vertex = n.Parent.Vertex
		n.Parent.Key = tmp
		n.Parent.Vertex = tmpV

		n = n.Parent
	}
}

func (h *Heap) Min() *Node {
	return h.List[0]
}

func (h *Heap) ExtractMin() *Node {
	min := &Node{
		Key:    h.List[0].Key,
		Vertex: h.List[0].Vertex,
	}

	if len(h.List) == 1 {
		h.List = nil
		return min
	}

	//assign the root node key to be the key of rightmost buttom node
	rightmostNode := h.List[len(h.List)-1]
	h.List[0].Key = rightmostNode.Key
	h.List[0].Vertex = rightmostNode.Vertex
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

			h.Map[n.Vertex.Label] = child
			h.Map[child.Vertex.Label] = n

			tmp := child.Key
			tmpV := child.Vertex
			child.Key = n.Key
			child.Vertex = n.Vertex
			n.Key = tmp
			n.Vertex = tmpV

			n = child
			child = nil

		} else {
			break
		}

	}

	return min

}

func (h *Heap) Insert(d int, v *Vertex) {

	n := new(Node)
	n.Key = d
	n.Vertex = v

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
