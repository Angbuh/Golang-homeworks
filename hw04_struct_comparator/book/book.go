package book

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

type Comparator struct {
	comparationMethod int
}

const (
	CompareBySize int = iota
	CompareByYear
	CompareByRate

	unknownComparationMethod
)

func New(id int, title string, author string, year int, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func NewComparator(comparationMethod int) (*Comparator, error) {
	if comparationMethod < CompareBySize || comparationMethod >= unknownComparationMethod {
		return nil, fmt.Errorf("invalid comparation method")
	}

	return &Comparator{
		comparationMethod: comparationMethod,
	}, nil
}

func (c *Comparator) Compare(book1, book2 *Book) bool {
	switch c.comparationMethod {
	case CompareByYear:
		return book1.year >= book2.year
	case CompareBySize:
		return book1.size >= book2.size
	case CompareByRate:
		return book1.rate >= book2.rate
	}

	panic("invalid comparation method")
}

func (b *Book) GetInfo() string {
	return fmt.Sprintf(`
		ID: %d; 
		Title: %s; 
		Author: %s; 
		Year: %d;
		Size: %d; 
		Rate: %f`, b.id, b.title, b.author, b.year, b.size, b.rate,
	)
}
