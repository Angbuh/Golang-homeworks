package main

import (
	"fmt"

	"github.com/angbuh/golang-homeworks/hw07_word_counter/wordcounter"
)

func main() {
	text := "I love my mom very very much"
	words := wordcounter.CountWords(text)
	fmt.Println(words)
}
