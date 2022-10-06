package main

import "fmt"

func main() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    //set {1, 5}
	fmt.Printf("%08b\n", y)    // set {1, 2}
	fmt.Printf("%08b\n", x^y)  // symmteric diff
	fmt.Printf("%08b\n", x&y)  //intersection
	fmt.Printf("%08b\n", x|y)  //union
	fmt.Printf("%08b\n", x&^y) //the difference

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1) //set {2,6}
	fmt.Printf("%08b\n", x>>1) // set {0,4}
}
