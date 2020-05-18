package main

import (
	"fmt"
	"chapter14/prose"
)

func main() {
	phrases := []string{"apple", "mango"}
	fmt.Println(prose.JoinWithCommas(phrases))
	phrases = []string{"apple", "oranges", "mango"}
	fmt.Println(prose.JoinWithCommas(phrases))
}
