package main

import "github.com/bbqgophers/qpid/gobot"

func main() {
	gb := gobot.NewController()
	err := gb.Run()
	if err != nil {
		panic(err)
	}
}
