package comparator

import (
	"errors"
	"testing"

	"github.com/angbuh/golang-homeworks/hw04_struct_comparator/book"
)

func TestCompare(t *testing.T) {
	books := []*book.Book{
		book.New(23, "Tree", "Kate God", 1999, 10, 7.5),
		book.New(54, "Tree", "Kate God", 2001, 12, 7.0),
	}

	comparator, err := NewComparator(CompareByYear)
	if err != nil {
		t.Fatal(err)
	}

	if !comparator.Compare(books[1], books[0]) {
		t.Fatal("unexpected Compare() result")
	}

	comparator, err = NewComparator(CompareBySize)
	if err != nil {
		t.Fatal(err)
	}
	if !comparator.Compare(books[1], books[0]) {
		t.Fatal("unexpected Compare() result")
	}

	comparator, err = NewComparator(CompareByRate)
	if err != nil {
		t.Fatal(err)
	}

	if comparator.Compare(books[1], books[0]) {
		t.Fatal("unexpected Compare() result")
	}

	comparator.comparationMethod = unknownComparationMethod

	defer func() {
		if err := recover(); errors.Is(err.(error), ErrInvalidMethod) {
			t.Fatal("the code did not panic")
		}
	}()

	comparator.Compare(books[1], books[0])
}

func TestNewComparator(t *testing.T) {
	var comparator Comparator
	compareMethod := unknownComparationMethod

	_, err := NewComparator(compareMethod)
	if err == nil {
		t.Fatal("error is expected")
	}

	for compareMethod = 0; compareMethod < unknownComparationMethod; compareMethod++ {
		if comparator, err = NewComparator(compareMethod); err != nil {
			t.Fatal("nil error is expected")
		} else if comparator.comparationMethod != compareMethod {
			t.Fatal("comparator method is now equal to the expected one")
		}
	}
}
