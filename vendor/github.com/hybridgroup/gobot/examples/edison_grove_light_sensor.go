package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("board")
	sensor := gpio.NewGroveLightSensorDriver(board, "sensor", "0")

	work := func() {
		sensor.On(sensor.Event("data"), func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
