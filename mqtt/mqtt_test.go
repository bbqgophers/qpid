package mqtt

import (
	"context"
	"testing"
	"time"

	"github.com/bbqgophers/qpid"
)

func TestSending(t *testing.T) {
	s := NewSink("bketelsen")

	ch := make(chan qpid.GrillStatus)
	go s.Listen(context.Background(), ch)
	for n := 1; n < 1000; n++ {
		var fst bool
		if n%2 == 0 {
			fst = true
		}
		stat := qpid.GrillStatus{
			FanOn: fst,
		}
		ch <- stat

	}
	time.Sleep(5 * time.Second)
	close(ch)
}
