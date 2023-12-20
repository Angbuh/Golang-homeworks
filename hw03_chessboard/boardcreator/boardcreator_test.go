package boardcreator

import "testing"

func TestCreateBoard_ZeroSize(t *testing.T) {
	size := 0
	board := CreateBoard(size)

	if board != "" {
		t.Fatal("an empty string is expected")
	}
}

func TestCreateBoard_NegotiveSize(t *testing.T) {
	sizes := []int{-1, -5, -1000}

	for _, size := range sizes {
		board := CreateBoard(size)

		if board != "" {
			t.Fatal("an empty string is expected")
		}
	}
}

func TestCreateBoard_1Size(t *testing.T) {
	size := 1
	expectedBoard := "#\n"

	if actualBoard := CreateBoard(size); actualBoard != expectedBoard {
		t.Fatalf("\"%v\" != \"%v\"", []rune(actualBoard), []rune(expectedBoard))
	}
}

func TestCreateBoard_10Size(t *testing.T) {
	size := 10
	expectedBoard := `# # # # # 
 # # # # #
# # # # # 
 # # # # #
# # # # # 
 # # # # #
# # # # # 
 # # # # #
# # # # # 
 # # # # #
`

	if actualBoard := CreateBoard(size); actualBoard != expectedBoard {
		t.Fatalf("\"%v\" != \"%v\"", []rune(actualBoard), []rune(expectedBoard))
	}
}

func TestCreateBoard_3Size(t *testing.T) {
	size := 3
	expectedBoard := `# #
 # 
# #
`

	if actualBoard := CreateBoard(size); actualBoard != expectedBoard {
		t.Fatalf("\"%v\" != \"%v\"", []rune(actualBoard), []rune(expectedBoard))
	}
}
