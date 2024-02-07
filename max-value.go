package main

import "fmt"

func maxValue(arrayNums []int) int {
	total := len(arrayNums)
	max := 0

	for nums := 0; nums < total; nums++ {
		if arrayNums[nums] > max {
			max = arrayNums[nums]
		}
	}

	return max
}

func main() {
	array1 := []int{3, 5, 1, 9, 2}
	array2 := []int{10,100,2,4,5,9}
	fmt.Println(maxValue(array1))
	fmt.Println(maxValue(array2))
}