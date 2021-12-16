package bank

type TransInReq struct {
	GID    string `json:"gid,omitempty"`
	ID     int64  `json:"id,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

type TransInResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type TransOutReq struct {
	GID    string `json:"gid,omitempty"`
	ID     int64  `json:"id,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

type TransOutResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type TMCallbackReq struct {
	Action   string `json:"action,omitempty"`
	GID      string `json:"gid,omitempty"`
	BranchID string `json:"branch_id,omitempty"`
}

type TMCallbackResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
