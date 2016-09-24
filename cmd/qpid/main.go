package main

import "github.com/bbqgophers/qpid/gobot"

func main() {
	gb := gobot.NewController()
	gb.GrillMonitor().Target(100)
	err := gb.Run()
	if err != nil {
		panic(err)
	}
}
