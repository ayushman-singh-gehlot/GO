package main

import (
	"fmt"
	"bufio"
	"os"
	"packages/datafile"
	"strings"
	"log"
)

func main(){
	readerObj:=bufio.NewReader(os.Stdin)
	filename,err:=readerObj.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	filename=strings.TrimSpace(filename) // triming was required ,without triming filename was not recognized by func GetStrings
	lines,err:=datafile.GetStrings(filename)
	if err!=nil{
		log.Fatal(err)
	}

	vote_count:=make(map[string]int)

	for _,line:=range lines{
		vote_count[line]++   // incrementing count of the name present
							 // and for name not present it will first declare key then increment count
	}
	for name,votes:=range vote_count{
		fmt.Printf("%21s : %d\n",name,votes)
	}	
}