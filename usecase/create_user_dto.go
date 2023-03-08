package usecase

type CreateUserDTOInput struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserDTOOutput struct {
	ID          int    `json:"id"`
	ErrorMesage string `json:"error_message"`
	Status      string `json:"status"`
}
