package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getNumber() (int, error) {
	readObject := bufio.NewReader(os.Stdin)
	input, err := readObject.ReadString('\n')
	if err != nil {
		return 0, errors.New("error occured in reading number")
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("user did not entered integer\nplease enter integer again")
	}
	if num < 1 || num > 50 {
		return 0, errors.New("Entered number is not in range 1 to 50 ")
	}
	return num, nil
}

func checkGuess(guess int, target int) bool {
	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
		return false
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
		return false
	} else {
		fmt.Println("GOOD JOB ! you guessed it.")
		return true
	}
}

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(50) + 1
	fmt.Println("I have chosen a random number between 1 and 50.")
	fmt.Println("Can you guess it? ")
	success := false
	for guesses := 0; guesses < 5; {
		fmt.Println("\n\nYou have", 5-guesses, "guesses left.")
		fmt.Print("Make a guess : ")
		guess, err := getNumber()
		if err != nil {
			fmt.Println("Error Occured : ", err)
			continue
		}
		guesses++
		success = checkGuess(guess, target)
		if success {
			fmt.Println("You managed to guess it in", guesses, "chances")
			break
		}

	}
	if !success {
		fmt.Println("Sorry ,you didn't guess my number. It was :", target)
	}

}
