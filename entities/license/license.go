package license

import "github.com/01luisfonseca/seligo/entities/common"

type LicenseInputDTO struct {
	ApplicationId common.Id         `json:"application_id"`
	Name          string            `json:"name"`
	Duration      int               `json:"duration,omitempty"`
	TimeUnit      string            `json:"time_unit,omitempty"`
	Cost          common.MoneyValue `json:"cost,omitempty"`
	Currency      string            `json:"currency,omitempty"`
	Public        bool              `json:"public"`
}

type LicenseRegistryDTO struct {
	common.CommonDTO
	LicenseInputDTO
}

type LicenseInterface interface {
	GetLicenses() []LicenseRegistryDTO
	GetLicense(id string) LicenseRegistryDTO
	CreateLicense(license LicenseInputDTO) LicenseRegistryDTO
	UpdateLicense(id string, license LicenseInputDTO) LicenseRegistryDTO
	DeleteLicense(id string) LicenseRegistryDTO
}
