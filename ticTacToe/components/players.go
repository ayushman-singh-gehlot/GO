package components

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Player string

func (e Player) Name() (string, error) {
	readerObj := bufio.NewReader(os.Stdin)
	name, err := readerObj.ReadString('\n')
	name = strings.TrimSpace(name)
	return name, err
}

func (e Player) Mark() (string, error) {
	readerObj := bufio.NewReader(os.Stdin)
	mark, err := readerObj.ReadString('\n')
	mark = strings.TrimSpace(mark)
	if mark != "X" && mark != "O" {
		return mark, errors.New("you did not entered mark")
	}
	return mark, err
}
