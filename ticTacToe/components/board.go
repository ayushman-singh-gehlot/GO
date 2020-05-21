package components

import (
	"errors"
	"fmt"
)

type NewBoard struct {
	size   int
	values []string
}

func (b *NewBoard) Update(inp int, mark string) error {
	if inp >= b.size*b.size || inp < 0 {
		return errors.New("please enter valid position")
	} else if b.values[inp] == "X" || b.values[inp] == "O" {
		return errors.New("place is already marked")
	}
	b.values[inp] = mark
	return nil
}

func (b *NewBoard) CreateBoard(size int) {
	b.size = size
	for i := 0; i < int(b.size*b.size); i++ {
		b.values = append(b.values, " ")
	}
	for i := 0; i < (b.size * b.size); i++ {
		if i%b.size == 0 {
			fmt.Println()
		}
		if i >= b.size*(b.size-1) {
			if i%b.size == (b.size - 1) {
				fmt.Printf("  %d  ", i)
			} else {
				fmt.Printf("  %d  |", i)
			}
		} else {
			if i%b.size == (b.size - 1) {
				fmt.Printf("__%d__", i)
			} else {
				fmt.Printf("__%d__|", i)
			}
		}
	}

}

func (b *NewBoard) DisplayBoard() { // try writing test
	for i := 0; i < (b.size * b.size); i++ {
		if i%b.size == 0 {
			fmt.Println()
		}
		if i >= b.size*(b.size-1) {
			if i%b.size == (b.size - 1) {
				fmt.Printf("  %s  ", b.values[i])
			} else {
				fmt.Printf("  %s  |", b.values[i])
			}
		} else {
			if i%b.size == (b.size - 1) {
				fmt.Printf("__%s__", b.values[i])
			} else {
				fmt.Printf("__%s__|", b.values[i])
			}
		}
	}
}
