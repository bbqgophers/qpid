package main

import (

	"fmt"
        "github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {

        r := raspi.NewRaspiAdaptor("raspi")
	errs := r.Connect()
	if errs != nil {
		panic(errs)
	}

	e := r.I2cStart(0x4d)
	if e != nil {
		panic(e)
	}
	b, e := r.I2cRead(0x4d,3)
	if e != nil {
		panic(e)
	}
	c := b[1] / 5
	fmt.Println(c)
}
