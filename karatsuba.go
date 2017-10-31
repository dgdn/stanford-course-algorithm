package main

import (
	"fmt"
	"strconv"
)

//class 1 test answer
func main() {

	a := "3141592653589793238462643383279502884197169399375105820974944592"
	b := "2718281828459045235360287471352662497757247093699959574966967627"

	x := make([]int, len(a))
	y := make([]int, len(b))
	for i, v := range a {
		x[i], _ = strconv.Atoi(string(v))
	}
	for i, v := range b {
		y[i], _ = strconv.Atoi(string(v))
	}

	out := multiplication(x, y)
	var str string
	//remove padding zero
	for _, v := range removePaddingZero(out) {
		str += fmt.Sprint(v)
	}

	fmt.Println(str)
	//output: 8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184

}

//implement karatsuba multiplication
func multiplication(x, y []int) []int {

	//remove redundent padding 0 first
	x = removePaddingZero(x)
	y = removePaddingZero(y)

	//ensure x y has the same even len
	//TODO this restriction may not true
	var n int
	if len(x) > len(y) {
		n = len(x)
	} else {
		n = len(y)
	}
	if n > 1 {
		if n%2 != 0 {
			n++
		}
		x = leftPaddingZero(x, n)
		y = leftPaddingZero(y, n)
	}

	//define base case
	if n == 1 {
		d := x[0] * y[0]
		l := d / 10
		r := d - l*10
		return []int{l, r}
	}

	xm := len(x) / 2
	ym := len(y) / 2

	a := x[:xm]
	b := x[xm:]
	c := y[:ym]
	d := y[ym:]

	ac := multiplication(a, c)
	bd := multiplication(b, d)

	s3 := multiplication(add(a, b), add(c, d))

	adPlusBc := sub(sub(s3, ac), bd)

	return add(add(pad(ac, n), pad(adPlusBc, n/2)), bd)
}

func removePaddingZero(a []int) []int {
	i := 0
	//keep at least one zero
	for ; i < len(a)-1; i++ {
		if a[i] != 0 {
			break
		}
	}
	return a[i:]
}

func leftPaddingZero(a []int, size int) []int {
	n := size - len(a)
	out := make([]int, size)
	for i := 0; i < n; i++ {
		out[i] = 0
	}
	for i := 0; i < len(a); i++ {
		out[i+n] = a[i]
	}
	return out
}

func add(a, b []int) []int {

	alen := len(a)
	blen := len(b)

	var size int
	if alen > blen {
		size = alen
	} else {
		size = blen
	}

	//+1 make room large enough for carry bit
	size = size + 1

	out := make([]int, size)
	carry := 0

	var av, bv int
	for i := 0; i < size; i++ {

		if blen-i-1 < 0 {
			bv = 0
		} else {
			bv = b[blen-i-1]
		}

		if alen-i-1 < 0 {
			av = 0
		} else {
			av = a[alen-i-1]
		}

		c := av + bv + carry
		if c >= 10 {
			out[size-i-1] = c - 10
			carry = 1
		} else {
			out[size-i-1] = c
			carry = 0
		}

	}

	return out
}

func sub(a, b []int) []int {

	//assume a >= b

	alen := len(a)
	blen := len(b)

	size := alen
	carry := 0

	out := make([]int, size)

	var av, bv int
	for i := 0; i < size; i++ {
		av = a[size-i-1]

		if blen-i-1 < 0 {
			bv = 0
		} else {
			bv = b[blen-i-1]
		}

		if av-carry >= bv {
			out[size-i-1] = av - carry - bv
			carry = 0
		} else {
			out[size-i-1] = av - carry + 10 - bv
			carry = 1
		}

	}

	return out
}

func pad(a []int, size int) []int {

	//copy to avoid potential slice reference
	out := make([]int, len(a)+size)
	for i := 0; i < len(a); i++ {
		out[i] = a[i]
	}

	for i := len(a); i < size; i++ {
		out[i] = 0
	}

	return out
}
