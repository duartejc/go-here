package here

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ImagesService provides for HERE Places api.
type ImagesService struct {
	httpClient *http.Client
	baseURL    string
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
func newImagesService(httpClient *http.Client, baseURL string) *ImagesService {
	return &ImagesService{
		httpClient: httpClient,
		baseURL:    baseURL,
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
func (s *ImagesService) Routing(params *ImagesParams) ([]byte, *http.Response, error) {
	apiError := new(APIError)

	v, _ := query.Values(params)

	buf := bytes.Buffer{}
	buf.WriteString(s.baseURL)
	buf.WriteString(v.Encode())

	reqURL := buf.String()

	resp, err := http.Get(reqURL)
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil, relevantError(err, *apiError)
}
