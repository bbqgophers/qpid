package prometheus

import "github.com/bbqgophers/qpid"

type PrometheusSink struct {
}

func (p *PrometheusSink) Listen(chan<- qpid.GrillStatus) {
	panic("not implemented")
}
