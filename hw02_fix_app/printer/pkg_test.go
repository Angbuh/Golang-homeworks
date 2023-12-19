package printer

import (
	"testing"

	"github.com/angbuh/golang-homeworks/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestStaffToString(t *testing.T) {
	data := []types.Employee{
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

	expected := "{User ID: 10, Age: 25, Name: Rob, Department ID: 3; User ID: 11, Age: 30, Name: George, Department ID: 2}"

	assert.Equal(t, expected, StaffToString(data))
	assert.Equal(t, "{}", StaffToString([]types.Employee{}))
}
