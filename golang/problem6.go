package main

import "fmt"

func main() {

	result := solve()

	fmt.Println("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is equal to ", result)
}

func solve() int {
	sumofsquares := 0
	squareofsums := 0

	for i := 0; i <= 100; i++ {
		sumofsquares += i * i
		squareofsums += i
	}
	return (squareofsums * squareofsums) - sumofsquares
}
