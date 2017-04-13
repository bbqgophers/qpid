package messages

import (
	"math"
	"time"
)

type GrillTemp struct {
	Temp     Temp
	Time     time.Time
	Location Location
}

type FoodTemp struct {
	Temp     Temp
	Time     time.Time
	Location Location
}

type GrillTarget struct {
	Temp Temp
	Time time.Time
}

type FoodTarget struct {
	Temp Temp
	Time time.Time
}

type FanStatus struct {
	FanOn bool
}

// Temp is a temperature in Celcius
type Temp int

// C returns the temperature in Celcius
func (t *Temp) C() int {
	return int(*t)
}

// F returns the temperature in Fahrenheit
func (t *Temp) F() int {
	temp := math.Floor(float64(*t)*9/5) + 32
	if temp < 0 {
		return int(temp - 1.0)
	}
	return int(temp)
}

// TempFromF returns Temp from a fahrenheit value
func TempFromF(f int) Temp {
	c := math.Floor(float64((f - 32)) / 1.8)
	return Temp(c)

}

// Location of where something is
type Location int

// TODO: Need other location constants, think more here
const (
	Food1 Location = iota
	Food2
	Food3
	Food4
	Food5
	Food6
	Top
	Bottom
	Outside
	Inside
	Right
	Left
)

// LocationMap returns string values for
// Locations
var LocationMap = map[Location]string{
	Food1:   "FOOD1",
	Food2:   "FOOD2",
	Food3:   "FOOD3",
	Food4:   "FOOD4",
	Food5:   "FOOD5",
	Food6:   "FOOD6",
	Top:     "TOP",
	Bottom:  "BOTTOM",
	Outside: "OUTSIDE",
	Inside:  "INSIDE",
	Right:   "RIGHT",
	Left:    "LEFT,",
}

func (l Location) String() string {
	return LocationMap[l]
}
