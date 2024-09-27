package policy

import "github.com/01luisfonseca/seligo/entities/common"

type PolicyInputDTO struct {
	ApplicationId common.Id `json:"application_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
}

type PolicyRegistryDTO struct {
	common.CommonDTO
	PolicyInputDTO
}

type PolicyInterface interface {
	GetPolicies() []PolicyRegistryDTO
	GetPolicy(id string) PolicyRegistryDTO
	CreatePolicy(policy PolicyInputDTO) PolicyRegistryDTO
	UpdatePolicy(id string, policy PolicyInputDTO) PolicyRegistryDTO
	DeletePolicy(id string) PolicyRegistryDTO
}
