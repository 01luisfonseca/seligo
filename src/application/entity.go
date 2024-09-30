package application

import "github.com/01luisfonseca/seligo/src/common"

type ApplicationInputDTO struct {
	Name        string `json:"name"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApplicationRegistryDTO struct {
	common.CommonDTO
	ApplicationInputDTO
}

type ApplicationInterface interface {
	GetApplications() []ApplicationRegistryDTO
	GetApplication(id string) ApplicationRegistryDTO
	CreateApplication(application ApplicationInputDTO) ApplicationRegistryDTO
	UpdateApplication(id string, application ApplicationInputDTO) ApplicationRegistryDTO
	DeleteApplication(id string) ApplicationRegistryDTO
}
