package main

import "fmt"

func main() {

	result := solve()
	fmt.Println("The result is equal to ", result)
}

func solve() int {

	for a := 1; a <= 1000; a++ {
		for b := a + 1; b <= (1000 - a); b++ {
			for c := b + 1; c <= (b + a); c++ {
				if c*c > b*b+a*a || a+b+c > 1000 {
					break
				}
				if a*a+b*b == c*c {
					if a+b+c == 1000 {
						return a * b * c
					}
				}

			}
		}
	}

	return -1
}
