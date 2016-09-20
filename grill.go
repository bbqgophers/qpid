package qpid

import (
	"math"
	"time"
)

// Location of where something is
type Location int

// TODO: Need other location constants, think more here
const (
	_ Location = iota
	Food
	Top
	Bottom
	Outside
	Inside
)

// Temp is a temperature in Celcius
type Temp int

// C returns the temperature in Celcius
func (t *Temp) C() int {
	return t
}

// F returns the temperature in Fahrenheit
func (t *Temp) F() int {
	return math.Floor(t*9/5) + 32
}

// A Grill represents the cooking chamber of a BBQ cooker.
// It can contain one or more Sensors which report the current
// temperature.
type Grill struct {
	GrillSensors  []Sensor
	MeatSensors   []Sensor
	FoodMonitors  []Monitor
	GrillMonitors []Monitor
	Reporter      GrillReporter
}

// A Probe measures a temperature at a unique location.
type Probe struct {
	ID          int // Unique ID
	Location    Location
	Description string
	Temperature Temp // Current Temperature
}

// GrillStatus reports the temperatures of the Grill and Food probes at
// a point in time
type GrillStatus struct {
	Time              time.Time
	GrillTemperatures []Probe
	FoodTemperatures  []Probe
}

// Sensor is the interface to retrieve the current temperature
// of a probe
type Sensor interface {
	TemperatureF() (int, error)
	TemperatureC() (int, error)
	Location() Location
	Description() string
}

// Sourcer provides a string representing the source of an Alert
type Sourcer interface {
	Source() string
}

// An Alert is triggered when a Threshold is reached.
type Alert struct {
	Time    time.Time
	Message string
	Source  Sourcer
}

// A Thresholder watches a probe for high and low values,
// firing an Alert on the channel returned by Alerts()
type Thresholder interface {
	HighThreshold(Temp) error
	LowThreshold(Temp) error
	Alerts() chan<- Alert
}

// A Targeter sets the desired temperature for a device.
type Targeter interface {
	// Target sets the desired temperature, returns current and/or error
	Target(int) (int, error)
	// Setpoint() returns the currently set desired temperature
	Setpoint() (int, error)
}

type Monitor interface {
	Targeter
	Thresholder
}

// A CookController manages the cook.  A new cook is started by
// calling Run()
type CookController interface {
	FoodMonitors() []Monitor
	GrillMonitor() Monitor
	Run() error
}

// GrillReporter outputs metrics from a Grill
type GrillReporter interface {
	Status() (GrillStatus, error)
}
