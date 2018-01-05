package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("2sum.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := HashTable{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		d, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		m.Put(d)
	}

	d := map[int]bool{}
	m.Range(func(k int) {
		for i := -10000; i <= 10000; i++ {
			if m.Exist(i - k) {
				d[i] = true
			}
		}
	})
	fmt.Println(len(d))

}

type HashTable map[int]bool

func (h HashTable) Put(d int) {
	h[d] = false
}

func (h HashTable) Exist(d int) bool {
	_, ok := h[d]
	if ok {
		h[d] = true
	}
	return ok
}

func (h HashTable) Range(f func(int)) {
	for k := range h {
		if !h[k] {
			h[k] = true
			f(k)
		}
	}
}
