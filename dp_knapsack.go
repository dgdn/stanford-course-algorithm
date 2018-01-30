package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	f_small, err := os.Open("knapsack1.txt")
	if err != nil {
		panic(err)
	}
	f_big, err := os.Open("knapsack_big.txt")
	if err != nil {
		panic(err)
	}

	small_datas, small_capacity := parseData(f_small)
	big_datas, big_capacity := parseData(f_big)

	small_value := knapsack_iterative(small_datas, small_capacity)
	fmt.Println(small_value)

	cache := make(map[string]int)
	big_value := knapsack_recursive(big_datas, big_capacity, cache)
	fmt.Println(big_value)

}

func parseData(r io.Reader) ([]Object, int) {
	scanner := bufio.NewScanner(r)
	var datas []Object
	var capacity int
	var err error
	for i := 0; scanner.Scan(); i++ {
		row := strings.Fields(scanner.Text())
		if i == 0 {
			capacity, err = strconv.Atoi(row[0])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			value, err := strconv.Atoi(row[0])
			if err != nil {
				fmt.Println(err)
			}
			weight, err := strconv.Atoi(row[1])
			if err != nil {
				fmt.Println(err)
			}
			datas = append(datas, Object{value, weight})
		}
	}
	return datas, capacity
}

type Object struct {
	Value  int
	Weight int
}

func knapsack_iterative(datas []Object, capacity int) int {

	A := make([][]int, len(datas)+1)
	for i := 0; i < len(A); i++ {
		A[i] = make([]int, capacity+1)
	}
	for v := 0; v <= capacity; v++ {
		A[0][v] = 0
	}
	for i := 1; i < len(A); i++ {
		wi := datas[i-1].Weight
		vi := datas[i-1].Value
		for w := 0; w <= capacity; w++ {
			if wi > w {
				A[i][w] = A[i-1][w]
			} else {
				if A[i-1][w] > A[i-1][w-wi]+vi {
					A[i][w] = A[i-1][w]
				} else {
					A[i][w] = A[i-1][w-wi] + vi
				}
			}
		}
	}
	return A[len(datas)][capacity]
}

func knapsack_recursive(datas []Object, capacity int, cache map[string]int) (out int) {

	//base case
	if len(datas) == 1 {
		if datas[0].Weight <= capacity {
			return datas[0].Value
		}
		return 0
	}

	lastIdx := len(datas) - 1

	//case 2 not include the last item
	case2Value := cache_knapsack_recursive(datas[:lastIdx], capacity, cache)

	//edge case return case 2 value directly
	if capacity-datas[lastIdx].Weight < 0 {
		return case2Value
	}

	//case 1 include the last item
	case1Value := cache_knapsack_recursive(datas[:lastIdx], capacity-datas[lastIdx].Weight, cache) + datas[lastIdx].Value
	if case1Value > case2Value {
		return case1Value
	}
	return case2Value
}

func cache_knapsack_recursive(datas []Object, capacity int, cache map[string]int) int {
	lastIdx := len(datas) - 1
	key := fmt.Sprintf("%d-%d", lastIdx, capacity)
	if v, ok := cache[key]; ok {
		return v
	}
	v := knapsack_recursive(datas, capacity, cache)
	cache[key] = v
	return v
}

func print2dArray(datas [][]int) {

	i_length := len(datas)
	j_length := len(datas[0])

	for j := j_length - 1; j > 0; j-- {
		for i := 0; i < i_length; i++ {
			fmt.Print(datas[i][j], " ")
		}
		fmt.Print("\n")
	}
}
