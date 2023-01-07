package model

type GreetingEvent struct {
	EventKey *GreetingEventKey
}

func NewGreetingEvent(key *GreetingEventKey) *GreetingEvent {
	return &GreetingEvent{
		EventKey: key,
	}
}

func (e *GreetingEvent) Key() *GreetingEventKey {
	return e.EventKey
}
