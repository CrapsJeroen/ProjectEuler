package main

import "fmt"

func main() {

	result := solve()

	fmt.Println("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is equal to ", result)
}

func solve() int {
	primes := []int{2}

	for i := 3; len(primes) <= 10000; i++ {
		fmt.Println(len(primes))
		for _, prime := range primes {
			if i%prime == 0 {
				break
			}
			if prime == primes[len(primes)-1] {
				primes = append(primes, i)
			}
		}
	}
	return primes[len(primes)-1]
}
