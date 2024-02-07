package main

import "fmt"

func triangleStars(stars int) {
	for rows := 1; rows <= stars; rows++ {
		for cols := 0; cols < rows; cols++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	triangleStars(10)
}