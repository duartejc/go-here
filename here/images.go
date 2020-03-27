package here

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ImagesService provides for HERE Places api.
type ImagesService struct {
	sling *sling.Sling
}

// ImagesParams parameters for Images Service.
type ImagesParams struct {
	Waypoint0  string `url:"waypoint0"`
	Waypoint1  string `url:"waypoint1"`
	Poi0       string `url:"poi0"`
	Poi1       string `url:"poi1"`
	LineColor0 string `url:"lc0"`
	LineColor1 string `url:"lc1"`
	LineWidth0 string `url:"lw0"`
	LineWidth1 string `url:"lw1"`
	Resolution int    `url:"ppi"`
	Width      int    `url:"w"`
	Height     int    `url:"h"`
	APIKey     string `url:"apikey"`
}

// newImagesService returns a new ImagesService.
func newImagesService(sling *sling.Sling) *ImagesService {
	return &ImagesService{
		sling: sling,
	}
}

// CreateImagesParams creates images parameters struct.
func (s *ImagesService) CreateImagesParams(waypoint0 [2]float32, waypoint1 [2]float32, apiKey string) ImagesParams {
	stringWaypoint0 := createWaypoint(waypoint0)
	stringWaypoint1 := createWaypoint(waypoint1)

	imagesParams := ImagesParams{
		Waypoint0: stringWaypoint0,
		Waypoint1: stringWaypoint1,
		APIKey:    apiKey,
	}
	return imagesParams
}

// Routing with given parameters.
func (s *ImagesService) Routing(params *ImagesParams) (interface{}, *http.Response, error) {
	routes := new([]bytes)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("routing").QueryStruct(params).Receive(routes, apiError)
	return routes, resp, relevantError(err, *apiError)
}
