package main

import "fmt"

func factorial(nums int) int {
	if nums == 0 {
		return 1
	}
	return nums * factorial(nums-1)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
	fmt.Println(factorial(7))
	fmt.Println(factorial(8))
	fmt.Println(factorial(9))
	fmt.Println(factorial(10))
}