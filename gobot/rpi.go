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
	Sleep = 10
	P     = 3.0
	I     = .05
	B     = 0.0
)

type GobotController struct {
	grillProbe *GobotProbe
	gobot      *gb.Gobot
	pi         *raspi.RaspiAdaptor
	api        *api.API
	pid        *pidctrl.PIDController
	heating    bool
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
		grillProbe: NewProbe(r),
		gobot:      g,
		pi:         r,
		pid:        pid,
		api:        a,
	}
}

func (g *GobotController) FoodMonitors() []qpid.Monitor {
	panic("not implemented")
}

func (g *GobotController) GrillMonitor() qpid.Monitor {
	return g.grillProbe
}

func (g *GobotController) Run() error {

	target, err := g.grillProbe.Setpoint()
	if err != nil {
		return err
	}
	g.pid.Set(float64(target))

	for x := 1; x < 1000; x++ {

		time.Sleep(1 * time.Second)
		temp, err := g.grillProbe.Temperature()
		if err != nil {
			return err
		}
		output := g.pid.Update(float64(temp.C()))
		fmt.Printf("%d - %d C - Output: %f\n", x, temp, output)

		for x := 1; x < 10; x++ {

			if output > float64(x^2) {
				if !g.heating {
					g.heating = true
					fmt.Println("turning on the blower")
					err := g.pi.DigitalWrite("15",0x1)
					if err != nil {
						fmt.Println("error turning on blower")
					}
				}
				fmt.Println("leaving blower on")
			} else {
				if g.heating {
					g.heating = false
					fmt.Println("turning blower off")
					err := g.pi.DigitalWrite("15",0x0)
					if err != nil {
						fmt.Println("error turning off blower")
					}
				}
			}
		}
		if output < float64(10) {
			fmt.Println("temperature reached, sleep 10")
			time.Sleep(10 * time.Second)
		}
	}
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
