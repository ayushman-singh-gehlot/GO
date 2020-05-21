package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"ticTacToe/components"
)

type player struct {
	name string
	mark string
}

func main() {
	// setting player 1 data
	var firstPlayer player
	readerObj := bufio.NewReader(os.Stdin)
	fmt.Println("enter name of first player : ")
	name, err := readerObj.ReadString('\n')
	name = strings.TrimSpace(name)
	check(err)
	firstPlayer.name = name
	fmt.Println("select your mark...\nEnter X or O")
	mark, err := readerObj.ReadString('\n')
	check(err)
	mark = strings.TrimSpace(mark)
	if mark != "X" && mark != "O" {
		log.Fatal(errors.New("you did not entered mark"))
	}
	firstPlayer.mark = mark

	// setting palyer 2 data
	var secondPlayer player
	fmt.Println("enter name of Second player : ")
	name, err = readerObj.ReadString('\n')
	name = strings.TrimSpace(name)
	check(err)
	secondPlayer.name = name
	if firstPlayer.mark == "X" {
		secondPlayer.mark = "O"
	} else {
		secondPlayer.mark = "X"
	}
	// display players and mark
	fmt.Println("\n\nplayer 1 :- (", firstPlayer.name, firstPlayer.mark, ")")
	fmt.Println("\nplayer 2 :- (", secondPlayer.name, secondPlayer.mark, ")")

	// setting board size and displaying positions
	var board components.NewBoard
	fmt.Println("\nselect board size...\nEnter 2 for 2x2\nEnter 3 for 3x3\nEnter 4 for 4x4")
	bSize, err := readerObj.ReadString('\n')
	check(err)
	bSize = strings.TrimSpace(bSize)
	size, err := strconv.Atoi(bSize)
	check(err)
	if size != 2 && size != 3 && size != 4 {
		log.Fatal(errors.New("entered wrong size"))
	}
	fmt.Printf("\nboard size is %dx%d\n", size, size)
	board.CreateBoard(size)

	// start game
	for {
		err := errors.New("  ")
		var temp string
		var inp int
		for err != nil {
			fmt.Printf("\nTurn of %s Enter position for your mark\ninput : ", firstPlayer.name)
			temp, err = readerObj.ReadString('\n')
			check(err)
			temp = strings.TrimSpace(temp)
			inp, err = strconv.Atoi(temp)
			if warning(err) {
				continue
			}
			err = board.Update(inp, firstPlayer.mark)
			if warning(err) {
				continue
			}
			board.DisplayBoard()
		}
		valid, result := board.Analyse(firstPlayer.mark)
		if valid && result == firstPlayer.mark {
			fmt.Printf("\n\n--------------%s won--------------\n", firstPlayer.name)
			break
		} else if valid {
			fmt.Println("\n\n--------------Draw--------------")
			break
		}

		err = errors.New("  ")
		for err != nil {
			fmt.Printf("\nTurn of %s Enter position for your mark\ninput : ", secondPlayer.name)
			temp, err = readerObj.ReadString('\n')
			check(err)
			temp = strings.TrimSpace(temp)
			inp, err = strconv.Atoi(temp)
			if warning(err) {
				continue
			}
			err = board.Update(inp, secondPlayer.mark)
			if warning(err) {
				continue
			}
			board.DisplayBoard()
		}
		valid, result = board.Analyse(secondPlayer.mark)
		if valid && result == secondPlayer.mark {
			fmt.Printf("\n\n--------------%s won--------------\n", secondPlayer.name)
			break
		} else if valid {
			fmt.Println("\n\n--------------Draw--------------")
			break
		}
	}

}

func warning(err error) bool {
	if err != nil {
		fmt.Println("Warning : ", err)
		return true
	}
	return false
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
