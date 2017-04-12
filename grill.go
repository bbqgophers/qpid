package qpid

import (
	"fmt"
	"time"
	"github.com/bbqgophers/messages"
)


// MessageType is a constant for different levels
// of messages
type MessageType int

const (
	// Info is an informational message type
	Info MessageType = iota
	// Warning is an warning message type
	Warning
	// Critical is a critical message type
	Critical
	// ThresholdAlert is a message type to indicate a threshold has been reached
	ThresholdAlert
)

// MessageMap returns string values for
// MessageTypes
var MessageMap = map[MessageType]string{
	Info:           "INFO",
	Warning:        "WARNING",
	Critical:       "CRITICAL",
	ThresholdAlert: "THRESHOLD ALERT",
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

// GrillStatus reports the temperatures of the Grill and Food probes at
// a point in time
type GrillStatus struct {
	Time         time.Time
	GrillSensors []Sensor
	FoodSensors  []Sensor
	FanOn        bool
}

// Sensor is the interface to retrieve the current temperature
// of a probe
type Sensor interface {
	Temperature() (messages.Temp, error)
	Location() messages.Location
	Description() string
	Setpointer
}

// Sourcer provides a string representing the source of an Alert
type Sourcer interface {
	Source() string
}

// An Notification is a message that can be sent from various devices
type Notification struct {
	Time        time.Time
	Message     string
	MessageType MessageType
	Source      Sourcer
}

// GoString implements the GoStringer interface to allow
// notifications to be printed as strings
func (n Notification) GoString() string {
	return fmt.Sprintf("%s: [%s] %s -  %s", n.Time.String(), MessageMap[n.MessageType], n.Source.Source(), n.Message)
}

// A Thresholder watches a probe for high and low values,
// firing an Alert on the channel returned by Alerts()
type Thresholder interface {
	HighThreshold(messages.Temp) error
	LowThreshold(messages.Temp) error
	Alerts() chan Notification
}

// A Targeter sets the desired temperature for a device.
type Targeter interface {
	// Target sets the desired temperature, returns current and/or error
	Target(messages.Temp) (messages.Temp, error)
}

// A Setpointer can return the current target temperature
type Setpointer interface {
	// Setpoint() returns the currently set desired temperature
	Setpoint() (messages.Temp, error)
}

// A Monitor is a device that implements both the
// Thresholder and Targeter interfaces
type Monitor interface {
	Targeter
	Thresholder
}

// A CookController manages the cook.  A new cook is started by
// calling Run()
type CookController interface {
	GrillReporter
	FoodMonitors() []Monitor
	GrillMonitor() Monitor
	Run() error
	Stop() error
}

// GrillReporter outputs metrics from a Grill
type GrillReporter interface {
	Status() (GrillStatus, error)
	Notifications() chan Notification
	Metrics() chan GrillStatus
}

// NotificationSink is an interface that must be implemented
// to receive Notifications from a GrillReporter.  There must
// be one NotificationSink registered or Notification reporting will
// block TODO: Figure out a way around blocking or create a nullsink
type NotificationSink interface {
	Listen(chan Notification)
}

// AlertSink is an interface that must be implemented
// to receive Alerts from a Thresholder.  There must
// be one AlertSink registered or Alert reporting will
// block TODO: Figure out a way around blocking or create a nullsink
type AlertSink interface {
	Listen(chan Notification)
}

// MetricSink is an interface that must be implemented
// to receive metrics from a GrillReporter.  There must
// be one MetricSink registered or metric reporting will
// block TODO: Figure out a way around blocking or create a nullsink
type MetricSink interface {
	Listen(chan GrillStatus)
}
