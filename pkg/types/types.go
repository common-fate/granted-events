package types

type EventType string

//Event is the basis for all events submitted to eventbridge
type Event struct {
	Type EventType `json:"type"`
}

func (e Event) EventType() EventType {
	return e.Type
}
