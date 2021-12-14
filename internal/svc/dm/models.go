package dm

type GlobalTx struct {
	GID string `json:"gid,omitempty" db:"gid"`
}

type SubTx struct {
	GID      string `json:"gid,omitempty" db:"gid"`
	Callback string `json:"callback,omitempty" db:"callback"`
}
