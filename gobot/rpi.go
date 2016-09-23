package gobot

import (
	"fmt"
	"time"

	"github.com/bbqgophers/qpid"
	"github.com/felixge/pidctrl"
	gb "github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

const i2cAddress = 0x4d

var (
	P = 3.0
	I = .05
	B = 0.0
)

type GobotController struct {
	grillProbe *GobotProbe
	gobot      *gb.Gobot
	pi         *raspi.RaspiAdaptor
	api        *api.API
	pid        *pidctrl.PIDController
}

func NewController() *GobotController {
	g := gb.NewGobot()
	r := raspi.NewRaspiAdaptor("qpid")
	robot := gb.NewRobot("qpid",
		[]gb.Connection{r},
		[]gb.Device{},
		nil,
	)
	errs := r.Connect()
	if errs != nil {
		return nil
	}
	g.AddRobot(robot)

	pid := pidctrl.NewPIDController(P, I, B)
	return &GobotController{
		grillProbe: NewProbe(r),
		gobot:      g,
		pi:         r,
		pid:        pid,
	}
}

func (g *GobotController) FoodMonitors() []qpid.Monitor {
	panic("not implemented")
}

func (g *GobotController) GrillMonitor() qpid.Monitor {
	return g.grillProbe
}

func (g *GobotController) Run() error {

	g.api = api.NewAPI(g.gobot)
	g.api.Port = "4000"
	g.api.AddHandler(api.BasicAuth("bbq", "gopher"))
	g.api.Start()
	e := g.pi.I2cStart(i2cAddress)
	if e != nil {
		return e
	}
	errs := g.gobot.Start()
	if errs != nil {
		// hack - maybe change interface?
		return errs[0]
	}

	g.pid.Set(100.0)

	for x := 1; x < 1000; x++ {

		time.Sleep(1 * time.Second)
		temp, err := g.grillProbe.Temperature()
		if err != nil {
			return err
		}
		output := g.pid.Update(float64(temp.C()))
		fmt.Printf("%d - %d C - Output: %f", x, temp, output)
	}

	return nil
}

func (g *GobotController) Stop() error {
	panic("not implemented")
}
func (g *GobotController) Status() (qpid.GrillStatus, error) {
	return qpid.GrillStatus{
		Time:         time.Now(),
		GrillSensors: []qpid.Sensor{g.grillProbe},
	}, nil
}
