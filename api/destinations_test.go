package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KongAirlines/destinations/api"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetDestinations(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new DestinationService
	destinationService := api.NewDestinationsService()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/destinations", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Create an Echo context
	ctx := e.NewContext(req, rec)

	// Call the GetAllDestinations function
	err := destinationService.GetAllDestinations(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the response status code is 200
	assert.Equal(t, http.StatusOK, rec.Code)
}
