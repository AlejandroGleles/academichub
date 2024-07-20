package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	Nome  string
	Email string
	CPF   string
}

type ProfessorResponse struct {
	ID        uint      `Json:"id"`
	CreatedAt time.Time `Json:"crestedAt"`
	UpdatedAt time.Time `Json:"updateAt"`
	DeletedAt time.Time `Json:"deletedAt,omitempty"`
	Nome      string    `Json:"nome"`
	Email     string    `Json:"email"`
	CPF       string    `Json:"cpf"`
}
