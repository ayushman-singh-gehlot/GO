package main

import (
	"fmt"
	"ticTacToe/components"
)

func main() {
	// setting player 1 data
	var player components.Player
	fmt.Println("enter name of first player : ")
	firstPlayerName, err := player.Name()
	components.Check(err)
	fmt.Println("select your mark...\nEnter X or O")
	firstPlayerMark, err := player.Mark()
	components.Check(err)
	// setting palyer 2 data
	fmt.Println("enter name of Second player : ")
	secondPlayerName, err := player.Name()
	components.Check(err)
	var secondPlayerMark string
	if firstPlayerMark == "X" {
		secondPlayerMark = "O"
	} else {
		secondPlayerMark = "X"
	}
	// display players and mark
	fmt.Println("\n\nplayer 1 :- (", firstPlayerName, firstPlayerMark, ")")
	fmt.Println("\nplayer 2 :- (", secondPlayerName, secondPlayerMark, ")")

	// setting board size and displaying positions
	var board components.Board
	fmt.Println("\nselect board size...\nEnter 2 for 2x2\nEnter 3 for 3x3\nEnter 4 for 4x4")
	temp, err := board.BoardSize()
	components.Check(err)
	board.Size = temp
	fmt.Printf("board size is %dx%d", board.Size, board.Size)
	board.Intialize()
	components.StartGame(&board, firstPlayerMark, firstPlayerName, secondPlayerMark, secondPlayerName)
}
