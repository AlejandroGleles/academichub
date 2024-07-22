package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Atividade struct {
	gorm.Model
	Turma string
	Valor int
	Data  string
}

type AtividadeResponse struct {
	ID        uint      `Json:"id"`
	CreatedAt time.Time `Json:"crestedAt"`
	UpdatedAt time.Time `Json:"updateAt"`
	DeletedAt time.Time `Json:"deletedAt,omitempty"`
	Turma     string    `Json:"turma"`
	Valor     string    `Json:"valor"`
	Data      string    `Json:"data"`
}
