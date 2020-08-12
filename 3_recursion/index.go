package main

import "fmt"

func main() {
	fmt.Println(fibonacciNumbersRec(1, 2, 5))
}

func countDown(i int) int {
	fmt.Println(i)
	if (i <= 0) {
		return 0
	} else {
		return countDown(i - 1);
	}
}

// 1, 1, 2, 3, 5, 8, 13

// xn = xn - 1 + xn - 2

func fibonacciNumbers(num int) {
	x := 1
	y := 2

	for i := 0; i < num; i++ {
		z := x + y;
		x = y
		y = z;
		fmt.Println(y)
	}
}

func fibonacciNumbersRec(x int, y int, n int) int {
	z := x + y;
	x = y
	y = z
	if (n <= 0) {
		return y
	} else {
		return fibonacciNumbersRec(x, y, n - 1)
	}
}