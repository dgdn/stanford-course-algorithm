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

	f, err := os.Open("SCC.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g := newGraphReader(f)

	scc(g)

	results := calculateNumComponentNodes(g)
	length := len(results)

	fmt.Println(results[length-5:])

}

func testData() io.Reader {
	return strings.NewReader(
		"5 4\n4 6\n6 5\n5 2\n2 1\n1 3\n3 2",
	)
}

func printGraph(g *Graph) {
	for _, v := range g.Vertices {
		var edges []string
		for idx, i := range v.Edges {
			edges = append(edges, fmt.Sprint(i.Tail.Label))
			edges = append(edges, fmt.Sprint(i.Head.Label))
			if idx != len(v.Edges)-1 {
				edges = append(edges, ",")
			}
		}
		fmt.Println("node: ", v.Label, edges, "f:", v.FinishTime, "leader:", v.Leader)
	}

	for _, edge := range g.Edges {
		var edges []string
		edges = append(edges, fmt.Sprint(edge.Tail.Label))
		edges = append(edges, fmt.Sprint(edge.Head.Label))
		fmt.Println("edge: ", edges)
	}

	for k, v := range g.VerticeMap {
		fmt.Println("node map: ", k, v.Label)
	}
}

func calculateNumComponentNodes(g *Graph) []int {

	leaderMap := make(map[int]int)
	for _, vertex := range g.Vertices {
		leaderMap[vertex.Leader] = leaderMap[vertex.Leader] + 1
	}
	var leaderNums []int
	var invalidLeader []int
	for k, v := range leaderMap {
		leaderNums = append(leaderNums, v)
		if v == 1 {
			invalidLeader = append(invalidLeader, k)
		}
	}
	sort.Ints(invalidLeader)

	sort.Ints(leaderNums)

	return leaderNums
}

func newGraphReader(reader io.Reader) *Graph {

	g := new(Graph)
	g.VerticeMap = make(map[int]*Vertex)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {

		edgeDatas := strings.Fields(scanner.Text())
		tail, err := strconv.Atoi(edgeDatas[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		head, err := strconv.Atoi(edgeDatas[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		ne := new(Edge)

		hv, ok := g.VerticeMap[head]
		if !ok {
			hv = new(Vertex)
			hv.Label = head
			g.Vertices = append(g.Vertices, hv)
			g.VerticeMap[head] = hv
		}
		ne.Head = hv

		tv, ok := g.VerticeMap[tail]
		if !ok {
			tv = new(Vertex)
			tv.Label = tail
			g.Vertices = append(g.Vertices, tv)
			g.VerticeMap[tail] = tv
		}
		ne.Tail = tv

		tv.Edges = append(tv.Edges, ne)
		g.Edges = append(g.Edges, ne)
	}
	return g
}

type Graph struct {

	//vertice index map for dfsLoop(may be label or finishingTime)
	VerticeMap map[int]*Vertex
	Vertices   []*Vertex
	Edges      []*Edge
}

type Edge struct {
	Tail *Vertex
	Head *Vertex
}

type Vertex struct {
	Label      int
	IsExplored bool
	Leader     int
	FinishTime int

	Edges []*Edge
}

// strongly connected component
func scc(g *Graph) {

	reversedGraph := createReverseArcGraph(g)
	//run DFS loop on reversed graph
	DFSLoop(reversedGraph)

	//process the origin graph node in the
	//decreasing order of rev graph's finishing time
	nm := make(map[int]*Vertex)
	for _, rv := range reversedGraph.Vertices {
		nm[rv.FinishTime] = g.VerticeMap[rv.Label]
	}
	g.VerticeMap = nm

	//run DFS loop on original graph
	DFSLoop(g)
}

//create a new graph with all arcs reversed
func createReverseArcGraph(g *Graph) *Graph {
	ng := new(Graph)

	ng.VerticeMap = make(map[int]*Vertex)

	for _, v := range g.Vertices {
		nv := new(Vertex)
		nv.Label = v.Label
		ng.Vertices = append(ng.Vertices, nv)

		ng.VerticeMap[nv.Label] = nv
	}

	//reversed arc
	for _, edge := range g.Edges {
		ne := new(Edge)
		ne.Tail = ng.VerticeMap[edge.Head.Label]
		ne.Head = ng.VerticeMap[edge.Tail.Label]
		ng.Edges = append(ng.Edges, ne)

		//add this edges to v
		ne.Tail.Edges = append(ne.Tail.Edges, ne)
	}

	return ng

}

//global finishing time
var finishingTime int

//global leader denoted by label
var leader int

func DFSLoop(g *Graph) {

	finishingTime = 0
	leader = 0

	for i := len(g.VerticeMap); i >= 1; i-- {
		node := g.VerticeMap[i]
		if !node.IsExplored {
			leader = node.Label
			DFS(g, node)

		}
	}
}

func DFS(g *Graph, node *Vertex) {
	node.Leader = leader
	node.IsExplored = true

	for _, edge := range node.Edges {
		anode := edge.Head
		if !anode.IsExplored {
			DFS(g, anode)
		}
	}
	finishingTime++
	node.FinishTime = finishingTime
}
