package twillio

import "github.com/bbqgophers/qpid"

type TwillioClient struct {
}

func (t *TwillioClient) Listen(chan<- qpid.Notification) {
	panic("not implemented")
}
