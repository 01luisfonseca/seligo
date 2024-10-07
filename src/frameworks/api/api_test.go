package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFrameworkAPI_Run(t *testing.T) {
	// Arrange
	framework := NewFrameworkAPI(":8080")
	handler := framework.Prepare()
	server := httptest.NewServer(handler)
	defer server.Close()
	t.Run("should run health endpoint", func(t *testing.T) {

		// Act
		resp, err := http.Get(server.URL + "/health")
		if err != nil {
			t.Fatalf("Error al realizar la solicitud: %v", err)
		}

		// Assert
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}
	})
}
