package main

func MoveRight(board [][]int64) [][]int64 {
	for y := 0; y < len(board); y++ {
		board[y] = moveRowRight(board[y])
	}

	return board
}

func moveRowRight(row []int64) []int64 {
	for i := 0; i < len(row)-1; i++ {
		if row[i+1] == 0 {
			row[i+1] = row[i]
			row[i] = 0
		} else if row[i] == row[i+1] {
			row[i+1] = row[i] + row[i+1]
			row[i] = 0
		}
	}

	return row
}
