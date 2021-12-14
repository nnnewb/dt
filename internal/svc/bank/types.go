package bank

type TransInReq struct {
	ID     int64 `json:"id,omitempty"`
	Amount int64 `json:"amount,omitempty"`
}

type TransInResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type TransOutReq struct {
	ID     int64 `json:"id,omitempty"`
	Amount int64 `json:"amount,omitempty"`
}

type TransOutResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
