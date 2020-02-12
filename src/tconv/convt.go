package main

import (
	"fmt"

	"github.com/srinav141/go-practice/src/tconv/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"

	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
	f := tempconv.CToF(tempconv.FreezingC)
	fmt.Println(f)        //32°F
	fmt.Printf("%g\n", f) //32
}
