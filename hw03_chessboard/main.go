package main

import "fmt"

func main() {
	var length int

	fmt.Scan(&length)
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			fmt.Print("#")
		}
		for j := 0; j < length; j++ {
			if i%2 == 0 && j == length-1 {
				break
			}
			if j%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
