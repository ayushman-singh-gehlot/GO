package components

import (
	"errors"
	"fmt"
	"log"
)

func PlayerTurn(board *Board, name, mark string) (bool, error) {
	fmt.Printf("\nTurn of %s Enter position for your mark\ninput : ", name)
	pos, err := board.Input()
	if err != nil {
		return false, err
		fmt.Println(err)
	}
	board.Values[pos] = mark
	board.DisplayBoard()
	valid, result := board.Analyse(mark)
	if valid && result == mark {
		fmt.Printf("\n\n--------------%s won--------------\n", name)
		return true, err
	} else if valid {
		fmt.Println("\n\n--------------Draw--------------")
		return true, err
	}
	return false, err
}

func StartGame(board *Board, firstPlayerMark, firstPlayerName, secondPlayerMark, secondPlayerName string) {
	// game starts
	for {
		err := errors.New("  ")
		ret := false
		for err != nil {
			ret, err = PlayerTurn(board, firstPlayerName, firstPlayerMark)
			warning(err)
		}
		if ret {
			break
		}
		err = errors.New("  ")
		for err != nil {
			ret, err = PlayerTurn(board, secondPlayerName, secondPlayerMark)
			warning(err)
		}
		if ret {
			break
		}

	}
}

func warning(err error) {
	if err != nil {
		fmt.Println("Warning : ", err)
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
