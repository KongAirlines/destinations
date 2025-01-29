//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=models.cfg.yaml ../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../openapi.yaml

package api

import (
	"net/http"

	"github.com/KongAirlines/destinations/api/models"
	"github.com/labstack/echo/v4"
)

type DestinationService struct {
	Destinations []models.Destination
}

func NewDestinationService() *DestinationService {
	rv := DestinationService{}
	rv.Destinations = []models.Destination{
		{Id: "JFK", City: "New York", Image: "https://images.unsplash.com/photo-1496442226666-8d4d0e62e6e9"},
		{Id: "SFO", City: "San Francisco", Image: "https://images.unsplash.com/photo-1554107136-57b138ea99df"},
		{Id: "DXB", City: "Dubai", Image: "https://images.unsplash.com/photo-1518684079-3c830dcef090"},
		{Id: "HKG", City: "Hong Kong", Image: "https://images.unsplash.com/photo-1536599018102-9f803c140fc1"},
		{Id: "BLR", City: "Bangalore", Image: "https://images.unsplash.com/photo-1591461321840-c7503e604a16"},
		{Id: "HND", City: "Tokyo", Image: "https://images.unsplash.com/photo-1540959733332-eab4deabeeaf"},
		{Id: "CPT", City: "Cape Town", Image: "https://images.unsplash.com/photo-1576485290814-1c72aa4bbb8e"},
		{Id: "SYD", City: "Sydney", Image: "https://images.unsplash.com/photo-1506973035872-a4ec16b8e8d9"},
		{Id: "SIN", City: "Singapore", Image: "https://images.unsplash.com/photo-1555217851-6141535bd771"},
		{Id: "LAX", City: "Los Angeles", Image: "https://images.unsplash.com/photo-1609924211018-5526c55bad5b"},
		{Id: "LHR", City: "London", Image: "https://images.unsplash.com/photo-1513635269975-59663e0ac1ad"},
		{Id: "PVG", City: "Shanghai", Image: "https://images.unsplash.com/photo-1627484986972-e544190b8abb"},
	}
	return &rv
}

func (s *DestinationService) GetHealthStatus(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (s *DestinationService) GetAllDestinations(ctx echo.Context) error {
	return ctx.JSON(200, s.Destinations)
}

func (s *DestinationService) GetDestinationById(ctx echo.Context, id string) error {
	destinations := s.Destinations
	for _, destination := range destinations {
		if destination.Id == id {
			err := ctx.JSON(200, destination)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ctx.JSON(404, nil)
}
