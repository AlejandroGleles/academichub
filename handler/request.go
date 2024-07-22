package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// create professor

type CreateOpenigRequest struct {
	Nome  string `Json:"nome"`
	Email string `Json:"email"`
	CPF   string `Json:"cpf"`
}

func (r *CreateOpenigRequest) Validate() error {
	if r.Nome == "" && r.Email == "" && r.CPF == "" {
		return fmt.Errorf("malformed request body")
	}
	if r.Nome == "" {
		return errParamIsRequired("nome", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("cpf", "string")
	}

	return nil
}

//update professor
type UpdateOpenigRequest struct {
	Nome  string `Json:"nome"`
	Email string `Json:"email"`
	CPF   string `Json:"cpf"`
}

func (r *UpdateOpenigRequest) Validate() error {
	//if any field is provided, validation is truthy
	if r.Nome != "" || r.Email != "" || r.CPF != "" {
		return nil
	}
	//if none of the fields were providad return false
	return fmt.Errorf("at least one valid field must be provided")
}

// create aluno

type CreateAlunoRequest struct {
	Nome      string `Json:"nome"`
	Matricula int    `Json:"matricula"`
	Turma     string `Json:"turma"`
}

func (r *CreateAlunoRequest) Validate() error {
	if r.Nome == "" && r.Matricula <= 0 && r.Turma == "" {
		return fmt.Errorf("malformed request body")
	}
	if r.Nome == "" {
		return errParamIsRequired("nome", "string")
	}
	if r.Matricula <= 0 {
		return errParamIsRequired("matricula", "int")
	}
	if r.Turma == "" {
		return errParamIsRequired("turma", "string")
	}

	return nil
}

//update aluno
type UpdateAlunoRequest struct {
	Nome      string `Json:"nome"`
	Matricula int    `Json:"matricula"`
	Turma     string `Json:"turma"`
}

func (r *UpdateAlunoRequest) Validate() error {
	//if any field is provided, validation is truthy
	if r.Nome != "" || r.Matricula <= 0 || r.Turma != "" {
		return nil
	}
	//if none of the fields were providad return false
	return fmt.Errorf("at least one valid field must be provided")
}

// create turma

type CreateTurmaRequest struct {
	Nome      string `Json:"nome"`
	Semestre  string `Json:"semestre"`
	Ano       string `Json:"ano"`
	Professor string `Json:"professor"`
}

func (r *CreateTurmaRequest) Validate() error {
	if r.Nome == "" && r.Semestre == "" && r.Ano == "" && r.Professor == "" {
		return fmt.Errorf("malformed request body")
	}
	if r.Nome == "" {
		return errParamIsRequired("nome", "string")
	}
	if r.Semestre == "" {
		return errParamIsRequired("Semestre", "string")
	}
	if r.Ano == "" {
		return errParamIsRequired("Ano", "string")
	}
	if r.Professor == "" {
		return errParamIsRequired("Professor", "string")
	}

	return nil
}

//update turma
type UpdateTurmaRequest struct {
	Nome      string `Json:"nome"`
	Semestre  string `Json:"semestre"`
	Ano       string `Json:"ano"`
	Professor string `Json:"professor"`
}

func (r *UpdateTurmaRequest) Validate() error {
	//if any field is provided, validation is truthy
	if r.Nome != "" || r.Semestre != "" || r.Ano != "" || r.Professor != "" {
		return nil
	}
	//if none of the fields were providad return false
	return fmt.Errorf("at least one valid field must be provided")
}

// create atividade

type CreateAtividadeRequest struct {
	Turma string `Json:"turma"`
	Valor int    `Json:"valor"`
	Data  string `Json:"data"`
}

func (r *CreateAtividadeRequest) Validate() error {
	if r.Turma == "" && r.Valor <= 0 && r.Data == "" {
		return fmt.Errorf("malformed request body")
	}
	if r.Turma == "" {
		return errParamIsRequired("Turma", "string")
	}
	if r.Valor <= 0 {
		return errParamIsRequired("valor", "int")
	}
	if r.Data == "" {
		return errParamIsRequired("data", "string")
	}

	return nil
}

//update atividade
type UpdateAtividadeRequest struct {
	Turma string `Json:"turma"`
	Valor int    `Json:"valor"`
	Data  string `Json:"data"`
}

func (r *UpdateAtividadeRequest) Validate() error {
	//if any field is provided, validation is truthy
	if r.Turma != "" || r.Valor <= 0 || r.Data != "" {
		return nil
	}
	//if none of the fields were providad return false
	return fmt.Errorf("at least one valid field must be provided")
}
