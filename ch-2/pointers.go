package main

import "fmt"

func main() {
	var x = 2
	fmt.Println(f() == f())
	fmt.Println(incr(&x))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
