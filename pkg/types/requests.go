package types

const RequestEventPrefix = "REQUEST_"
const (
	//RequestCreated is emitted when a request is created
	RequestCreated EventType = RequestEventPrefix + "CREATED"
	//RequestApproved is emitted when a request is approved
	RequestApproved EventType = RequestEventPrefix + "APPROVED"
)

//RequestEvent
type RequestEvent struct {
	Event
	RequestID string
}
