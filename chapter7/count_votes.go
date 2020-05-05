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
	filename=strings.TrimSpace(filename)
	lines,err:=datafile.GetStrings(filename)
	if err!=nil{
		log.Fatal(err)
	}
	//fmt.Println(lines)
	var names [] string
	var vote_count [] int
	for _,line:=range lines{
		matched:=false
		for i,name:=range names{
			if name==line{
				vote_count[i]++
				matched=true
			}
		}
		if matched==false{
			names=append(names,line)
			vote_count=append(vote_count,1)
		}
	}
	for i,name:=range names{
		fmt.Printf("%21s : %d\n",name,vote_count[i])
	}	
}