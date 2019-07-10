package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func estimatePi(n int) float64 {
	var x float64
	var y float64
	var pointsInside float64
	var totalPoints float64

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := 0; i < n; i++ {
		x = rng.Float64()
		y = rng.Float64()

		distance := math.Sqrt(x*x + y*y)

		if distance <= 1.0 {
			pointsInside++
		}
		totalPoints++
	}

	return 4.0 * pointsInside / totalPoints
}

func main() {
	fmt.Println(estimatePi(1000000))
}
