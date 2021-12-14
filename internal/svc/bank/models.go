package bank

type Wallet struct {
	BankID  int32 `json:"bank_id,omitempty" db:"bank_id"`
	Balance int64 `json:"balance,omitempty" db:"balance"`
}
