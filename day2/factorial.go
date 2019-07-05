package main

import "fmt"

func factorial(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	var num int64
	fmt.Print("Please input a number: ")
	fmt.Scan(&num)

	fmt.Printf("%d! = %d\n", num, factorial(num))
}
