package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Name string `json:"name,omitempty" yaml:"name"`
}
