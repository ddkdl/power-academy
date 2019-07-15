package main

import "fmt"

func area(x, y []float64) float64 {
	var result float64

	if len(x) != len(y) {
		return 0
	}

	result = 0

	// Isert logic here

	return result
}

func main() {
	x := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	y := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(area(x, y))
}
