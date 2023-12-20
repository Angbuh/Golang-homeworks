package main

import (
	"github.com/angbuh/golang-homeworks/hw04_struct_comparator/book"
	"github.com/angbuh/golang-homeworks/hw04_struct_comparator/comparator"
)

func main() {
	books := []*book.Book{
		book.New(23, "Tree", "Kate God", 1999, 10, 7.5),
		book.New(54, "Tree", "Kate God", 2001, 12, 7.0),
	}

	comparator, err := comparator.NewComparator(comparator.CompareByYear)
	if err != nil {
		panic(err)
	}

	if !comparator.Compare(books[1], books[0]) {
		panic("unexpected Compare() result")
	}
}
