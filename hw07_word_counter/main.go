package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw07_word_counter/wordcounter"
)

func main() {
	text := "I love my mom very very much"
	words := wordcounter.CountWords(text)
	fmt.Println(words)
}
