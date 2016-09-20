package qpid

import "time"

// A Grill represents the cooking chamber of a BBQ cooker.
// It can contain one or more Sensors which report the current
// temperature.
type Grill struct {
	GrillSensors []Sensor
	MeatSensors  []Sensor
	Controllers  []GrillController
	Reporter     GrillReporter
}

// A GrillThermometer reports the current temperature of a location
// inside a Grill
type GrillThermometer struct {
	Location    string
	Temperature int
}

type FoodThermometer struct {
	Location    string
	Temperature int
}

type GrillStatus struct {
	Time              time.Time
	GrillTemperatures []GrillThermometer
	FoodTemperatures  []FoodThermometer
}

// Sensor is the interface to retrieve the current temperature
// of a probe
type Sensor interface {
	Temperature() (int, error)
}

// GrillController is the interface to set the target cooking temperature
// of a grill
type GrillController interface {
	Target(int) (int, error)
	Run() error
}

// FoodController is the interface to set the target temperature
// of a food item
type FoodController interface {
	Target(int) (int, error)
}

// GrillReporter outputs metrics from a Grill
type GrillReporter interface {
	Status() (GrillStatus, error)
}
