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
// reporting messages received to Prometheus
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
	prometheus.MustRegister(tempGauge)

	http.Handle("/metrics", prometheus.Handler())
	go http.ListenAndServe(":8080", nil)

	for message := range s {
		for _, s := range message.GrillSensors {
			t, err := s.Temperature()
			if err != nil {
				log.Println(errors.Wrap(err, "get temperature"))
			}
			tempGauge.WithLabelValues(s.Description()).Set(float64(t.F()))
			// take this out later
			log.Printf("Metrics: %#v", message)
		}
	}
}
