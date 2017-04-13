package gobot

import (
	"fmt"
	"time"

	"github.com/bbqgophers/qpid"
	"github.com/felixge/pidctrl"
	gb "github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/pkg/errors"
)

// i2cAddress is the location that we read
// to get the temperatures from the i2c bus
const i2cAddress = 0x4d

var (
	// Sleep is the duration in seconds we hold when
	// temperature is met
	Sleep = 10
	// P is the P in PID :)
	P = 3.0
	// I is the I in PID
	I = .05
	// B is a booster.  Currently unused
	B = 0.0
	// MetricsIntervalSeconds is the metric reporting interval
	MetricsIntervalSeconds = 5
)

type Controllerer interface {
	FoodMonitors() []qpid.Monitor
	GrillMonitor() qpid.Monitor
	Run() error
	Stop() error
	Status() (qpid.GrillStatus, error)
	Notifications() chan qpid.Notification
	Metrics() chan qpid.GrillStatus
	Source() string
}

// Controller represents all the electronics that
// make the PID work
type Controller struct {
	grillProbe    *Probe
	gobot         *gb.Gobot
	pi            *raspi.RaspiAdaptor
	api           *api.API
	pid           *pidctrl.PIDController
	heating       bool
	notifications chan qpid.Notification
	metrics       chan qpid.GrillStatus
}

// NewController returns a new Controller
// initialized.
func NewController() *Controller {
	n := make(chan qpid.Notification)
	metrics := make(chan qpid.GrillStatus)
	g := gb.NewGobot()
	r := raspi.NewRaspiAdaptor("qpid")
	robot := gb.NewRobot("qpid",
		[]gb.Connection{r},
		[]gb.Device{},
		nil,
	)
	errs := r.Connect()
	if errs != nil {
		panic(errs)
	}
	g.AddRobot(robot)

	a := api.NewAPI(g)
	a.Port = "4000"
	a.AddHandler(api.BasicAuth("bbq", "gopher"))
	a.Start()
	e := r.I2cStart(i2cAddress)
	if e != nil {
		panic(e)
	}
	pid := pidctrl.NewPIDController(P, I, B)
	pid.SetOutputLimits(-100.0, 100.0)
	return &Controller{
		grillProbe:    NewProbe(r),
		gobot:         g,
		pi:            r,
		pid:           pid,
		api:           a,
		notifications: n,
		metrics:       metrics,
	}
}

// FoodMonitors returns the food monitors
func (g *Controller) FoodMonitors() []qpid.Monitor {
	panic("not implemented")
}

// GrillMonitor returns the grill probe
func (g *Controller) GrillMonitor() qpid.Monitor {
	return g.grillProbe
}

// Run starts the grill's run loop and blocks
func (g *Controller) Run() error {

	go func() {
		errs := g.gobot.Start()
		if errs != nil {
			// hack - maybe change interface?
			panic(errs)
		}
	}()

	// TODO: Decide where the blocking happens.
	// is this routine where we block and run until
	// receiving some signal to exit?  Or do we block in main?
	target, err := g.grillProbe.Setpoint()
	if err != nil {
		return err
	}
	g.pid.Set(float64(target))

	// start goroutine to send status on interval
	go func() {
		for {
			status, err := g.Status()
			if err != nil {
				g.notify(qpid.Critical, errors.Wrap(err, "grill status").Error())
			}
			g.metrics <- status
			time.Sleep(time.Duration(MetricsIntervalSeconds) * time.Second)
		}
	}()

	for { // block here for now
		time.Sleep(1 * time.Second)
		temp, err := g.grillProbe.Temperature()
		if err != nil {
			return err
		}
		output := g.pid.Update(float64(temp.C()))
		message := fmt.Sprintf("Temp: %d, Integral: %f", temp, output)
		g.notify(qpid.Info, message)

		for x := 1; x < 10; x++ {

			if output > float64(x^2) {
				if !g.heating {
					g.heating = true

					g.notify(qpid.Info, "Turning on Blower")
					err := g.pi.DigitalWrite("15", 0x1)
					if err != nil {
						g.notify(qpid.Critical, errors.Wrap(err, "turning on blower").Error())
					}
				}

				g.notify(qpid.Info, "Leaving Blower On")
			} else {
				if g.heating {
					g.heating = false
					g.notify(qpid.Info, "Turning off Blower")
					err := g.pi.DigitalWrite("15", 0x0)
					if err != nil {
						g.notify(qpid.Critical, errors.Wrap(err, "turning off blower").Error())
					}
				}
			}
		}
		if output < float64(10) {

			g.notify(qpid.Info, "Temp Good, Sleep 10")
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}

// Stop stops the grill's run loop
// but actually probably wont.  because
// there's nothing to do that yet.
func (g *Controller) Stop() error {

	g.notify(qpid.Info, "Received Stop Request")
	errs := g.gobot.Stop()
	if errs != nil {
		// hack - maybe change interface?
		return errs[0]
	}

	return nil
}

// Status returns the current Grill status.
func (g *Controller) Status() (qpid.GrillStatus, error) {
	return qpid.GrillStatus{
		Time:         time.Now(),
		GrillSensors: []qpid.Sensor{g.grillProbe},
		FanOn:        g.heating,
	}, nil
}

// Notifications returns the Grill's notification channel
func (g *Controller) Notifications() chan qpid.Notification {
	return g.notifications
}

// notify sends a message on the grill's notification channel
func (g *Controller) notify(mt qpid.MessageType, message string) {
	n := qpid.Notification{
		Time:        time.Now(),
		Message:     message,
		MessageType: mt,
		Source:      g,
	}
	g.notifications <- n

}

// Metrics returns the grill's metric channel
func (g *Controller) Metrics() chan qpid.GrillStatus {
	return g.metrics
}

// Source implements Sourcer for pretty printing of events
// and notifications
func (g *Controller) Source() string {
	return fmt.Sprintf("Pi Grill Controller")
}
