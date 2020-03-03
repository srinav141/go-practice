package main

import (
	"flag"
	"fmt"

	"github.com/srinav141/go-practice/src/tconv/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"

	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
	f := tempconv.CToF(tempconv.FreezingC)
	fmt.Println(f)        //32°F
	fmt.Printf("%g\n", f) //32

	flag.Parse()
	fmt.Println(*temp)
}
