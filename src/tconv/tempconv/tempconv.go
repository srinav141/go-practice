package tempconv

import "fmt"

//Celsius celsius temp
type Celsius float64

//Fahrenheit fahrenheit temp
type Fahrenheit float64

type celsiusFlag struct {
	Celsius
}

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

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {

	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("Invalid temperature %q", s)
}
