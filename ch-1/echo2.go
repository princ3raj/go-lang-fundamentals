package main

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _, arg:= range os.Args[1:]{
		s += sep + arg
		sep =" "
	}
	fmt.Println(s) 
}

//declaration of string variable

// s := ""
// var s string
// var s string = ""
// var s = ""