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

const i2cAddress = 0x4d

var (
	Sleep                  = 10
	P                      = 3.0
	I                      = .05
	B                      = 0.0
	MetricsIntervalSeconds = 1
)

type GobotController struct {
	grillProbe    *GobotProbe
	gobot         *gb.Gobot
	pi            *raspi.RaspiAdaptor
	api           *api.API
	pid           *pidctrl.PIDController
	heating       bool
	notifications chan qpid.Notification
	metrics       chan qpid.GrillStatus
}

func NewController() *GobotController {
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
	go func() {
		errs := g.Start()
		if errs != nil {
			// hack - maybe change interface?
			panic(errs)
		}
	}()

	pid := pidctrl.NewPIDController(P, I, B)
	return &GobotController{
		grillProbe:    NewProbe(r),
		gobot:         g,
		pi:            r,
		pid:           pid,
		api:           a,
		notifications: n,
		metrics:       metrics,
	}
}

func (g *GobotController) FoodMonitors() []qpid.Monitor {
	panic("not implemented")
}

func (g *GobotController) GrillMonitor() qpid.Monitor {
	return g.grillProbe
}

func (g *GobotController) Run() error {
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

func (g *GobotController) Stop() error {

	g.notify(qpid.Info, "Received Stop Request")
	errs := g.gobot.Stop()
	if errs != nil {
		// hack - maybe change interface?
		return errs[0]
	}

	return nil
}
func (g *GobotController) Status() (qpid.GrillStatus, error) {
	return qpid.GrillStatus{
		Time:         time.Now(),
		GrillSensors: []qpid.Sensor{g.grillProbe},
	}, nil
}

func (g *GobotController) Notifications() chan qpid.Notification {
	return g.notifications
}
func (g *GobotController) notify(mt qpid.MessageType, message string) {
	n := qpid.Notification{
		Time:        time.Now(),
		Message:     message,
		MessageType: mt,
		Source:      g,
	}
	g.notifications <- n

}
func (g *GobotController) Metrics() chan qpid.GrillStatus {
	return g.metrics
}

func (g *GobotController) Source() string {
	return fmt.Sprintf("Pi Grill Controller")
}
