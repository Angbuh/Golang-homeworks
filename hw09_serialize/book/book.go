package book

import (
	"encoding/json"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

type _book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(&_book{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var book _book
	err := json.Unmarshal(data, &book)

	b.ID = book.ID
	b.Title = book.Title
	b.Author = book.Author
	b.Year = book.Year
	b.Size = book.Size
	b.Rate = book.Rate
	return err
}
