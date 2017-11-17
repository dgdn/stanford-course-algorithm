package main

import "fmt"

func main() {
	a := []int{3, 8, 2, 5, 1, 4, 7, 6}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(a []int) {

	//base case
	if len(a) == 1 {
		return
	}

	//choose pivot
	pivot := choosePivot(a)

	//patition
	partition(pivot, a)

	//recursively sort
	if pivot-1 > 0 {
		quickSort(a[:pivot-1])
	}
	if pivot+1 < len(a) {
		quickSort(a[pivot+1:])
	}
}

func choosePivot(a []int) int {
	return 0
}

func partition(pivot int, a []int) {
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
}
