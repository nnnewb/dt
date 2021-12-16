package tm

type GlobalTx struct {
	GID string `json:"gid,omitempty" db:"gid"`
}

type LocalTx struct {
	GID         string `json:"gid,omitempty" db:"gid"`
	BranchID    string `json:"branch_id,omitempty" db:"branch_id"`
	CallbackUrl string `json:"callback_url,omitempty" db:"callback_url"`
}
