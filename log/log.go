package log

import (
	"log"

	"github.com/bbqgophers/qpid"
)

// Log implements the qpid.NotificationSink interface
// logging messages to standard out (for now)
type Log struct {
}

func New() *Log {
	return &Log{}
}

func (l *Log) Listen(n chan qpid.Notification) {
	for message := range n {
		log.Printf("LOG: %#v", message)
	}
}
