package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome      string
	Matricula float64
	Turma     string
}

type AlunoResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Nome      string     `json:"nome"`
	Matricula float64    `json:"matricula"`
	Turma     string     `json:"turma"`
}
