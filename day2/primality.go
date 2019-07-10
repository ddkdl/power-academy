package main

import "fmt"

func isPrime(number int) bool {
	// Insert logic here

	return true
}

func sieveOfEratostenes(n int) []bool {
	isPrime := make([]bool, n+1)

	// Insert logic here

	return isPrime
}

func main() {
	var limit int

	fmt.Print("Please input the upper limit on numbers: ")
	fmt.Scan(&limit)

	for i := 2; i < limit; i++ {
		if isPrime(i) {
			fmt.Println(i, "is prime.")
		} else {
			fmt.Println(i, "is not prime.")
		}
	}

	primes := sieveOfEratostenes(limit)

	for number, prime := range primes {
		if number < 2 {
			continue
		}

		if prime {
			fmt.Println(number, "is prime.")
		} else {
			fmt.Println(number, "is not prime.")
		}
	}
}
