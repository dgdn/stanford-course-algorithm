package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var vertices []int
	var edges []int

	vertice_edges := map[int][]int{}
	edge_vertices := map[int][2]int{}

	vertices = []int{0, 1, 2, 3}
	edges = []int{0, 1, 2, 3, 4}

	vertice_edges = map[int][]int{
		0: {0, 1, 2},
		1: {1, 3},
		2: {0, 3, 4},
		3: {2, 4},
	}
	edge_vertices = map[int][2]int{
		0: {0, 2},
		1: {0, 1},
		2: {0, 3},
		3: {1, 2},
		4: {2, 3},
	}

	minCut(vertices, edges, vertice_edges, edge_vertices)
}

func minCut(vertices, edges []int, vertice_edges map[int][]int, edge_vertices map[int][2]int) int {

	for len(vertices) > 2 {

		//1. random select an edges
		rnd := rand.Intn(len(edges))
		rand.Seed(time.Now().UTC().UnixNano())
		edge := edges[rnd]
		fmt.Println("fuse edge", edge)

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

		for i, e := range edges {

			vs := edge_vertices[e]
			if vs[0] == vs[1] {

				//delete
				edges = append(edges[:i], edges[i+1:]...)
				delete(edge_vertices, e)

				//remove related vertice_edges
				vertice_edges[vs[0]] = removeElement(vertice_edges[vs[0]], e)

				break

			}

		}

		fmt.Println(vertices, vertice_edges)

	}

	return len(vertice_edges[0])
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
