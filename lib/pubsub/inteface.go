package pubsub

import "time"

type Message struct {
	Time    time.Time
	Message []byte
}
type Interface interface {
	Register() error
	Subscribe(topic string, sink func(message Message), params ...interface{})
	Publish(topic string, message interface{})

	// SetMarshaller set interface{} to []byte marshalling function
	SetMarshaller(func(input interface{}) ([]byte, error))

	// SetUnMarshaller set []byte to interface{} unmarshalling function
	SetUnMarshaller(func(bytes []byte, out interface{}) error)
}
