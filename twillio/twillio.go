package twillio

import (
	"log"

	"github.com/bbqgophers/qpid"
)

type TwillioClienter interface {
	Listen(a chan qpid.Notification)
}

// Client is a struct to hold
// twillio config
type Client struct {
}

// New returns an initialized
// twillio client
func New() *Client {
	return &Client{}
}

// Listen starts a listener on the notification channel.  It
// must be started in a goroutine before the grill's run loop
// starts or the grill will block when it sends a critical alert.
func (t *Client) Listen(a chan qpid.Notification) {
	for message := range a {
		log.Printf("ALERT: %#v", message)
	}
}
