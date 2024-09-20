package configuration

import "github.com/01luisfonseca/seligo/entities/common"

type ConfigurationInputDTO struct {
	Kind        common.Id `json:"kind"`
	KindName    string    `json:"kind_name"`
	Key         string    `json:"key"`
	Value       string    `json:"value"`
	Description string    `json:"description,omitempty"`
}

type ConfigurationRegistryDTO struct {
	common.CommonDTO
	ConfigurationInputDTO
}

type ConfigurationInterface interface {
	GetConfigurations() []ConfigurationRegistryDTO
	GetConfiguration(id string) ConfigurationRegistryDTO
	CreateConfiguration(configuration ConfigurationInputDTO) ConfigurationRegistryDTO
	UpdateConfiguration(id string, configuration ConfigurationInputDTO) ConfigurationRegistryDTO
	DeleteConfiguration(id string) ConfigurationRegistryDTO
}
