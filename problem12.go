package main

import "fmt"

func main() {

	divisors := []int{}
	triangleNum := 0
	for i := 1; len(divisors) < 500; i++ {
		triangleNum += i
		divisors = []int{}
		for j := 1; j <= triangleNum; j++ {
			if triangleNum%j == 0 {
				divisors = append(divisors, j)
				divisors = append(divisors, triangleNum/j)
				if triangleNum/j-j <= 1 {
					break
				}
			}
		}
	}

	fmt.Println("The result is equal to ", triangleNum)
}
