package model

type GreetingEventKey struct {
	Key string
}

func NewGreetingEventKey(key string) *GreetingEventKey {
	return &GreetingEventKey{
		Key: key,
	}
}

func (e *GreetingEventKey) Val() string {
	return e.Key
}
