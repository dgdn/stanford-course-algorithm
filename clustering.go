package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g := newGraphReader(f)
	clustering(g, 4)
}

type Graph struct {
	Vertices map[string]*Vertex
	Edges    []Edge
}

type Vertex struct {
	Leader string
	Label  string
	Edges  []Edge
}

type Edge struct {
	Cost int
	V1   *Vertex
	V2   *Vertex
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
		if len(rowItems) == 0 {
			continue
		}

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

type UnionFind struct {
	Groups map[string][]*Vertex
}

func (u *UnionFind) Find(label string) *Vertex {
	return nil
}

func (u *UnionFind) GroupLength() int {
	return len(u.Groups)
}

func (u *UnionFind) AllGroupLength() []int {
	var ls []int
	for _, g := range u.Groups {
		ls = append(ls, len(g))
	}
	return ls
}

func (u *UnionFind) Add(leader string, v *Vertex) {
	group := u.Groups[leader]
	group = append(group, v)
	u.Groups[leader] = group
}

func (u *UnionFind) Fuse(leader1, leader2 string) {

	var bigGroupLeader, smallGroupLeader string
	//fuse the small group to the big groups
	if len(u.Groups[leader1]) > len(u.Groups[leader2]) {
		bigGroupLeader = leader1
		smallGroupLeader = leader2
	} else {
		bigGroupLeader = leader2
		smallGroupLeader = leader1
	}

	bigGroup := u.Groups[bigGroupLeader]
	for _, v := range u.Groups[smallGroupLeader] {
		v.Leader = bigGroupLeader
		bigGroup = append(bigGroup, v)
	}
	u.Groups[bigGroupLeader] = bigGroup
	delete(u.Groups, smallGroupLeader)

}

func NewUnionFind() *UnionFind {
	u := new(UnionFind)
	u.Groups = make(map[string][]*Vertex)
	return u
}

func clustering(g *Graph, k int) []int {

	//sort the edges in order of increasing cost
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Cost < g.Edges[j].Cost
	})

	unionFind := NewUnionFind()
	//initialize all vertices to union find data structure
	//the vertex leader should be itself
	for _, v := range g.Vertices {
		v.Leader = v.Label
		unionFind.Add(v.Leader, v)
	}

	edgeIdx := 0
	for unionFind.GroupLength() != k {
		edge := g.Edges[edgeIdx]
		if edge.V1.Leader != edge.V2.Leader {
			unionFind.Fuse(edge.V1.Leader, edge.V2.Leader)
		}
		edgeIdx++
	}

	mins := map[string]int{}
	for leader, group := range unionFind.Groups {

		groupVertices := map[string]*Vertex{}
		for _, v := range group {
			groupVertices[v.Label] = v
		}

		var min = 10000000
		for _, v := range group {
			for _, e := range v.Edges {
				if keyExsist(e.V1.Label, groupVertices) && !keyExsist(e.V2.Label, groupVertices) ||
					!keyExsist(e.V1.Label, groupVertices) && keyExsist(e.V2.Label, groupVertices) {
					if e.Cost < min {
						min = e.Cost
					}
				}
			}
		}

		mins[leader] = min

	}

	for _, v := range unionFind.Groups {
		var labels []string
		for _, n := range v {
			labels = append(labels, n.Label)
		}
		fmt.Println(labels)
	}

	fmt.Println(mins)

	return unionFind.AllGroupLength()
}

func keyExsist(key string, vertices map[string]*Vertex) bool {
	if _, ok := vertices[key]; ok {
		return true
	}
	return false
}
