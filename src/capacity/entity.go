package capacity

import "github.com/01luisfonseca/seligo/src/common"

type CapacityInputDTO struct {
	LicenseId   common.Id `json:"license_id"`
	Unit        string    `json:"unit"`
	Value       int       `json:"value"`
	Description string    `json:"description,omitempty"`
}

type CapacityRegistryDTO struct {
	common.CommonDTO
	CapacityInputDTO
}

type CapacityInterface interface {
	GetCapacities() []CapacityRegistryDTO
	GetCapacity(id string) CapacityRegistryDTO
	CreateCapacity(capacity CapacityInputDTO) CapacityRegistryDTO
	UpdateCapacity(id string, capacity CapacityInputDTO) CapacityRegistryDTO
	DeleteCapacity(id string) CapacityRegistryDTO
}
