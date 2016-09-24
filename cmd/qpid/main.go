package main

import (
	"github.com/bbqgophers/qpid/gobot"
	"github.com/bbqgophers/qpid/log"
	"github.com/bbqgophers/qpid/prometheus"
	"github.com/bbqgophers/qpid/twillio"
)

func main() {
	gb := gobot.NewController()
	gb.GrillMonitor().Target(100)

	l := log.New()
	l.Listen(gb.Notifications())

	p := prometheus.NewSink()
	p.Listen(gb.Metrics())

	t := twillio.NewClient()
	t.Listen(gb.GrillMonitor().Alerts())
	err := gb.Run()
	if err != nil {
		panic(err)
	}
}
