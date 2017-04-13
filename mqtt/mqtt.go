package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bbqgophers/messages"
	"github.com/bbqgophers/qpid"
	"github.com/gomqtt/client"
	"github.com/pkg/errors"
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
	config := client.NewConfigWithClientID("mqtt://bketelsen:ncc1701c@mqtt.bbq.live", "qpid")
	config.CleanSession = false
	s := client.NewService()

	s.OnlineCallback = func(resumed bool) {
		fmt.Println("Online with mqtt")
		fmt.Printf("resumed: %v\n", resumed)
	}
	s.OfflineCallback = func() {
		fmt.Println("offline with mqtt!")
	}
	err := client.ClearSession(config)
	if err != nil {
		panic(err)
	}
	s.Start(config)

	return &Sink{
		config:  config,
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
		var fst int
		if status.FanOn {
			fst = 1
		}
		fsm := messages.FanStatus{
			FanOn: fst,
		}

		b, err := json.Marshal(fsm)
		if err != nil {
			fmt.Println("Err marshaling Fan Status", err)
		}
		fut := p.service.Publish(p.FanTopic(), b, 0, false)

		err = fut.Wait(5 * time.Second)
		if err != nil {
			panic(err)
		}
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
