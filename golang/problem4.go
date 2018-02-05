package main

import "fmt"
import "strconv"

func main(){

	result := 0

	for max := 997000 ; max >= 10000 ; max = max - 1000 {
		for left := 999 ; left >= 100 ; left-- {
			for right := left ; right >= 100 ; right-- {
				number := left * right
				if number < max {
					break
				}
				if isPalindrome(strconv.Itoa(number)) {
					result = number
				}
			} 
			if left <= max/1000 {
				break
			}
		}
		if result != 0 {
			break
		}
	}

	fmt.Println("The largest palindrome number of from 3x3 digit multiplication is equal to ", result)
}

func isPalindrome(word string) bool{
	length := len(word)

	for i := 0 ; i < length/2 ; i++ {
		if word[i] != word[length-i-1] {
	        return false;
	    }
 	}
 	return true
}