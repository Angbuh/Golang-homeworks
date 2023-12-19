package main

import (
	"fmt"

	"github.com/angbuh/golang-homeworks/hw03_chessboard/boardcreator"
)

func main() {
	var size int

	fmt.Print("Введите размер доски: ")
	fmt.Scan(&size)
	fmt.Print(boardcreator.CreateBoard(size))
}
