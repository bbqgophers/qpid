package main

import (
	"github.com/bbqgophers/qpid/gobot"
	"github.com/bbqgophers/qpid/http"

	"github.com/bbqgophers/qpid/mqtt"
	"github.com/bbqgophers/qpid/log"
	    //"github.com/bbqgophers/qpid/prometheus"
	"github.com/bbqgophers/qpid/twillio"
)

func main() {
	gb := gobot.NewController()




	 l := log.New()

	//p := prometheus.NewSink()
	p := mqtt.NewSink("bketelsen")

	t := twillio.New()

	s := http.NewServer(gb, t, l, p)
	err := s.ListenAndServe(":8081")
	if err != nil {
		panic(err)
	}
}
