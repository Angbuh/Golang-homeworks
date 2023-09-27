package main

import "fmt"

func main() {
	var widht, heigth int
	fmt.Scan(&widht, &heigth)
	if widht == 0 {
		return
	}
	for i := 0; i < heigth; i++ {
		if i%2 == 0 {
			fmt.Print("#")
		}
		for j := 0; j < widht; j++ {
			if i%2 == 0 && j == widht-1 {
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
