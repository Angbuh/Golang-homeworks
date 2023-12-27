package main

import (
	"encoding/json"
)

// "fmt"

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"autor"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, b)
}

func main() {
	book1 := &Book{
		Id:     23,
		Title:  "Tree",
		Author: "Kate God",
		Year:   1999,
		Size:   10,
		Rate:   7.5,
	}

}
