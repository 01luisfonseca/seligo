package application

import "github.com/01luisfonseca/seligo/src/cases/common"

type ApplicationInputDTO struct {
	Name        string `json:"name"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApplicationRegistryDTO struct {
	common.CommonDTO
	ApplicationInputDTO
}

type AplicationListDTO struct {
	Pagination common.CommonPagination
	Items      []ApplicationRegistryDTO
}

type ApplicationInterfaceDAO interface {
	GetApplications(filter common.CommonFilter) (AplicationListDTO, error)
	GetApplication(id string) (ApplicationRegistryDTO, error)
	CreateApplication(application ApplicationRegistryDTO) (ApplicationRegistryDTO, error)
	UpdateApplication(id string, application ApplicationRegistryDTO) (ApplicationRegistryDTO, error)
	DeleteApplication(id string) (ApplicationRegistryDTO, error)
}
