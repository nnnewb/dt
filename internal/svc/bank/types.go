package bank

type TransInReq struct {
	ID     int64 `json:"id,omitempty"`
	Amount int64 `json:"amount,omitempty"`
}

type TransOutReq struct {
	ID     int64 `json:"id,omitempty"`
	Amount int64 `json:"amount,omitempty"`
}
