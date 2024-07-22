package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome      string
	Matricula int
	Turma     string
}

type AlunoResponse struct {
	ID        uint      `Json:"id"`
	CreatedAt time.Time `Json:"crestedAt"`
	UpdatedAt time.Time `Json:"updateAt"`
	DeletedAt time.Time `Json:"deletedAt,omitempty"`
	Nome      string    `Json:"nome"`
	Matricula string    `Json:"matricula"`
	Turma     string    `Json:"turma"`
}
