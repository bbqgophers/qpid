package gobot

import (
	"time"

	"github.com/bbqgophers/qpid"
	gb "github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

const i2cAddress = 0x4d

type GobotController struct {
	grillProbe *GobotProbe
	gobot      *gb.Gobot
	pi         *raspi.RaspiAdaptor
	api        *api.API
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
	return &GobotController{
		grillProbe: NewProbe(r),
		gobot:      g,
		pi:         r,
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
	go func(){
		errs := g.gobot.Start()
		if errs != nil {
			// hack - maybe change interface?
			panic(errs)
		}
	}()
	return nil
}

func (g *GobotController) Stop() error {
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
