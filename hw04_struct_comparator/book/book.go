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
