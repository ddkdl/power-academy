package main

import "fmt"

func isPrime(number int) bool {
	if number < 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func sieveOfEratostenes(n int) []bool {
	isPrime := make([]bool, n+1)

	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

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
