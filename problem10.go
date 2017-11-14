package main

import "fmt"
import "math"

func main() {

	list := []bool{}

	for i := 0; i < 2000000; i++ {
		list = append(list, true)
	}

	for i := 2; float64(i) < math.Sqrt(2000000); i++ {
		if list[i] {
			prime := i
			for j := prime * prime; j < 2000000; j = j + i {
				list[j] = false
			}
		}
	}

	var result uint64 = 0

	for i := 2; i < len(list); i++ {
		if list[i] {
			result += uint64(i)
		}
	}


	fmt.Println("The result is equal to ", result)
}
