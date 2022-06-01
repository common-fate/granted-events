package types

const GrantEventPrefix = "GRANT_"

const (
	//GrantStart is emitted when a grant is provisioned successfully
	GrantStart EventType = GrantEventPrefix + "START"
	//GrantEnd is emitted when a grant is deprovisioned successfully
	GrantEnd EventType = GrantEventPrefix + "END"
	//GrantRevoke is emitted when a grant is deprovisioned on demand successfully
	GrantRevoke EventType = GrantEventPrefix + "REVOKE"
	//GrantStartFailure is emitted when a grant fails to be provisioned successfully
	GrantStartFailure EventType = GrantEventPrefix + "START_FAILURE"
	//GrantEndFailure is emitted when a grant fails to be deprovisioned successfully
	GrantEndFailure EventType = GrantEventPrefix + "END_FAILURE"
	//GrantRevokeFailure is emitted when a grant fails to be deprovisioned on demand successfully
	GrantRevokeFailure EventType = GrantEventPrefix + "REVOKE_FAILURE"
)

type GrantError struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}

//GrantEvent
type GrantEvent struct {
	Event
	GrantID string
	Error   *GrantError `json:"error,omitempty"`
}

func (e GrantEvent) HasError() bool {
	return e.Error != nil
}
