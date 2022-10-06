package main

import "fmt"

type Celsius float64
type Fahreneit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	c := FToC(212.0)
	fmt.Println(c.String()) // 100 C
	fmt.Printf("%v\n",c) 	// 100 C
	fmt.Printf("%s\n", c) 	// 100 C
	fmt.Println(c) 			//100 C
	fmt.Printf("%g\n", c)	// 100; does not call string
	fmt.Println(float64(c)) // 100; does not call string


}

func CToF(c Celsius) Fahreneit {
	return Fahreneit((c * 9 / 5) + 32)
}

func FToC(f Fahreneit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string { return fmt.Sprintf("%gC", c)}
