package licensed

import "github.com/01luisfonseca/seligo/src/common"

type LicensedInputDTO struct {
	Kind               common.Id `json:"kind"`
	KindName           string    `json:"kind_name"`
	LicenseId          common.Id `json:"license_id"`
	StartDate          string    `json:"start_date,omitempty"`
	EndDate            string    `json:"end_date,omitempty"`
	ConsumptionUnits   string    `json:"consumption_units"`
	ConsumptionAmmount uint      `json:"consumption_ammount"`
	Description        string    `json:"description,omitempty"`
}

type LicensedRegistryDTO struct {
	common.CommonDTO
	LicensedInputDTO
}

type LicensedInterface interface {
	GetLicenseds() []LicensedRegistryDTO
	GetLicensed(id string) LicensedRegistryDTO
	CreateLicensed(licensed LicensedInputDTO) LicensedRegistryDTO
	UpdateLicensed(id string, licensed LicensedInputDTO) LicensedRegistryDTO
	DeleteLicensed(id string) LicensedRegistryDTO
}
