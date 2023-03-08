package usecase

import (
	"test-architecture/entity"
)

type CreateUser struct {
	Repository entity.UserRepository
}

func NewCreateUser(repository entity.UserRepository) *CreateUser {
	return &CreateUser{repository}
}

func (cu *CreateUser) Execute(input CreateUserDTOInput) (CreateUserDTOOutput, error) {
	user := entity.NewUserApplication()
	user.ID = input.ID
	user.Name = input.Name
	user.Email = input.Email

	if err := user.IsValid(); err != nil {
		return failCreateUser(input.ID, err), err
	}
	output, err := cu.insertUser(*user)
	if err != nil {
		return failCreateUser(input.ID, err), err
	}

	return output, nil
}

func (cu *CreateUser) insertUser(user entity.User) (CreateUserDTOOutput, error) {
	err := cu.Repository.Insert(user.ID, user.Name, user.Email)
	if err != nil {
		return CreateUserDTOOutput{}, err
	}

	output := CreateUserDTOOutput{
		ID:          user.ID,
		ErrorMesage: "",
		Status:      "SUCCESS",
	}
	return output, nil
}

func failCreateUser(id int, err error) CreateUserDTOOutput {
	output := CreateUserDTOOutput{
		ID:          id,
		ErrorMesage: err.Error(),
		Status:      "INSERT ERROR",
	}

	return output
}
