package types

import (
	"testing"
)

func TestString(t *testing.T) {
	testdata := Employee{
		UserID:       10,
		Age:          25,
		Name:         "Rob",
		DepartmentID: 3,
	}
	expected := "User ID: 10, Age: 25, Name: Rob, Department ID: 3"

	res := testdata.String()
	if res != expected {
		t.Fatal("error")
	}
}
