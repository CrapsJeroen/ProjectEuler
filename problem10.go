package main

import (
        "fmt"
        "github.com/gonum/matrix/mat64"
)

func main() {
        
        result := 0


        for prime := 2 ; prime < 2000000 ; prime = nextPrime(prime) {
                result += prime
        }

	fmt.Println("The result is equal to ", result)
}

func nextPrime(n int) int{
	for x := n+1 ; x > 0 ; x++ {
		if isPrime(x) {
			return x
		}
	}
	return -1
}

func isPrime(n int) bool{
	if n < 2 {
		return false
	} 
	for i := 2 ; i < n ; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
