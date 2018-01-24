package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("mwis.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var datas []int
	for i := 0; scanner.Scan(); i++ {
		//skip the first number
		if i == 0 {
			continue
		}
		data := scanner.Text()
		d, err := strconv.Atoi(data)
		if err != nil {
			continue
		}
		datas = append(datas, d)
	}
	m := wis(datas)

	checkData := []int{1, 2, 3, 4, 17, 117, 517, 997}
	var out string
	for _, d := range checkData {
		if m[d] {
			out += "1"
		} else {
			out += "0"
		}
	}
	fmt.Println(out)

}

func wis(W []int) []bool {

	A := make([]int, len(W)+1)
	A[0] = 0
	A[1] = W[0]

	for i := 2; i < len(A); i++ {

		if A[i-1] > A[i-2]+W[i-1] {
			A[i] = A[i-1]
		} else {
			A[i] = A[i-2] + W[i-1]
		}

	}

	//construction process
	m := make([]bool, len(A))
	for i := len(A) - 1; i >= 1; {
		if i == 1 {
			m[i] = true
			break
		}

		if A[i-1] > A[i-2]+W[i-1] {
			i = i - 1
		} else {
			m[i] = true
			i = i - 2
		}
	}

	return m

}
