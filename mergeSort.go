package main

import "fmt"

func main() {

	out := mergeSort([]int{100, 1, 2, 6, 4})

	fmt.Println(out)

}

//2-way-mergeSort
func mergeSort(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return arr
	}

	mid := n / 2

	sortedLeft := mergeSort(arr[:mid])
	sortedRight := mergeSort(arr[mid:])

	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {

	i := 0
	j := 0
	n := len(left) + len(right)
	out := make([]int, n)

	for idx := 0; idx < n; idx++ {

		if i == len(left) {
			out[idx] = right[j]
			j++
			continue
		}

		if j == len(right) {
			out[idx] = left[i]
			i++
			continue
		}

		if left[i] >= right[j] {
			out[idx] = right[j]
			j++
		} else {
			out[idx] = left[i]
			i++
		}

	}

	return out

}
