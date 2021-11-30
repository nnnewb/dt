package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name string `json:"name" yaml:"name"`
}
