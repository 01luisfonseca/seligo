package project

import "github.com/01luisfonseca/seligo/entities/common"

type ProjectInputDTO struct {
	ApplicationId common.Id `json:"application_id"`
	StartDate     string    `json:"start_date,omitempty"`
	Name          string    `json:"name"`
}

type ProjectRegistryDTO struct {
	common.CommonDTO
	ProjectInputDTO
}

type ProjectInterface interface {
	GetProjects() []ProjectRegistryDTO
	GetProject(id string) ProjectRegistryDTO
	CreateProject(project ProjectInputDTO) ProjectRegistryDTO
	UpdateProject(id string, project ProjectInputDTO) ProjectRegistryDTO
	DeleteProject(id string) ProjectRegistryDTO
}
