package types

const (
	//GrantStart is emitted when a grant is provisioned successfully
	GrantStart EventType = "GRANT_START"
	//GrantEnd is emitted when a grant is deprovisioned successfully
	GrantEnd EventType = "GRANT_END"
	//GrantRevoke is emitted when a grant is deprovisioned on demand successfully
	GrantRevoke EventType = "GRANT_REVOKE"
	//GrantStartFailure is emitted when a grant fails to be provisioned successfully
	GrantStartFailure EventType = "GRANT_START_FAILURE"
	//GrantEndFailure is emitted when a grant fails to be deprovisioned successfully
	GrantEndFailure EventType = "GRANT_END_FAILURE"
	//GrantRevokeFailure is emitted when a grant fails to be deprovisioned on demand successfully
	GrantRevokeFailure EventType = "GRANT_REVOKE_FAILURE"
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
