package application

import (
	"errors"
	"strconv"

	"github.com/01luisfonseca/seligo/src/cases/common"
)

var applicationList = []ApplicationRegistryDTO{
	{
		ApplicationInputDTO: ApplicationInputDTO{
			Name:        "Test application",
			Description: "Test application description",
		},
		CommonDTO: common.CommonDTO{
			Id:        "1",
			Status:    common.Status(true),
			CreatedBy: common.CreatedBy("system"),
			UpdatedBy: common.UpdatedBy("system"),
			DeletedBy: common.DeletedBy(""),
			CreatedAt: common.CreatedAt("2023-01-01T00:00:00Z"),
			UpdatedAt: common.UpdatedAt("2023-01-01T00:00:00Z"),
		},
	},
	{
		ApplicationInputDTO: ApplicationInputDTO{
			Name:        "Test application 2",
			Description: "Test application description 2",
		},
		CommonDTO: common.CommonDTO{
			Id:        "2",
			Status:    common.Status(true),
			CreatedBy: common.CreatedBy("system"),
			UpdatedBy: common.UpdatedBy("system"),
			DeletedBy: common.DeletedBy(""),
			CreatedAt: common.CreatedAt("2023-01-01T00:00:00Z"),
			UpdatedAt: common.UpdatedAt("2023-01-01T00:00:00Z"),
		},
	},
}

type MockApplicationDAO struct {
}

func NewMockApplicationDAO() *MockApplicationDAO {
	return &MockApplicationDAO{}
}

func (m *MockApplicationDAO) GetApplications(filter common.CommonFilter) (AplicationListDTO, error) {
	pagination := common.CommonPagination{
		Total:  len(applicationList),
		Limit:  filter.Limit,
		Offset: filter.Offset,
	}
	toReturn := AplicationListDTO{
		Items:      applicationList,
		Pagination: pagination,
	}
	return toReturn, nil
}

func (m *MockApplicationDAO) GetApplication(id common.Id) (ApplicationRegistryDTO, error) {
	for _, application := range applicationList {
		if application.Id == id {
			return application, nil
		}
	}
	return ApplicationRegistryDTO{}, errors.New("application not found")
}

func (m *MockApplicationDAO) CreateApplication(application ApplicationRegistryDTO) (ApplicationRegistryDTO, error) {
	application.Id = common.Id(strconv.Itoa((len(applicationList) + 1)))
	applicationList = append(applicationList, application)
	return application, nil
}

func (m *MockApplicationDAO) UpdateApplication(id common.Id, application ApplicationRegistryDTO) (ApplicationRegistryDTO, error) {
	for i, application := range applicationList {
		if application.Id == id {
			applicationList[i].ApplicationInputDTO.Description = application.ApplicationInputDTO.Description
			applicationList[i].ApplicationInputDTO.Name = application.ApplicationInputDTO.Name
			applicationList[i].CommonDTO.UpdatedAt = application.UpdatedAt
			applicationList[i].CommonDTO.UpdatedBy = application.UpdatedBy
			applicationList[i].CommonDTO.CreatedAt = application.CreatedAt
			applicationList[i].CommonDTO.CreatedBy = application.CreatedBy
			applicationList[i].CommonDTO.DeletedAt = application.DeletedAt
			applicationList[i].CommonDTO.DeletedBy = application.DeletedBy
			return applicationList[i], nil
		}
	}
	return ApplicationRegistryDTO{}, nil
}

func (m *MockApplicationDAO) DeleteApplication(id common.Id) (ApplicationRegistryDTO, error) {
	for i, application := range applicationList {
		if application.Id == id {
			applicationList = append(applicationList[:i], applicationList[i+1:]...)
			return application, nil
		}
	}
	return ApplicationRegistryDTO{}, nil
}
