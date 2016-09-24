package twillio

import (
	"log"

	"github.com/bbqgophers/qpid"
)

type TwillioClient struct {
}

func NewClient() *TwillioClient {
	return &TwillioClient{}
}

func (t *TwillioClient) Listen(a chan qpid.Notification) {
	for message := range a {
		log.Printf("ALERT: %#v", message)
	}
}
