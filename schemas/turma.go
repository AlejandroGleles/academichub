package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Turma struct {
	gorm.Model
	Nome      string
	Semestre  string
	Ano       string
	Professor string
}

type TurmaResponse struct {
	ID        uint      `Json:"id"`
	CreatedAt time.Time `Json:"crestedAt"`
	UpdatedAt time.Time `Json:"updateAt"`
	DeletedAt time.Time `Json:"deletedAt,omitempty"`
	Nome      string    `Json:"nome"`
	Semestre  string    `Json:"semestre"`
	Ano       string    `Json:"ano"`
	Professor string    `Json:"professor"`
}
