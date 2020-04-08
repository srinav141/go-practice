package space

import("math"
"errors"
)

type Planet string

func Age(seconds float64, planet Planet) float64{

const earthYear = 31557600.0
yearValue := 0.0
switch planet {
case "Earth":
  yearValue = 1 * earthYear
case "Mercury":
  yearValue = 0.2408467 * earthYear
case "Venus":
  yearValue = 0.61519726 * earthYear
case "Mars":
  yearValue = 1.8808158 * earthYear
case "Jupiter":
  yearValue =11.862615 * earthYear
case "Saturn":
  yearValue = 29.447498 * earthYear
case "Uranus":
  yearValue = 84.016846 * earthYear
case "Neptune":
  yearValue = 164.79132 * earthYear
default:
  errors.New("Planet not defined")
}
x := seconds/yearValue
return math.Round(x*100)/100

}
