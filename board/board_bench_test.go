package board

import "testing"

func BenchmarkCompressBoardGrid(b *testing.B) {
	board := [][]int64{
		{0, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}

	for n := 0; n < b.N; n++ {
		CompressBoardGrid(board)
	}
}

func BenchmarkUncompressBoard(b *testing.B) {
	board := [][]int64{
		{0, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	}
	compressedBoard := CompressBoardGrid(board)
	for n := 0; n < b.N; n++ {
		UncompressBoard(compressedBoard)
	}
}

func BenchmarkAreMovesLeft(b *testing.B) {
	board := [][]int64{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 0},
	}
	compressedBoard := CompressBoardGrid(board)
	for n := 0; n < b.N; n++ {
		AreMovesLeft(compressedBoard)
	}
}

func BenchmarkMoveRight(b *testing.B) {
	board := [][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	}
	compressedBoard := CompressBoardGrid(board)

	for n := 0; n < b.N; n++ {
		MoveRight(compressedBoard)
	}
}
