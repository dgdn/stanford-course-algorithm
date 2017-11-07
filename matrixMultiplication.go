package main

import "fmt"

func main() {

	A := [][]int{
		{1, 2, 3, 4},
		{3, 4, 5, 6},
		{5, 6, 7, 8},
		{6, 7, 8, 9},
	}
	B := [][]int{
		{1, 2, 2, 4},
		{3, 4, 4, 6},
		{5, 6, 6, 8},
		{6, 7, 7, 9},
	}
	out := matrixMultiplication(A, B)
	fmt.Println(out)

}

//assume X Y are square matrix with same dimension
func matrixMultiplication(X, Y [][]int) [][]int {

	//base case
	if len(X) == 1 && len(Y) == 1 {
		return [][]int{{X[0][0] * Y[0][0]}}
	}

	//divide
	n := len(X)
	mid := n / 2
	A := divideMatrix(X, 0, mid, 0, mid)
	B := divideMatrix(X, 0, mid, mid, n)
	C := divideMatrix(X, mid, n, 0, mid)
	D := divideMatrix(X, mid, n, mid, n)

	E := divideMatrix(Y, 0, mid, 0, mid)
	F := divideMatrix(Y, 0, mid, mid, n)
	G := divideMatrix(Y, mid, n, 0, mid)
	H := divideMatrix(Y, mid, n, mid, n)

	//conquer
	P1 := matrixMultiplication(A, matrixSub(F, H))
	P2 := matrixMultiplication(matrixAdd(A, B), H)
	P3 := matrixMultiplication(matrixAdd(C, D), E)
	P4 := matrixMultiplication(D, matrixSub(G, E))
	P5 := matrixMultiplication(matrixAdd(A, D), matrixAdd(E, H))
	P6 := matrixMultiplication(matrixSub(B, D), matrixAdd(G, H))
	P7 := matrixMultiplication(matrixSub(A, C), matrixAdd(E, F))

	//combine
	AEpBG := matrixAdd(matrixSub(matrixAdd(P5, P4), P2), P6)
	AFpBH := matrixAdd(P1, P2)
	CEpDG := matrixAdd(P3, P4)
	CFpDH := matrixSub(matrixSub(matrixAdd(P1, P5), P3), P7)
	return combineMatrix(AEpBG, AFpBH, CEpDG, CFpDH)
}

func combineMatrix(A, B, C, D [][]int) [][]int {
	row := len(A) + len(C)
	out := make([][]int, row)
	for i := 0; i < row; i++ {
		if i < len(A) {
			out[i] = append(out[i], A[i]...)
			out[i] = append(out[i], B[i]...)
		} else {
			out[i] = append(out[i], C[i-len(A)]...)
			out[i] = append(out[i], D[i-len(A)]...)
		}
	}
	return out
}

func matrixAdd(A, B [][]int) [][]int {
	row := len(A)
	col := len(A[0])
	out := make([][]int, row)
	for i := 0; i < row; i++ {
		out[i] = make([]int, col)
		for j := 0; j < col; j++ {
			out[i][j] = A[i][j] + B[i][j]
		}
	}
	return out
}

func matrixSub(A, B [][]int) [][]int {
	row := len(A)
	col := len(A[0])
	out := make([][]int, row)
	for i := 0; i < row; i++ {
		out[i] = make([]int, col)
		for j := 0; j < col; j++ {
			out[i][j] = A[i][j] - B[i][j]
		}
	}
	return out
}

func divideMatrix(A [][]int, rowStart, rowEnd, columnStart, columnEnd int) [][]int {
	row := rowEnd - rowStart
	out := make([][]int, row)
	for i := 0; i < row; i++ {
		out[i] = A[rowStart+i][columnStart:columnEnd]
	}
	return out
}
