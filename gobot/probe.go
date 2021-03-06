package gobot

import (
	"fmt"
	"log"
	"sync"

	"github.com/bbqgophers/messages"
	"github.com/bbqgophers/qpid"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/pkg/errors"
)

// Probe is a thermocoupler connected to a raspberry pi
type Probe struct {
	id          int
	location    messages.Location
	description string
	setpoint    messages.Temp
	high        messages.Temp
	low         messages.Temp
	temperature messages.Temp
	pi          *raspi.RaspiAdaptor
	alerts      chan qpid.Notification
	tempMu      sync.Mutex
}

// NewProbe returns an initialized Probe.
// Location and description hard-coded for now, since
// we only support one thermocoupler.
func NewProbe(pi *raspi.RaspiAdaptor) *Probe {

	a := make(chan qpid.Notification)
	return &Probe{
		alerts:      a,
		pi:          pi,
		id:          1,
		setpoint:    messages.TempFromF(225),
		location:    messages.Inside,
		description: "Grill Internal Probe 1",
	}
}

// Target is the temperature we'd like to reach
func (g *Probe) Target(temp messages.Temp) (messages.Temp, error) {
	g.setpoint = temp
	// todo get temp and return that instead
	// if possible
	return g.Temperature()
}

// Setpoint is the current Target
func (g *Probe) Setpoint() (messages.Temp, error) {
	return g.setpoint, nil
}

// HighThreshold is the temperature max before a critical
// alert should be sent
func (g *Probe) HighThreshold(temp messages.Temp) error {
	g.high = temp
	return nil
}

// LowThreshold is the temperature min before a critical alert
// should be sent
func (g *Probe) LowThreshold(temp messages.Temp) error {
	g.low = temp
	return nil
}

// Alerts returns a channel of notifications for probe
// alerts
func (g *Probe) Alerts() chan qpid.Notification {
	return g.alerts
}

// Temperature reads and returns the current temperature
// from the probe
func (g *Probe) Temperature() (messages.Temp, error) {
	g.tempMu.Lock()
	defer g.tempMu.Unlock()
	var t messages.Temp
	b, e := g.pi.I2cRead(i2cAddress, 2)
	if e != nil {
		return t, e
	}

	var final uint
	fmt.Println("b0,b1:", b[0],b[1])

	final = uint(b[0]) << 8
	final = final + uint(b[1])
	final = final / 5

	fmt.Println("b0,b1,final:", b[0],b[1],final)

	g.temperature = messages.Temp(int(final))
	return g.temperature, e
}

// Location returns the probe's location
func (g *Probe) Location() messages.Location {
	return g.location
}

// Description returns the probe's description
func (g *Probe) Description() string {
	return g.description
}

// Source implements qpid.Sourcer and returns
// the source of a notification
func (g *Probe) Source() string {
	return fmt.Sprintf("Probe %d: %s", g.id, g.description)
}

func (g *Probe) String() string {
	t, err := g.Temperature()
	if err != nil {
		log.Println(errors.Wrap(err, "sensor get temperature"))
	}
	return fmt.Sprintf("Temp %d F at %s for %s", t.F(), messages.LocationMap[g.Location()], g.Description)
}
func (g *Probe) GoString() string {
	t, err := g.Temperature()
	if err != nil {
		log.Println(errors.Wrap(err, "sensor get temperature"))
	}
	return fmt.Sprintf("Temp %d F at %s for %s", t.F(), messages.LocationMap[g.Location()], g.Description)
}
