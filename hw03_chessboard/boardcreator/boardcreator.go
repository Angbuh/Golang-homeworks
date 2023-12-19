package boardcreator

func CreateBoard(size int) string {
	var board []rune
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			board = append(board, '#')
		}
		for j := 0; j < size; j++ {
			if i%2 == 0 && j == size-1 {
				break
			}
			if j%2 == 0 {
				board = append(board, ' ')
			} else {
				board = append(board, '#')
			}
		}
		board = append(board, '\n')
	}

	return string(board)
}
