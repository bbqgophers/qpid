package main

import (
	"fmt"
	"time"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {

	g := gobot.NewGobot()
	r := raspi.NewRaspiAdaptor("raspi")
	errs := r.Connect()
	if errs != nil {
		panic(errs)
	}

	e := r.I2cStart(0x4d)
	if e != nil {
		panic(e)
	}
	work := func() {
		gobot.Every(1*time.Second, func() {
			// get temperature reading
			b, e := r.I2cRead(0x4d, 3)
			if e != nil {
				panic(e)
			}
			// 2nd byte is temp C * 5
			// 3rd byte is temp over 127
			c := b[1] / 5
			c = c + b[2]
			fmt.Println(c)
		})
	}
	robot := gobot.NewRobot("raspi",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)
	g.AddRobot(robot)
	go func(){

  	server := api.NewAPI(g)
  	server.Port = "4000"
  	server.AddHandler(api.BasicAuth("bbq", "gopher"))
  	server.Start()
}()
	g.Start()

}
