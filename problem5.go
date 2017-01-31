package main

import "fmt"

func main() {

	result := solve()

	fmt.Println("The smallest positive number that is evenly divisible by number 1 to 20 is equal to ", result)
}

func solve() int {
	max := 20

	list := []int{11, 13, 14, 16, 17, 18, 19, 20}

	for result := 20; result <= 1000000000; result = result + 20 {
		for _, number := range list {
			if result%number != 0 {
				break
			}
			if number == max {
				return result
			}
		}
	}
	return -1
}
