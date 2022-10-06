package main

import "fmt"

func main() {

	var x uint = 1 << 1
	fmt.Printf("%010b\n", x)

	medals := []string{"1", "2", "3"}
	for i:= len(medals) - 1; i >= 0 ; i--{
		fmt.Println(medals[i])
	}
}
