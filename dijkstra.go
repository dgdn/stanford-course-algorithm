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

	f, err := os.Open("dijkstraData.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g := newGraphReader(f)

	//printGraph(g)

	//assignment part
	//select vertex 1 as source
	sourceVertex := "1"
	A := dijkstra(g, sourceVertex)

	var result []string
	for _, d := range strings.Split("7,37,59,82,99,115,133,165,188,197", ",") {
		result = append(result, fmt.Sprint(A[d]))
	}
	fmt.Print(strings.Join(result, ","))
}

func printGraph(g *Graph) {
	for k, v := range g.Vertices {
		fmt.Printf("node: key %v, label %v edges %v \n", k, v.Label, len(v.Edges))
	}
	for _, e := range g.Edges {
		fmt.Printf("edge:  (%v, %v) len %v \n", e.V1.Label, e.V2.Label, e.Len)
	}
}

func testData() io.Reader {
	return strings.NewReader(
		"s v,1 w,4\n" +
			"v s,1 w,2 t,6\n" +
			"w s,4 v,2 t,3\n" +
			"t v,6 w,3\n",
	)
}

type Graph struct {
	Vertices map[string]*Vertex
	Edges    []Edge
}

type Vertex struct {
	Label string
	Edges []Edge
}

type Edge struct {
	Len int
	V1  *Vertex
	V2  *Vertex
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
	newEdge.Len = length
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
	for scanner.Scan() {

		rowItems := strings.Fields(scanner.Text())
		v1 := g.saveAndFetchVertex(rowItems[0])

		for _, item := range rowItems[1:] {

			subItem := strings.Split(item, ",")
			edgeLength, _ := strconv.Atoi(subItem[1])

			v2 := g.saveAndFetchVertex(subItem[0])

			edge := g.saveAndFetchEdge(v1, v2, edgeLength)

			v1.Edges = append(v1.Edges, edge)

		}
	}

	return g
}

func dijkstra(g *Graph, sourceVertex string) map[string]int {

	//vertices processed so far
	X := map[string]bool{
		sourceVertex: true,
	}

	//computed shrotest path distance
	A := map[string]int{
		sourceVertex: 0,
	}

	for len(X) != len(g.Vertices) {

		//among edges (v,w) with v in X,  w in V-X
		//pick the one minimize  A[v] + lvw, call it (v*,m*)
		min := 100000
		var mv, mw string
		for _, edge := range g.Edges {
			v, w, exists := fetchEdge(edge, X)
			if exists {
				score := A[v] + edge.Len
				if score < min {
					min = score
					mv, mw = v, w
				}
			}
		}

		if mv == "" || mw == "" {
			panic("assert fail")
		}

		X[mw] = true
		A[mw] = min

	}

	return A

}

func fetchEdge(edge Edge, X map[string]bool) (v, w string, exists bool) {

	if X[edge.V1.Label] && !X[edge.V2.Label] {
		return edge.V1.Label, edge.V2.Label, true
	}

	if !X[edge.V1.Label] && X[edge.V2.Label] {
		return edge.V2.Label, edge.V1.Label, true

	}

	return
}
