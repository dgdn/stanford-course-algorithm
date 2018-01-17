package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("clustering_big.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	vm := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for line := 0; scanner.Scan(); line++ {
		if line == 0 {
			continue
		}
		text := scanner.Text()
		ts := strings.Fields(text)
		vm[strings.Join(ts, "")] = true
	}

	bigClustering(vm)
}

type Node struct {
	Leader string
	Label  string
}

type UnionFind struct {
	Groups map[string][]*Node
	Map    map[string]*Node
}

func (u *UnionFind) Find(label string) *Node {
	return u.Map[label]
}

func (u *UnionFind) Add(v *Node) {
	group := u.Groups[v.Leader]
	group = append(group, v)
	u.Groups[v.Leader] = group
	u.Map[v.Label] = v
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

func NewUnionFind() *UnionFind {
	u := new(UnionFind)
	u.Groups = make(map[string][]*Node)
	u.Map = make(map[string]*Node)
	return u
}

func bigClustering(vm map[string]bool) {

	unionFind := NewUnionFind()
	for v := range vm {
		node := new(Node)
		node.Label = v
		node.Leader = v
		unionFind.Add(node)
	}

	for v := range vm {
		candidates := generateDistance(v)
		for _, c := range candidates {
			if _, ok := vm[c]; ok {

				n1 := unionFind.Find(v)
				n2 := unionFind.Find(c)
				if n2 != nil && n1.Leader != n2.Leader {
					unionFind.Fuse(n1.Leader, n2.Leader)
				}
			}

		}

	}
	fmt.Println("need ", len(unionFind.Groups), " cluster")
}

func generateDistance(v string) []string {

	vr := []rune(v)

	var result []string
	//for 1 distance
	for i, c := range vr {
		var nvr = make([]rune, len(vr))
		copy(nvr, vr)
		nvr[i] = rev(c)
		result = append(result, string(nvr))
	}

	//for 2 distance
	for i := 0; i < len(vr)-1; i++ {
		for j := i + 1; j < len(vr); j++ {
			var nvr = make([]rune, len(vr))
			copy(nvr, vr)
			nvr[i] = rev(vr[i])
			nvr[j] = rev(vr[j])
			result = append(result, string(nvr))
		}
	}

	return result

}

func rev(c rune) rune {
	if c == '0' {
		return '1'
	}
	return '0'
}
