package tempconv

import "fmt"

//Celsius celsius temp
type Celsius float64

//Fahrenheit fahrenheit temp
type Fahrenheit float64

const (
	//AbsoluteZeroC abs zero temp
	AbsoluteZeroC Celsius = -273.15
	//FreezingC freezing temp
	FreezingC Celsius = 0
	//BoilingC boiling temp
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
