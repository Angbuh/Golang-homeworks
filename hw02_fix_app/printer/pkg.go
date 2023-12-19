package printer

import (
	"github.com/angbuh/golang-homeworks/hw02_fix_app/types"
)

func StaffToString(staff []types.Employee) string {
	result := "{"

	for i := range staff {
		result += staff[i].String()
		if i != len(staff)-1 {
			result += "; "
		}
	}

	result += "}"
	return result
}
