package http

import (
	"html/template"
	h "net/http"
	"strconv"

	"github.com/bbqgophers/qpid"
	"github.com/gorilla/mux"
)

// Server is an HTTP server that controls
// the cook
type Server struct {
	controller qpid.CookController
	alerter    qpid.AlertSink
	notifier   qpid.NotificationSink
	metricer   qpid.MetricSink
	router     *mux.Router
}

// NewServer returns an initialized Server
func NewServer(c qpid.CookController,
	a qpid.AlertSink,
	n qpid.NotificationSink,
	m qpid.MetricSink) *Server {
	r := mux.NewRouter()
	s := &Server{
		controller: c,
		alerter:    a,
		notifier:   n,
		metricer:   m,
		router:     r,
	}
	return s
}

// ListenAndServe starts an HTTP listener at address `addr`,
// blocking until it's stopped
func (s *Server) ListenAndServe(addr string) error {

	s.router.HandleFunc("/", s.Index).Methods("GET")
	s.router.HandleFunc("/run", s.Run).Methods("POST")
	s.router.HandleFunc("/status", s.Status).Methods("GET")
	err := h.ListenAndServe(addr, s.router)
	return err

}

// Index serves the / route
func (s *Server) Index(w h.ResponseWriter, r *h.Request) {
	t, err := template.New("index").Parse(indexT)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	t.Execute(w, nil)
}

// Run starts a cook
func (s *Server) Run(w h.ResponseWriter, r *h.Request) {
	t := r.PostFormValue("temp")
	if t == "" {
		w.WriteHeader(400)
		w.Write([]byte("Temperature not provided"))
		return
	}
	tf, err := strconv.Atoi(t)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Can't parse temperature as a number"))
		return
	}
	temp := qpid.TempFromF(tf)
	s.controller.GrillMonitor().Target(temp) //ignoring return temp and error TODO

	// start that cook!

	s.controller.GrillMonitor().Target(temp)

	// TODO done channel on all listeners and runner
	go s.notifier.Listen(s.controller.Notifications())

	go s.metricer.Listen(s.controller.Metrics())

	go s.alerter.Listen(s.controller.GrillMonitor().Alerts())
	go s.controller.Run()

	h.Redirect(w, r, "/status", h.StatusSeeOther)
	return
}

// Status serves the /status route
func (s *Server) Status(w h.ResponseWriter, r *h.Request) {
	t, err := template.New("status").Parse(statusT)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	status, err := s.controller.Status()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	t.Execute(w, status)
}
