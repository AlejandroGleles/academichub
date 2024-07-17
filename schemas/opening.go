package schemas

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	Nome  string
	Email string
	CPF   string
}
