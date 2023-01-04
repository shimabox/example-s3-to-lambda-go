package model

type GreetingEventList struct {
	Events []*GreetingEvent
}

func NewGreetingEventList(events []*GreetingEvent) *GreetingEventList {
	return &GreetingEventList{
		events,
	}
}

func (l GreetingEventList) Exists() bool {
	return len(l.Events) > 0
}

func (l GreetingEventList) GetFirstEvent() *GreetingEvent {
	return l.Events[0]
}
