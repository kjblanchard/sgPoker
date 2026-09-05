package events

type Event struct {
	Type int
	Data any
}

type EventTypes int

const (
	EventTypeTCell = iota
	EventTypeStatusMessage
)

var (
	EventBus chan Event
)

func UpdateEvents() {

}
