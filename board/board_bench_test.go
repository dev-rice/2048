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
		NewBoardFromGrid(board)
	}
}

func BenchmarkUncompressBoard(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{0, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 2},
	})

	for n := 0; n < b.N; n++ {
		ExtractGridFromBoard(board)
	}
}

func BenchmarkAreMovesLeft(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{2, 4, 2, 4},
		{4, 2, 4, 2},
		{2, 4, 2, 4},
		{4, 2, 4, 0},
	})

	for n := 0; n < b.N; n++ {
		AreMovesLeft(board)
	}
}

func BenchmarkMoveRight(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	})

	for n := 0; n < b.N; n++ {
		MoveRight(board)
	}
}

func BenchmarkMoveLeft(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	})

	for n := 0; n < b.N; n++ {
		MoveLeft(board)
	}
}

func BenchmarkMoveDown(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	})

	for n := 0; n < b.N; n++ {
		MoveDown(board)
	}
}

func BenchmarkMoveUp(b *testing.B) {
	board := NewBoardFromGrid([][]int64{
		{2, 128, 0, 0},
		{4, 0, 0, 4},
		{2, 0, 8, 8},
		{2, 2, 2, 32},
	})

	for n := 0; n < b.N; n++ {
		MoveUp(board)
	}
}
