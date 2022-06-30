package main

import (
    "fmt"
)

func main() {
	fullset := 
	numbers1 := [3]int{1, 2, 0}
	numbers2 := [4]int{3, 4, -1, 1}
	numbers3 := [5]int{7, 8, 9, 11,12}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println(numbers3)

	for i := 0; i < len(numbers1); i ++ {
		fmt.Println(numbers1[i])
	} 
}