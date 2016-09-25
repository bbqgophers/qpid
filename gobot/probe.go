package gobot

import (
	"fmt"
	"log"

	"github.com/bbqgophers/qpid"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/pkg/errors"
)

// GobotProbe is a thermocoupler connected to a raspberry pi
type GobotProbe struct {
	id          int
	location    qpid.Location
	description string
	setpoint    qpid.Temp
	high        qpid.Temp
	low         qpid.Temp
	temperature qpid.Temp
	pi          *raspi.RaspiAdaptor
	alerts      chan qpid.Notification
}

// NewProbe returns an initialized GobotProbe.
// Location and description hard-coded for now, since
// we only support one thermocoupler.
func NewProbe(pi *raspi.RaspiAdaptor) *GobotProbe {

	a := make(chan qpid.Notification)
	return &GobotProbe{
		alerts:      a,
		pi:          pi,
		id:          1,
		location:    qpid.Inside,
		description: "Grill Internal Probe 1",
	}
}

// Target is the temperature we'd like to reach
func (g *GobotProbe) Target(temp qpid.Temp) (qpid.Temp, error) {
	g.setpoint = temp
	// todo get temp and return that instead
	// if possible
	return g.Temperature()
}

// Setpoint is the current Target
func (g *GobotProbe) Setpoint() (qpid.Temp, error) {
	return g.setpoint, nil
}

// HighThreshold is the temperature max before a critical
// alert should be sent
func (g *GobotProbe) HighThreshold(temp qpid.Temp) error {
	g.high = temp
	return nil
}

// LowThreshold is the temperature min before a critical alert
// should be sent
func (g *GobotProbe) LowThreshold(temp qpid.Temp) error {
	g.low = temp
	return nil
}

// Alerts returns a channel of notifications for probe
// alerts
func (g *GobotProbe) Alerts() chan qpid.Notification {
	return g.alerts
}

// Temperature reads and returns the current temperature
// from the probe
func (g *GobotProbe) Temperature() (qpid.Temp, error) {
	var t qpid.Temp
	b, e := g.pi.I2cRead(i2cAddress, 3)
	if e != nil {
		return t, e
	}
	// 2nd byte is temp C * 5
	// 3rd byte is temp over 127
	c := b[1] << 8
	c = c + b[2]
	c = c / 5
	g.temperature = qpid.Temp(c)
	return g.temperature, e
}

// Location returns the probe's location
func (g *GobotProbe) Location() qpid.Location {
	return g.location
}

// Description returns the probe's description
func (g *GobotProbe) Description() string {
	return g.description
}

// Source implements qpid.Sourcer and returns
// the source of a notification
func (g *GobotProbe) Source() string {
	return fmt.Sprintf("Probe %d: %s", g.id, g.description)
}

func (g *GobotProbe) String() string {
	t, err := g.Temperature()
	if err != nil {
		log.Println(errors.Wrap(err, "sensor get temperature"))
	}
	return fmt.Sprintf("Temp %d F at %s for %s", t.F(), qpid.LocationMap[g.Location()], g.Description)
}
func (g *GobotProbe) GoString() string {
	t, err := g.Temperature()
	if err != nil {
		log.Println(errors.Wrap(err, "sensor get temperature"))
	}
	return fmt.Sprintf("Temp %d F at %s for %s", t.F(), qpid.LocationMap[g.Location()], g.Description)
}
