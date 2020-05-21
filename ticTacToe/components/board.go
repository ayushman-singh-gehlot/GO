package components

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Size   int
	Values []string
}

func (b Board) BoardSize() (int, error) {
	readerObj := bufio.NewReader(os.Stdin)
	bSize, err := readerObj.ReadString('\n')
	bSize = strings.TrimSpace(bSize)
	size, err := strconv.Atoi(bSize)

	if size != 2 && size != 3 && size != 4 {
		return 0, errors.New("entered wrong size")
	}
	return size, err
}

func (b Board) Input() (int, error) {
	readerObj := bufio.NewReader(os.Stdin)
	temp, err := readerObj.ReadString('\n')
	temp = strings.TrimSpace(temp)
	inp, err := strconv.Atoi(temp)
	if err != nil {
		return 0, errors.New("please enter integer value")
	} else if inp >= b.Size*b.Size || inp < 0 {
		return 0, errors.New("enter a valid position")
	} else if b.Values[inp] == "X" || b.Values[inp] == "O" {
		return 0, errors.New("place is already marked")
	}
	return inp, err
}

func (b *Board) Intialize() {
	for i := 0; i < int(b.Size*b.Size); i++ {
		b.Values = append(b.Values, " ")
	}
	for i := 0; i < (b.Size * b.Size); i++ {
		if i%b.Size == 0 {
			fmt.Println()
		}
		if i >= b.Size*(b.Size-1) {
			if i%b.Size == (b.Size - 1) {
				fmt.Printf("  %d  ", i)
			} else {
				fmt.Printf("  %d  |", i)
			}
		} else {
			if i%b.Size == (b.Size - 1) {
				fmt.Printf("__%d__", i)
			} else {
				fmt.Printf("__%d__|", i)
			}
		}
	}

}

func (b *Board) DisplayBoard() {
	for i := 0; i < (b.Size * b.Size); i++ {
		if i%b.Size == 0 {
			fmt.Println()
		}
		if i >= b.Size*(b.Size-1) {
			if i%b.Size == (b.Size - 1) {
				fmt.Printf("  %s  ", b.Values[i])
			} else {
				fmt.Printf("  %s  |", b.Values[i])
			}
		} else {
			if i%b.Size == (b.Size - 1) {
				fmt.Printf("__%s__", b.Values[i])
			} else {
				fmt.Printf("__%s__|", b.Values[i])
			}
		}
	}
}
