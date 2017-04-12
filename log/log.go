package log

import (
	"log"

	"github.com/bbqgophers/qpid"
)

type Logger interface {
	Listen(n chan qpid.Notification)
}

// Log implements the qpid.NotificationSink interface
// logging messages to standard out (for now)
type Log struct {
}

// New returns a new Log struct
// no config for now TODO
func New() *Log {
	return &Log{}
}

// Listen starts a listener on the notification
// channel.  Must be started in a goroutine before
// starting grill run loop or grill will block sending
// notifications
func (l *Log) Listen(n chan qpid.Notification) {
	for message := range n {
		log.Printf("LOG: %#v", message)
	}
}
