package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/chip"
	"github.com/hybridgroup/gobot/platforms/i2c"
)

func main() {
	gbot := gobot.NewGobot()

	board := chip.NewChipAdaptor("chip")
	accel := i2c.NewGroveAccelerometerDriver(board, "accel")

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			if x, y, z, err := accel.XYZ(); err == nil {
				fmt.Println(x, y, z)
				fmt.Println(accel.Acceleration(x, y, z))
			} else {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("accelBot",
		[]gobot.Connection{board},
		[]gobot.Device{accel},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
