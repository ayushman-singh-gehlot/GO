package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var broken string="G# r#cks!"
	temp:= strings.NewReplacer("#","o")  
	var fixed string =temp.Replace(broken)
	fmt.Println(reflect.TypeOf(temp),"\t",temp,"\n",fixed)

}
// methods belong to a individual value 
// ex Replace method belong to value(temp in this case) of a particular type *strings.Replacer
//for more info refer page 71