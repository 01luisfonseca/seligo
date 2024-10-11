package application

import (
	"time"

	"github.com/01luisfonseca/seligo/src/cases/common"
)

type ApplicationCRUDCase struct {
	application ApplicationInterfaceDAO
}

func NewApplicationCRUDCase(application ApplicationInterfaceDAO) *ApplicationCRUDCase {
	return &ApplicationCRUDCase{
		application: application,
	}
}

func (c *ApplicationCRUDCase) GetApplications(context common.CommonContext, filter common.CommonFilter) (AplicationListDTO, error) {
	return c.application.GetApplications(filter)
}

func (c *ApplicationCRUDCase) GetApplication(context common.CommonContext, id common.Id) (ApplicationRegistryDTO, error) {
	return c.application.GetApplication(id)
}

func (c *ApplicationCRUDCase) CreateApplication(
	context common.CommonContext,
	application ApplicationInputDTO,
) (ApplicationRegistryDTO, error) {
	timeNow := time.Now()
	applicationRegistry := ApplicationRegistryDTO{
		ApplicationInputDTO: application,
		CommonDTO: common.CommonDTO{
			Status:    common.Status(true),
			CreatedBy: common.CreatedBy(context.UserId),
			UpdatedBy: common.UpdatedBy(context.UserId),
			DeletedBy: common.DeletedBy(""),
			CreatedAt: common.CreatedAt(common.Date(timeNow.Format(time.RFC3339))),
			UpdatedAt: common.UpdatedAt(common.Date(timeNow.Format(time.RFC3339))),
		},
	}
	return c.application.CreateApplication(applicationRegistry)
}

func (c *ApplicationCRUDCase) UpdateApplication(
	context common.CommonContext,
	id common.Id,
	application ApplicationInputDTO,
) (ApplicationRegistryDTO, error) {
	currentApplication, err := c.application.GetApplication(id)
	if err != nil {
		return ApplicationRegistryDTO{}, err
	}
	currentApplication.UpdatedBy = common.UpdatedBy(context.UserId)
	currentApplication.UpdatedAt = common.UpdatedAt(common.Date(time.Now().Format(time.RFC3339)))
	return c.application.UpdateApplication(id, currentApplication)
}

func (c *ApplicationCRUDCase) DeleteApplication(
	context common.CommonContext,
	id common.Id,
	softdelete bool,
) (ApplicationRegistryDTO, error) {
	if softdelete {
		currentApplication, err := c.application.GetApplication(id)
		if err != nil {
			return ApplicationRegistryDTO{}, err
		}
		currentApplication.DeletedBy = common.DeletedBy(context.UserId)
		currentApplication.DeletedAt = common.DeletedAt(common.Date(time.Now().Format(time.RFC3339)))
		return c.application.UpdateApplication(id, currentApplication)
	}
	return c.application.DeleteApplication(id)
}
