package types

const (
	//RequestCreated is emitted when a request is created
	RequestCreated EventType = "REQUEST_CREATED"
	//RequestApproved is emitted when a request is approved
	RequestApproved EventType = "REQUEST_APPROVED"
)

//RequestEvent
type RequestEvent struct {
	Event
	RequestID string
}
