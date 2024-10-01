package application

import (
	"testing"

	"github.com/01luisfonseca/seligo/src/common"
)

func TestApplicationCRUDCase_GetApplications(t *testing.T) {
	applicationDao := NewMockApplicationDAO()
	t.Run("should return a list of applications", func(t *testing.T) {
		// Arrange
		application := NewApplicationCRUDCase(applicationDao)
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
		applications := application.GetApplications(filter)

		// Assert
		if len(applications.Items) != 0 {
			t.Errorf("Expected 0 applications, got %d", len(applications.Items))
		}
	})
}
