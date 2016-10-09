package prometheus

import (
	"log"
	"net/http"

	"github.com/bbqgophers/qpid"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

// Sink is a qpid.MetricSink that
// Reports metrics to Prometheus
type Sink struct {
}

//NewSink returns a new PrometheusSink
func NewSink() *Sink {
	return &Sink{}
}

// Listen starts a GrillStatus listener on GrillStatus channel
// reporting messages received to Prometheus.  Must be started
// in a goroutine before starting grill run loop or grill will block
// when it tries to send first status
func (p *Sink) Listen(s chan qpid.GrillStatus) {

	tempGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "qpid",
		Subsystem: "grill",
		Name:      "temp_f",
		Help:      "Grill temperature readings.",
	}, []string{
		"sensor",
	},
	)
	setGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "qpid",
		Subsystem: "grill",
		Name:      "setpoint_f",
		Help:      "Grill setpoint.",
	}, []string{
		"sensor",
	},
	)
	prometheus.MustRegister(setGauge)

	fanGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "qpid",
		Subsystem: "grill",
		Name:      "fan_bool",
		Help:      "Blower ON = 1 OFF = 0.",
	}, []string{
		"sensor",
	},
	)
	prometheus.MustRegister(fanGauge)
	http.Handle("/metrics", prometheus.Handler())
	go http.ListenAndServe(":8080", nil)

	for message := range s {
		for _, s := range message.GrillSensors {
			t, err := s.Temperature()
			if err != nil {
				log.Println(errors.Wrap(err, "get temperature"))
			}
			log.Println("Temp: c ", s.Temperature())
			tempGauge.WithLabelValues(s.Description()).Set(float64(t.F()))

			set, err := s.Setpoint()
			if err != nil {
				log.Println(errors.Wrap(err, "get setpoint"))
			}
			setGauge.WithLabelValues(s.Description()).Set(float64(set.F()))

			fan := message.FanOn
			v := 0.0
			if fan {
				v = 1.0
			}
			fanGauge.WithLabelValues(s.Description()).Set(float64(v))

		}

	}

}
