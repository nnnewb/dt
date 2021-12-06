package bank

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	BankID  int32 `json:"bank_id,omitempty" yaml:"bank_id"`
	Balance int64 `json:"balance,omitempty" yaml:"balance"`
}
