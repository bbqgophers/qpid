package mqtt

import (
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type MqttAdaptor struct {
	name     string
	Host     string
	clientID string
	username string
	password string
	client   *mqtt.Client
}

// NewMqttAdaptor creates a new mqtt adaptor with specified name, host and client id
func NewMqttAdaptor(name string, host string, clientID string) *MqttAdaptor {
	return &MqttAdaptor{
		name:     name,
		Host:     host,
		clientID: clientID,
	}
}

func NewMqttAdaptorWithAuth(name, host, clientID, username, password string) *MqttAdaptor {
	return &MqttAdaptor{
		name:     name,
		Host:     host,
		clientID: clientID,
		username: username,
		password: password,
	}
}

func (a *MqttAdaptor) Name() string { return a.name }

// Connect returns true if connection to mqtt is established
func (a *MqttAdaptor) Connect() (errs []error) {
	a.client = mqtt.NewClient(createClientOptions(a.clientID, a.Host, a.username, a.password))
	if token := a.client.Connect(); token.Wait() && token.Error() != nil {
		errs = append(errs, token.Error())
	}

	return
}

// Disconnect returns true if connection to mqtt is closed
func (a *MqttAdaptor) Disconnect() (err error) {
	if a.client != nil {
		a.client.Disconnect(500)
	}
	return
}

// Finalize returns true if connection to mqtt is finalized successfully
func (a *MqttAdaptor) Finalize() (errs []error) {
	a.Disconnect()
	return
}

// Publish a message under a specific topic
func (a *MqttAdaptor) Publish(topic string, message []byte) bool {
	if a.client == nil {
		return false
	}
	a.client.Publish(topic, 0, false, message)
	return true
}

// Subscribe to a topic, and then call the message handler function when data is received
func (a *MqttAdaptor) On(event string, f func(s []byte)) bool {
	if a.client == nil {
		return false
	}
	a.client.Subscribe(event, 0, func(client *mqtt.Client, msg mqtt.Message) {
		f(msg.Payload())
	})
	return true
}

func createClientOptions(clientId, raw, username, password string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(raw)
	opts.SetClientID(clientId)
	if username != "" && password != "" {
		opts.SetPassword(password)
		opts.SetUsername(username)
	}
	opts.AutoReconnect = false
	return opts
}
