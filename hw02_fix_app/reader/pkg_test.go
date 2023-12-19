package reader

import (
	"errors"
	"io"
	"testing"

	"github.com/angbuh/golang-homeworks/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

type ErrorReader struct{}

type EmployeeReader struct {
	reader *employeeReader
}

type employeeReader struct {
	bytes []byte
}

func (f ErrorReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("expected error")
}

func (f EmployeeReader) Read(p []byte) (n int, err error) {
	readBytesCount := min(cap(p), len(f.reader.bytes))

	copy(p, f.reader.bytes[:readBytesCount])
	f.reader.bytes = f.reader.bytes[readBytesCount:]

	if readBytesCount == 0 {
		return 0, io.EOF
	}

	return readBytesCount, nil
}

func TestReadJsonFromReaderReadError(t *testing.T) {
	_, err := readJSONFromReader(ErrorReader{})

	assert.NotNil(t, err)

	_, err = ReadJSON("")
	assert.NotNil(t, err)
}

func TestReadJsonFromReaderUnmarshalError(t *testing.T) {
	reader := employeeReader{
		bytes: []byte("{"),
	}

	_, err := readJSONFromReader(EmployeeReader{&reader})
	assert.NotNil(t, err)
}

func TestReadJsonFromReaderSuccess(t *testing.T) {
	reader := employeeReader{
		bytes: []byte(`
		[
			{
				"userId": 10,
				"age": 25,
				"name": "Rob",
				"departmentId": 3
			},
			{
				"userId": 11,
				"age": 30,
				"name": "George",
				"departmentId": 2
			}
		]
		`),
	}
	expected := []types.Employee{
		{
			UserID:       10,
			Age:          25,
			Name:         "Rob",
			DepartmentID: 3,
		},
		{
			UserID:       11,
			Age:          30,
			Name:         "George",
			DepartmentID: 2,
		},
	}

	res, err := readJSONFromReader(EmployeeReader{&reader})

	assert.Nil(t, err)
	assert.Equal(t, expected, res)

	res, err = readJSONFromReader(EmployeeReader{&employeeReader{bytes: []byte("[]")}})

	assert.Nil(t, err)
	assert.Equal(t, []types.Employee{}, res)
}
