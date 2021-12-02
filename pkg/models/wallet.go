package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Name    string `json:"name,omitempty" yaml:"name"`
	Balance int64  `json:"balance,omitempty" yaml:"balance"`
}
