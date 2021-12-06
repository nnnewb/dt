package dm

import "gorm.io/gorm"

type GlobalTx struct {
	gorm.Model
	GID string `json:"gid,omitempty"`
}

type SubTx struct {
	gorm.Model
	GID      string `json:"gid,omitempty"`
	Callback string `json:"callback,omitempty"`
}
