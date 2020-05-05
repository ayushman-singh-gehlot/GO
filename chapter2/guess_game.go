package main

import(
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"	
)

func main(){
	seconds:=time.Now().Unix()
	rand.Seed(seconds)
	target:=rand.Intn(100)+1
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it? ")
	//fmt.Println(target)
	read_object:=bufio.NewReader(os.Stdin)
	success:=false
	for guesses:=0;guesses<10;guesses++{
		fmt.Println("\n\nYou have",10-guesses,"guesses left.")
		fmt.Print("Make a guess : ")
		input,err:=read_object.ReadString('\n')
		if err!=nil{
			log.Fatal(err)
		}
		input=strings.TrimSpace(input)
		guess,err:=strconv.Atoi(input)
		if err!=nil{
			log.Fatal(err)
		}
		if guess < target{
			fmt.Println("Oops. Your guess was LOW.")
		}else if guess>target{
			fmt.Println("Oops. Your guess was HIGH.")
		}else{
			success=true
			fmt.Println("GOOD JOB ! you guessed it.")
			break
		}
	}
	if !success {
		fmt.Println("Sorry ,you didn't guess my number. It was :",target)
	}
	
}