package mock

import (
	"log"

	"github.com/bbqgophers/qpid"
)

type AlertSink struct {
}

func NewAlertSink() *AlertSink {
	return &AlertSink{}
}
func (a *AlertSink) Listen(alerts chan qpid.Notification) {
	for message := range alerts {
		log.Printf("ALERT: %#v", message)
	}
}
