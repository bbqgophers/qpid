package log

import "github.com/bbqgophers/qpid"

type Log struct {
}

func (l *Log) Listen(chan<- qpid.Notification) {
	panic("not implemented")
}
