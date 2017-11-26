package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var edgeStore = make(map[[2]int]int)
var edgeStoreCount = 0

func saveAndFetchEdge(e1, e2 int) (int, bool) {
	var max, min int
	if e1 > e2 {
		max = e1
		min = e2
	} else {
		max = e2
		min = e1
	}

	key := [2]int{min, max}
	if v, exist := edgeStore[key]; exist {
		return v, false
	}
	edgeStore[key] = edgeStoreCount
	edgeStoreCount++
	return edgeStore[key], true
}

func main() {

	var vertices []int
	var edges []int

	vertice_edges := map[int][]int{}
	edge_vertices := map[int][2]int{}

	f, err := os.Open("kargerMinCut.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var datas [][]int
	for scanner.Scan() {
		var nums []int
		for _, n := range strings.Fields(scanner.Text()) {

			d, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nums = append(nums, d)
		}
		datas = append(datas, nums)

	}

	for _, data := range datas {

		//vertice
		vertice := data[0]
		vertices = append(vertices, vertice)

		for i := 1; i < len(data); i++ {

			//add edge
			edge, isNew := saveAndFetchEdge(vertice, data[i])
			if isNew {
				edges = append(edges, edge)
			}

			//add vertice_edge
			vertice_edges[vertice] = append(vertice_edges[vertice], edge)

			//add edge_vertices
			edge_vertices[edge] = [2]int{vertice, data[i]}
		}

	}

	minCount := 1000
	for i := 0; i < 1000; i++ {
		c_vertices := make([]int, len(vertices))
		copy(c_vertices, vertices)
		c_edges := make([]int, len(edges))
		copy(c_edges, edges)

		c_vertice_edges := make(map[int][]int)
		for k, v := range vertice_edges {
			_v := []int{}
			for _, d := range v {
				_v = append(_v, d)
			}
			c_vertice_edges[k] = _v
		}
		c_edge_vertices := make(map[int][2]int)
		for k, v := range edge_vertices {
			c_edge_vertices[k] = [2]int{v[0], v[1]}
		}

		//provide clean data on every call
		length := minCut(c_vertices, c_edges, c_vertice_edges, c_edge_vertices)
		if length < minCount {
			minCount = length
		}
		if i != 0 && i%100 == 0 {
			fmt.Println(minCount)
		}
	}

	fmt.Println("min cut is", minCount)

}

func minCut(vertices, edges []int, vertice_edges map[int][]int, edge_vertices map[int][2]int) int {

	for len(vertices) > 2 {

		//1. random select an edges
		rnd := rand.Intn(len(edges))
		rand.Seed(time.Now().UTC().UnixNano())
		edge := edges[rnd]

		//2.fuse vertice
		v1, v2 := edge_vertices[edge][0], edge_vertices[edge][1]
		//2.1 merge v2 edges to v1
		for _, e := range vertice_edges[v2] {
			v := edge_vertices[e]

			//ignore the selected edge
			if e == edge {
				continue
			}

			if v[0] == v2 {
				edge_vertices[e] = [2]int{v1, v[1]}
			} else {
				edge_vertices[e] = [2]int{v[0], v1}
			}

			//TODO  dont't append duplicated edge
			vertice_edges[v1] = append(vertice_edges[v1], e)
		}
		//2.2 delete v1 edges
		vertice_edges[v1] = removeElement(vertice_edges[v1], edge)
		//2.3 delete seleted edge
		edges = removeElement(edges, edge)
		delete(edge_vertices, edge)
		//2.4 delete v2
		vertices = removeElement(vertices, v2)
		delete(vertice_edges, v2)

		//3. delete self loop

		for i := 0; i < len(edges); i++ {

			e := edges[i]

			vs := edge_vertices[e]
			if vs[0] == vs[1] {

				//delete
				edges = append(edges[:i], edges[i+1:]...)
				i--
				delete(edge_vertices, e)

				//remove related vertice_edges
				vertice_edges[vs[0]] = removeElement(vertice_edges[vs[0]], e)

			}

		}

	}

	return len(vertice_edges[vertices[0]])
}

func removeElement(a []int, n int) []int {

	for i := 0; i < len(a); i++ {

		if a[i] == n {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}

	return a
}
