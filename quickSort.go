package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("quick_sort_data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var datas []int
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		datas = append(datas, d)
	}
	quickSort(datas)
	fmt.Println(datas)
}

func quickSort(a []int) {

	//base case
	if len(a) == 1 {
		return
	}

	//choose pivot
	pivot := choosePivot(a)

	//patition
	p := partition(pivot, a)

	//recursively sort
	if p-1 > 0 {
		quickSort(a[:p])
	}
	if p+1 < len(a) {
		quickSort(a[p+1:])
	}
}

func choosePivot(a []int) int {
	return 0
}

func partition(pivot int, a []int) int {
	//place pivot to the left most
	p := a[pivot]
	a[pivot] = a[0]
	a[0] = p

	var i = 1
	for j := 1; j < len(a); j++ {
		if p > a[j] {
			tmp := a[j]
			a[j] = a[i]
			a[i] = tmp
			i++
		}
	}
	a[0] = a[i-1]
	a[i-1] = p

	return i - 1
}
