package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// create openig

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
	//boolean
	//if r.boolean == nil{
	//	return errParamIsRequired("bolean", "bolean")
	//}
	//numero
	//if r.numero <= 0{
	//	return errParamIsRequired("numero", "int")
	//}
	return nil
}

//update
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
