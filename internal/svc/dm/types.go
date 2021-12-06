package dm

type CreateGlobalTransactionReq struct {
	GID string `json:"gid,omitempty"`
}

type RegisterLocalTransactionReq struct {
	GID      string `json:"gid,omitempty"`
	BranchID string `json:"branch_id,omitempty"`
}

type CommitGlobalTransactionReq struct {
	GID string `json:"gid,omitempty"`
}

type RollbackGlobalTransactionReq struct {
	GID string `json:"gid,omitempty"`
}
