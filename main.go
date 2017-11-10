package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	board := NewEmptyBoard()
	board[0][3] = 2
	board[0][1] = 2
	board[0][0] = 2048

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" { // break the loop if text == "q"
		fmt.Println("")
		fmt.Println(BoardToString(board))
		fmt.Println("")

		fmt.Print("Enter your move: ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			if text == "r" {
				board = MoveRight(board)
			}
		}
	}

}
