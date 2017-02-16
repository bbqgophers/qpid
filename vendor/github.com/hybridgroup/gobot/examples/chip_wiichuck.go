package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/chip"
	"github.com/hybridgroup/gobot/platforms/i2c"
)

func main() {
	gbot := gobot.NewGobot()

	chipAdaptor := chip.NewChipAdaptor("chip")
	wiichuck := i2c.NewWiichuckDriver(chipAdaptor, "wiichuck")

	work := func() {
		wiichuck.On(wiichuck.Event("joystick"), func(data interface{}) {
			fmt.Println("joystick", data)
		})

		wiichuck.On(wiichuck.Event("c"), func(data interface{}) {
			fmt.Println("c")
		})

		wiichuck.On(wiichuck.Event("z"), func(data interface{}) {
			fmt.Println("z")
		})
		wiichuck.On(wiichuck.Event("error"), func(data interface{}) {
			fmt.Println("Wiichuck error:", data)
		})
	}

	robot := gobot.NewRobot("chuck",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{wiichuck},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
