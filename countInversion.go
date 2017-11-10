package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//txt file from the course assignment containing 10,00000 integers
	f, err := os.Open("count_inversion.txt")
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

	_, n := sort_and_split(datas)
	fmt.Println(n)
	//output correct answer 2407905288
}

func sort_and_split(data []int) ([]int, int) {

	//base case
	if len(data) == 0 || len(data) == 1 {
		return data, 0
	}

	//divide
	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	//conqur
	a, x := sort_and_split(left)
	b, y := sort_and_split(right)
	z := count_split_inversion(a, b)

	//combine
	return mergeSort(data), x + y + z
}

func count_split_inversion(a, b []int) int {

	i := 0
	j := 0

	total := 0
	n := len(a) + len(b)
	for idx := 0; idx < n; idx++ {

		if j == len(b) {
			break
		}

		if i == len(a) {
			break
		}

		if a[i] > b[j] {
			total += len(a) - i
			j++
		} else {
			i++
		}

	}

	return total
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
			for k := 0; k < len(right)-j; k++ {
				out[idx+k] = right[j+k]
			}
			break
		}

		if j == len(right) {
			for k := 0; k < len(left)-i; k++ {
				out[idx+k] = left[i+k]
			}
			break
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
