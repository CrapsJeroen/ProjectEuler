package main

import "fmt"

func main(){
	fib := 0

	i := 1
	prev := 1
	prevprev := 0

	for i < 4000000 {
		prevprev = prev
		prev = i

		i = prev + prevprev



		if (i%2 == 0){
			fib += i
		} 
	}

	fmt.Println("The sum of all even-valued fibonacci numbers smaller than 4,000,000 is equal to ", fib)
}