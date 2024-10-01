package application

import (
	"time"

	"github.com/01luisfonseca/seligo/src/common"
)

type ApplicationCRUDCase struct {
	application ApplicationInterfaceDAO
}

func NewApplicationCRUDCase(application ApplicationInterfaceDAO) *ApplicationCRUDCase {
	return &ApplicationCRUDCase{
		application: application,
	}
}

func (c *ApplicationCRUDCase) GetApplications(filter common.CommonFilter) AplicationListDTO {
	return c.application.GetApplications(filter)
}

func (c *ApplicationCRUDCase) GetApplication(id string) ApplicationRegistryDTO {
	return c.application.GetApplication(id)
}

func (c *ApplicationCRUDCase) CreateApplication(application ApplicationInputDTO) ApplicationRegistryDTO {
	timeNow := time.Now()
	applicationRegistry := ApplicationRegistryDTO{
		ApplicationInputDTO: application,
		CommonDTO: common.CommonDTO{
			Status:    common.Status(true),
			CreatedBy: common.CreatedBy("system"),
			UpdatedBy: common.UpdatedBy("system"),
			DeletedBy: common.DeletedBy(""),
			CreatedAt: common.CreatedAt(common.Date(timeNow.Format(time.RFC3339))),
			UpdatedAt: common.UpdatedAt(common.Date(timeNow.Format(time.RFC3339))),
		},
	}
	return c.application.CreateApplication(applicationRegistry)
}

func (c *ApplicationCRUDCase) UpdateApplication(id string, application ApplicationInputDTO) ApplicationRegistryDTO {
	currentApplication := c.application.GetApplication(id)
	return c.application.UpdateApplication(id, application)
}

func (c *ApplicationCRUDCase) DeleteApplication(id string) ApplicationRegistryDTO {
	return c.application.DeleteApplication(id)
}
