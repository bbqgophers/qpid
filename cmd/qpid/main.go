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
	go l.Listen(gb.Notifications())

	p := prometheus.NewSink()
	go p.Listen(gb.Metrics())

	t := twillio.NewClient()
	go t.Listen(gb.GrillMonitor().Alerts())
	err := gb.Run()
	if err != nil {
		panic(err)
	}
}
