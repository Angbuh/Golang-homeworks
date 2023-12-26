package main

import (
	"fmt"

	"github.com/angbuh/golang-homeworks/hw08_binary_search/binarysearch"
)

func main() {
	myarr := []int{1, 3, 7, 9, 10, 13, 107}

	fmt.Println(binarysearch.BinarySearch(myarr, 3))
}
