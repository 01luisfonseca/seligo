package test

import (
	"fmt"
	"testing"

	app "github.com/01luisfonseca/seligo/src/cases/application"
	"github.com/01luisfonseca/seligo/src/cases/common"
)

func TestApplicationCRUDCase_GetApplications(t *testing.T) {
	applicationDao := app.NewMockApplicationDAO()
	application := app.NewApplicationCRUDCase(applicationDao)
	var id common.Id
	name := "Test application"
	nameUpdated := "Test application 2"

	context := common.CommonContext{
		UserId: "1",
	}

	t.Run("should return a list of applications", func(t *testing.T) {
		// Arrange
		filter := common.CommonFilter{
			Limit:  10,
			Offset: 0,
			Sort: []common.CommonSort{
				{
					Field: "id",
					Order: "asc",
				},
			},
		}

		// Act
		applications, err := application.GetApplications(context, filter)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if len(applications.Items) == 0 {
			t.Errorf("Expected more than 0 applications, got %d", len(applications.Items))
		}
	})

	t.Run("should create an application", func(t *testing.T) {
		// Create
		applicationToCreate := app.ApplicationInputDTO{
			Name:        name,
			Description: "Test application description",
		}

		// Act
		createdApplication, err := application.CreateApplication(context, applicationToCreate)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if createdApplication.Id == "" {
			t.Errorf("Expected an application id, got %s", createdApplication.Id)
		}
		id = createdApplication.Id
		if createdApplication.Name != name {
			t.Errorf("Expected name %s, got %s", name, createdApplication.Name)
		}
	})

	t.Run("should get an application", func(t *testing.T) {
		// Act
		application, err := application.GetApplication(context, id)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if application.Id != id {
			t.Errorf("Expected id %s, got %s", id, application.Id)
		}
		if application.Name != name {
			t.Errorf("Expected name %s, got %s", name, application.Name)
		}
	})

	t.Run("should update an application", func(t *testing.T) {
		// Update
		applicationToUpdate := app.ApplicationInputDTO{
			Name:        nameUpdated,
			Description: "Test application description 2",
		}

		// Act
		updatedApplication, err := application.UpdateApplication(context, id, applicationToUpdate)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if updatedApplication.Id != id {
			t.Errorf("Expected id %s, got %s", id, updatedApplication.Id)
		}
		if updatedApplication.Name != name {
			t.Errorf("Expected name %s, got %s", applicationToUpdate.Name, name)
		}
	})

	t.Run("should soft delete an application", func(t *testing.T) {
		// Act
		deletedApplication, err := application.DeleteApplication(context, id, true)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if deletedApplication.Id != id {
			t.Errorf("Expected id %s, got %s", id, deletedApplication.Id)
		}
	})

	t.Run("should get an application soft deleted", func(t *testing.T) {
		// Act
		application, err := application.GetApplication(context, id)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		fmt.Printf("%+v\n", application)

		// Assert
		if application.Id != id {
			t.Errorf("Expected id %s, got %s", id, application.Id)
		}
		if application.Name != name {
			t.Errorf("Expected name %s, got %s", name, application.Name)
		}
		if application.DeletedAt == "" {
			t.Errorf("Expected deleted_at to be set, got %s", application.DeletedAt)
		}
	})

	t.Run("should hard delete an application", func(t *testing.T) {
		// Act
		deletedApplication, err := application.DeleteApplication(context, id, false)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		// Assert
		if deletedApplication.Id != id {
			t.Errorf("Expected id %s, got %s", id, deletedApplication.Id)
		}
	})
}
