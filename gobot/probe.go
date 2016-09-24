package gobot

import (
	"fmt"

	"github.com/bbqgophers/qpid"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

type GobotProbe struct {
	id          int
	location    qpid.Location
	description string
	setpoint    qpid.Temp
	high        qpid.Temp
	low         qpid.Temp
	temperature qpid.Temp
	pi          *raspi.RaspiAdaptor
	alerts      chan<- qpid.Notification
}

func NewProbe(pi *raspi.RaspiAdaptor) *GobotProbe {

	a := make(chan<- qpid.Notification)
	return &GobotProbe{
		alerts:      a,
		pi:          pi,
		id:          1,
		location:    qpid.Inside,
		description: "Grill Internal Probe 1",
	}
}

func (g *GobotProbe) Target(temp qpid.Temp) (qpid.Temp, error) {
	g.setpoint = temp
	// todo get temp and return that instead
	// if possible
	return g.Temperature()
}

func (g *GobotProbe) Setpoint() (qpid.Temp, error) {
	return g.setpoint, nil
}

func (g *GobotProbe) HighThreshold(temp qpid.Temp) error {
	g.high = temp
	return nil
}

func (g *GobotProbe) LowThreshold(temp qpid.Temp) error {
	g.low = temp
	return nil
}

func (g *GobotProbe) Alerts() chan<- qpid.Notification {
	return g.alerts
}
func (g *GobotProbe) Temperature() (qpid.Temp, error) {
	var t qpid.Temp
	b, e := g.pi.I2cRead(i2cAddress, 3)
	if e != nil {
		return t, e
	}
	// 2nd byte is temp C * 5
	// 3rd byte is temp over 127
	c := b[1] / 5
	c = c + b[2]
	g.temperature = qpid.Temp(c)
	return g.temperature, e
}

func (g *GobotProbe) Location() qpid.Location {
	return g.location
}

func (g *GobotProbe) Description() string {
	return g.description
}

func (g *GobotProbe) Source() string {
	return fmt.Sprintf("Probe %d: %s", g.id, g.description)
}
