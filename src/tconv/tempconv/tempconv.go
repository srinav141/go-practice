package tempconv

import "fmt"

//Celsius celsius temp
type Celsius float64

//Fahrenheit fahrenheit temp
type Fahrenheit float64

const (
	//AbsolluteAero abs zero temp
	AbsolluteAero Celsius = -273.15
	//Freezingc freezing temp
	Freezingc Celsius = 0
	//Boilingc boiling temp
	Boilingc Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
