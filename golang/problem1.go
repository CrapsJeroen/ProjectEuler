package main

import "fmt"

func main(){
	sum := 0

	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		} 
	}

	fmt.Println("The sum of all multiples of 3 and 5 smaller than 1000 is equal to ", sum)
}