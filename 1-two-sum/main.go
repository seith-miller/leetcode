// Given an array of integers nums and an integer target, 
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.


package main

import (
	"fmt"
	// "sort"
)

func main()  {
	numbers := []int{3,2,11,7, 8}

	// twoSumNaive(numbers, 9)
	twoSumB(numbers, 9)
}

func twoSumNaive(number []int, target int) (number1 int, number2 int) {
	// fmt.Println(number)
	// fmt.Println(target)
	fmt.Println("")

	for i := 0; i < len(number); i++ {
		fmt.Printf("i is: %v\n", number[i])
		
		for j := i + 1; j < len(number); j++ {
			fmt.Printf("j is: %v\n", number[j])

			if number[i] + number[j] == target {
				sum := number[i] + number[j]
				fmt.Printf("i plus j is: %v\n", sum)
				fmt.Printf("i is %v, j is %v\n", i, j)
				return i, j
			}
		}
	}

	return 0, 0
}

func twoSumB(number []int, target int) (number1 int, number2 int) {
	myMap := make(map[int]int)
	fmt.Println(myMap)

	for i := range(number) {
		fmt.Printf("i is: %v\n", number[i])

		x := target - number[i]
		fmt.Printf("we are looking for: %v", x)
		fmt.Println("")

		for j
	}

	return 0, 0
}