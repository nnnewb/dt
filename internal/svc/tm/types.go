package tm

type CreateGlobalTxReq struct {
	GID string `json:"gid,omitempty"`
}

type CreateGlobalTxResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type RegisterLocalTxReq struct {
	GID         string `json:"gid,omitempty"`
	BranchID    string `json:"branch_id,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"`
}

type RegisterLocalTxResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type CommitGlobalTxReq struct {
	GID string `json:"gid,omitempty"`
}

type CommitGlobalTxResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type RollbackGlobalTxReq struct {
	GID string `json:"gid,omitempty"`
}

type RollbackGlobalTxResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
