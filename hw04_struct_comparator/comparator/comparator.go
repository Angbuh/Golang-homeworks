package comparator

import (
	"errors"

	"github.com/angbuh/golang-homeworks/hw04_struct_comparator/book"
)

type Comparator struct {
	comparationMethod uint
}

const (
	CompareBySize uint = iota
	CompareByYear
	CompareByRate

	unknownComparationMethod
)

var ErrInvalidMethod = errors.New("invalid comparation method")

func NewComparator(comparationMethod uint) (Comparator, error) {
	if comparationMethod >= unknownComparationMethod {
		return Comparator{}, ErrInvalidMethod
	}

	return Comparator{
		comparationMethod: comparationMethod,
	}, nil
}

func (c *Comparator) Compare(book1, book2 *book.Book) bool {
	switch c.comparationMethod {
	case CompareByYear:
		return book1.GetYear() >= book2.GetYear()
	case CompareBySize:
		return book1.GetSize() >= book2.GetSize()
	case CompareByRate:
		return book1.GetRate() >= book2.GetRate()
	}

	panic(ErrInvalidMethod)
}
