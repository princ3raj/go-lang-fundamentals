package main

import (
	"fmt"

	"employee/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
}
func CToF(c tempconv.Celsius) tempconv.Fahreneit {
	return tempconv.Fahreneit((c * 9 / 5) + 32)
}

func FToC(f tempconv.Fahreneit) tempconv.Celsius {
	return tempconv.Celsius((f - 32) * 5 / 9)
}
