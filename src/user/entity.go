package user

import "github.com/01luisfonseca/seligo/src/common"

type UserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Address  string `json:"address,omitempty"`
	Country  string `json:"country,omitempty"`
}

type UserRegistryDTO struct {
	common.CommonDTO
	UserInputDTO
}

type UserInterface interface {
	GetUsers() []UserRegistryDTO
	GetUser(id string) UserRegistryDTO
	CreateUser(user UserInputDTO) UserRegistryDTO
	UpdateUser(id string, user UserInputDTO) UserRegistryDTO
	DeleteUser(id string) UserRegistryDTO
}
