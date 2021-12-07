package broker

type Producer interface {
	Publish(msg interface{}, key []byte, topic string) error
}
