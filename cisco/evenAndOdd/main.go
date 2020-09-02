package main

import (
	"fmt"
)

func main() {

	var numbers []int

	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
		if numbers[i]%2 == 0 {
			fmt.Printf("%d is even \n", i)
		} else {
			fmt.Printf("%d is odd \n", i)
		}
	}

}
