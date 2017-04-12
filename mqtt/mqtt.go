package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/bbqgophers/messages"
	"github.com/bbqgophers/qpid"
	"github.com/gomqtt/client"
	"github.com/pkg/errors"
	"log"
)

// Sink is a qpid.MetricSink that
// Reports metrics to MQTT
type Sink struct {
	config  *client.Config
	service *client.Service
	topic   string
}

//NewSink returns a new PrometheusSink
func NewSink(topic string) *Sink {
	c := client.NewConfigWithClientID("mqtt://try:try@broker.shiftr.io", "gomqtt/service")
	c.CleanSession = false

	s := client.NewService()

	s.OnlineCallback = func(resumed bool) {
		fmt.Println("Online with mqtt")
		fmt.Println("resumed: %v\n", resumed)
	}
	s.OfflineCallback = func() {
		fmt.Println("offline with mqtt!")
	}
	err := client.ClearSession(c)
	if err != nil {
		panic(err)
	}
	s.Start(c)

	return &Sink{
		config:  c,
		service: s,
		topic:   topic,
	}
}

// Listen starts a GrillStatus listener on GrillStatus channel
// reporting messages received to Prometheus.  Must be started
// in a goroutine before starting grill run loop or grill will block
// when it tries to send first status
func (p *Sink) Listen(s chan qpid.GrillStatus) {

	for status := range s {
		for _, s := range status.GrillSensors {
			t, err := s.Temperature()
			if err != nil {
				log.Println(errors.Wrap(err, "get temperature"))
			}

			tm := messages.GrillTemp{
				Temp: t,
			}

			b, err := json.Marshal(tm)
			if err != nil {
				fmt.Println("Err marshaling Grill Temp", err)
			}

			p.service.Publish(p.GrillTopic(), b, 0, false)

			set, err := s.Setpoint()
			if err != nil {
				log.Println(errors.Wrap(err, "get setpoint"))
			}

			gtm := messages.GrillTarget{
				Temp: set,
			}

			b, err = json.Marshal(gtm)
			if err != nil {
				fmt.Println("Err marshaling Grill Setpoint", err)
			}

			p.service.Publish(p.SetTopic(), b, 0, false)

			fsm := messages.FanStatus{
				FanOn: status.FanOn,
			}

			b, err = json.Marshal(fsm)
			if err != nil {
				fmt.Println("Err marshaling Fan Status", err)
			}

			p.service.Publish(p.FanTopic(), b, 0, false)
		}
	}

}

func (p *Sink) FanTopic() string {
	return p.topic + "/fan"
}

func (p *Sink) GrillTopic() string {

	return p.topic + "/grill"
}

func (p *Sink) SetTopic() string {

	return p.topic + "/setpoint"

}
