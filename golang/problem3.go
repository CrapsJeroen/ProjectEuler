package main

import "fmt"

func main(){
	prime := 2
	number := 600851475143

	for number != 1 {
		if (number%prime == 0){
			number = number/prime
		} else {
			prime = nextPrime(prime)
		} 
	}

	fmt.Println("The largest prime number of 600851475143 is equal to ", prime)
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