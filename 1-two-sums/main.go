// Given an array of integers nums and an integer target, 
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.


package main

import "fmt"

func main()  {
	numbers := []int{3,2,11,7}

	twoSum(numbers, 9)
}

func twoSum(number []int, target int) (number1 int, number2 int) {
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