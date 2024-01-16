package book

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	testdata := Book{
		ID:     23,
		Title:  "Tree",
		Author: "Kate God",
		Year:   1999,
		Size:   10,
		Rate:   7.5,
	}
	expected := `{"id":23,"title":"Tree","author":"Kate God","year":1999,"size":10,"rate":7.5}`
	res, err := json.Marshal(&testdata)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(res))
}

func TestUnmarshalJSON(t *testing.T) {
	testdata := `{"Id":23,"Title":"Tree","Author":"Kate God","Year":1999,"Size":10,"Rate":7.5}`
	expected := Book{23, "Tree", "Kate God", 1999, 10, 7.5}
	var data Book
	err := json.Unmarshal([]byte(testdata), &data)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}

func TestMarshalJSONslice(t *testing.T) {
	testdata := []Book{
		{
			ID:     23,
			Title:  "Tree",
			Author: "Kate God",
			Year:   1999,
			Size:   10,
			Rate:   7.5,
		},
		{
			ID:     23,
			Title:  "Flower",
			Author: "Kate Bad",
			Year:   2001,
			Size:   12,
			Rate:   7.0,
		},
	}
	expected := `[{"id":23,"title":"Tree","author":"Kate God","year":1999,"size":10,"rate":7.5},
{"id":23,"title":"Flower","author":"Kate Bad","year":2001,"size":12,"rate":7}]`
	res, err := json.Marshal(&testdata)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(res))
}

func TestUnmarshalJSONslice(t *testing.T) {
	testdata := `[{"id":23,"Title":"Tree","Author":"Kate God","Year":1999,"Size":10,"Rate":7.5},
{"id":23,"Title":"Flower","Author":"Kate Bad","Year":2001,"Size":12,"Rate":7}]`
	expected := []Book{{23, "Tree", "Kate God", 1999, 10, 7.5}, {23, "Flower", "Kate Bad", 2001, 12, 7.0}}
	var data []Book
	err := json.Unmarshal([]byte(testdata), &data)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
}
