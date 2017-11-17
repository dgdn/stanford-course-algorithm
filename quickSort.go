package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var total int

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
	fmt.Println(total, "comparisons")
}

func quickSort(a []int) {

	total += len(a) - 1

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

//implement first-median-last choosing pivot algorithm
func choosePivot(a []int) int {
	first := a[0]

	var middle int
	var middlePosition int
	if len(a)%2 == 0 {
		middlePosition = len(a)/2 - 1
		middle = a[middlePosition]
	} else {
		middlePosition = len(a) / 2
		middle = a[middlePosition]
	}

	last := a[len(a)-1]

	if last < first && first < middle {
		return 0
	}
	if middle < first && first < last {
		return 0
	}

	if first < middle && middle < last {
		return middlePosition
	}

	if last < middle && middle < first {
		return middlePosition
	}

	if first < last && last < middle {
		return len(a) - 1
	}
	if middle < last && last < first {
		return len(a) - 1
	}

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
