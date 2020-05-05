//package datafile allows reading data samples from files.
package datafile

import(
	"os"
	"bufio"
)
// GetStrings read a string from each line of a file.
func GetStrings(filename string) ([] string,error){
	var lines [] string
	file,err:=os.Open(filename)
	if err!=nil{
		//log.Fatal(err)
		return nil,err
	}
	scannerObj:=bufio.NewScanner(file)
	for scannerObj.Scan(){
		line:=scannerObj.Text()
		lines=append(lines,line)
	}
	err=file.Close()
	if err!=nil{
		return nil,err
	}
	if scannerObj.Err()!=nil{
		return nil,scannerObj.Err()
	}
	return lines,nil
}